package main

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net"
	"strconv"
	"sync"
	"time"

	pb "example.com/token_client_server_rpc/token_management"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
)

type TokenManagerServer struct {
	pb.UnimplementedTokenManagerServer
}

type token struct {
	id        int32
	name      string
	domain    [3]uint64
	state     [2]uint64
	timestamp string
	mutex     sync.RWMutex
}

// Reading YAML Structure
type read_yaml struct {
	Token  int32    `yaml:"token_id"`
	Writer string   `yaml:"writer"`
	Reader []string `yaml:"reader,flow"`
}

// Maintaining readlist for write-back
type readlist struct {
	Timestamp string
	FinalVal  uint64
	Low       uint64
	Mid       uint64
	High      uint64
	Name      string
}

var ReadConfig []read_yaml
var TokenList []token
var ReadList []readlist
var reading_flag bool
var cnt int
var read_flag int

// var TokenList map[int32]token
// Creating Token
func (s *TokenManagerServer) CreateNewToken(ctx context.Context, in *pb.Token) (*pb.Response, error) {
	log.Printf("Token ID Received to Create Token: %v", in.GetId())

	var create_res = "Server Response -> Token Created!"
	for _, token := range TokenList {
		if token.id == in.GetId() {
			token.mutex.Lock()
			token.state[0] = 0
			token.state[1] = 0
			token.mutex.Unlock()
			create_res = "Server Response -> Token Exists. State values of the Token Updated"
			log.Printf(create_res)
			return &pb.Response{CreateResponse: create_res}, nil
		}
	}

	created_token := token{id: in.GetId(), mutex: sync.RWMutex{}}
	TokenList = append(TokenList, created_token)
	// TokenList[in.GetId()] = created_token
	log.Printf("Server Response -> Token Created!")
	return &pb.Response{CreateResponse: create_res}, nil
}

//Reference - https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
func RemoveIndex(s []token, index int) []token {
	return append(s[:index], s[index+1:]...)
}

// Delete Token
func (s *TokenManagerServer) DropToken(ctx context.Context, in *pb.Token) (*pb.Response, error) {
	log.Printf("Token ID received to Drop Token: %v", in.GetId())
	for i, token_desc := range TokenList {
		if token_desc.id == in.GetId() {
			token_desc.mutex.Lock()
			TokenList = RemoveIndex(TokenList, i)
			token_desc.mutex.Unlock()
			log.Printf("Server Response -> Token Deleted!")
			return &pb.Response{CreateResponse: "Server Response -> Success! Token Deleted!"}, nil
		}
	}
	return &pb.Response{CreateResponse: "Failed as token doesn't exists."}, nil
}

// Calculating Hash value
func Hash(name string, nonce uint64) uint64 {
	hasher := sha256.New()
	hasher.Write([]byte(fmt.Sprintf("%s %d", name, nonce)))
	return binary.BigEndian.Uint64(hasher.Sum(nil))
}

func (s *TokenManagerServer) WriteToken(ctx context.Context, in *pb.WriteTokenMsg) (*pb.WriteResponse, error) {
	log.Printf("Token ID Received to Write: %v", in.GetId())
	var flag_token_list int
	// Checking if token exists or not
	for _, token := range TokenList {
		if token.id == in.GetId() {
			flag_token_list = 1
			break
		}
	}
	var read_high, read_mid, read_low uint64
	var read_name string
	var partial_val uint64 = math.MaxUint64
	if flag_token_list != 1 {
		log.Printf("Token ID doesn't exist!")
		return &pb.WriteResponse{CreateWriteResponse: partial_val}, nil
	}
	for x := in.GetLow(); x < in.GetMid(); x++ {
		var h_val = Hash(in.GetName(), x)
		if h_val < partial_val {
			partial_val = h_val
		}
	}
	// var ts *timestamppb.Timestamp
	for i, token := range TokenList {
		if token.id == in.GetId() {
			token.mutex.Lock()
			token_write := TokenList[i]
			token_write.name = in.GetName()
			token_write.domain[0] = in.GetLow()
			token_write.domain[1] = in.GetMid()
			token_write.domain[2] = in.GetHigh()
			token_write.state[0] = partial_val
			read_name = in.GetName()
			read_high = in.GetHigh()
			read_mid = in.GetMid()
			read_low = in.GetLow()
			token_write.state[1] = 0
			// Updating value with current timestamp
			token_write.timestamp = (time.Now()).String()
			TokenList[i] = token_write
			token.mutex.Unlock()
		}
	}

	fmt.Println("Main Server -> Write Broadcast Started")
	cnt_ch := make(chan int)
	// var wg sync.WaitGroup
	// wg.Add(3)
	cnt = 0
	reading_flag = false
	index := 0
	// ack_cnt := 0
	for idx, i := range TokenList {
		if i.id == in.GetId() {
			index = idx
		}
	}
	for _, server := range ReadConfig {
		if server.Token == in.GetId() {
			for _, reader_server := range server.Reader {
				// fmt.Println(reader_server)
				// Starting Broadcast asynchronously
				// Using 99 because assuming len_readlist would be max 99(i.e nodes)
				go startBroadcast(partial_val, read_name, read_high, read_mid, read_low, 99, TokenList[index].timestamp, reader_server, reading_flag, &cnt, cnt_ch, server.Token)
				// time.Sleep(5 * time.Second)
				// cnt := <-cnt_ch
			}
			for {
				//Taking count of ack thru channeling
				ack_cnt := <-cnt_ch
				// time.Sleep(5 * time.Second)
				fmt.Println(ack_cnt)
				// Checking for the majority if N/2 servers are written then it will return the partial value.
				if ack_cnt > len(server.Reader)/2 {
					fmt.Println(ack_cnt)
					log.Printf("Main Server -> Majority Reached!")
					ack_cnt = 0
					// break
					return &pb.WriteResponse{CreateWriteResponse: partial_val}, nil
				}
			}
		}
	}
	// wg.Wait()
	log.Printf("Server Response -> Token Write Completed!")
	return &pb.WriteResponse{CreateWriteResponse: partial_val}, nil
}

func startBroadcast(partial_val uint64, read_name string, read_high uint64, read_mid uint64, read_low uint64, len_readlist int, wts string, server string, reading_flag bool, cnt *int, cnt_ch chan int, token int32) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Not able to Connect: %v", err)
	}
	// defer conn.Close()
	c := pb.NewTokenManagerClient(conn)
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()

	write_req, err := c.WriteBroadcast(context.Background(), &pb.WriteBroadcastRequest{HashVal: partial_val, Low: read_low, Mid: read_mid, High: read_high, Name: read_name, Wts: wts, Server: server, TokenId: token, ReadingFlag: reading_flag})
	if write_req.GetAck() == 1 {
		*cnt += 1
		cnt_ch <- *cnt
	}
	if err != nil {
		log.Fatalf("Couldn't Write Token: %v", err)
	}
	if *cnt > len_readlist/2 {
		read_flag = 1
	}

}

func (s *TokenManagerServer) WriteBroadcast(ctx context.Context, in *pb.WriteBroadcastRequest) (*pb.WriteBroadcastResponse, error) {

	if in.ReadingFlag == true {
		fmt.Println("Write-Back Started")
	} else {
		if in.Server == "127.0.0.1:50052" {
			fmt.Println("Simulation -> High Latency")
			time.Sleep(60 * time.Second)
			// return &pb.WriteBroadcastResponse{Ack: 0}, nil
		}
		fmt.Println("Write Started")
	}
	for i, token := range TokenList {
		if token.id == in.TokenId {
			if in.GetWts() > token.timestamp {
				token.mutex.Lock()
				token_write := TokenList[i]
				token_write.name = in.GetName()
				token_write.domain[0] = in.GetLow()
				token_write.domain[1] = in.GetMid()
				token_write.domain[2] = in.GetHigh()
				token_write.timestamp = in.GetWts()
				token_write.state[0] = in.HashVal
				// fmt.Println(token_write)
				TokenList[i] = token_write
				token.mutex.Unlock()
			} else {
				fmt.Println("Write-Back not done as timestamp is less then or equal to current timestamp.")
			}
		}
	}
	fmt.Println(TokenList)
	log.Printf(`Response from the Server: ACK`)
	return &pb.WriteBroadcastResponse{Ack: 1}, nil
}

func (s *TokenManagerServer) ReadToken(ctx context.Context, in *pb.Token) (*pb.WriteResponse, error) {
	log.Printf("Token ID Received to Read: %v", in.GetId())
	// var final_val uint64 = math.MaxUint64
	var read_mid uint64
	var read_high uint64
	var read_partial_val uint64 = math.MaxUint64
	var read_name string
	var read_low uint64
	var flag int
	for _, curr_token := range TokenList {
		if curr_token.id == in.GetId() {
			read_low = curr_token.domain[0]
			read_mid = curr_token.domain[1]
			read_high = curr_token.domain[2]
			read_name = curr_token.name
			read_partial_val = curr_token.state[0]
			flag = 1
		}
	}
	// print(read_name)
	if flag != 1 {
		return &pb.WriteResponse{CreateWriteResponse: read_partial_val}, nil
	}

	// Commented this part as we won't be working on final value i.e we won't be calculating
	// hash value for mid to high range

	// for x := read_mid; x < read_high; x++ {
	// 	var h_val = Hash(read_name, x)
	// 	if h_val < final_val {
	// 		final_val = h_val
	// 	}
	// }
	// if final_val > read_partial_val {
	// 	final_val = read_partial_val
	// }
	// for i, token := range TokenList {
	// 	if token.id == in.GetId() {
	// 		token.mutex.RLock()
	// 		token_write := TokenList[i]
	// 		token_write.state[1] = final_val
	// 		token_write.timestamp = time.Now().String()
	// 		TokenList[i] = token_write
	// 		token.mutex.RUnlock()
	// 	}
	// }

	cnt_ch := make(chan int)
	readlist_ch := make(chan readlist)
	// ack_cnt := 0
	// var wg sync.WaitGroup
	// wg.Add(3)

	maxts := time.Date(2020, 8, 15, 14, 30, 45, 100, time.Local).String()
	reading_flag := true
	cnt = 0
	read_flag = 0
	index := 0
	for idx, i := range TokenList {
		if i.id == in.GetId() {
			index = idx
		}
	}
	for _, server := range ReadConfig {
		if server.Token == in.GetId() {
			for _, reader_server := range server.Reader {
				// fmt.Println(reader_server)
				go startReadBroadcast(server.Token, TokenList[index].timestamp, reader_server, readlist_ch)
				// time.Sleep(2 * time.Second)
			}
			for {
				// Maintaing the count of readlist
				ReadList = append(ReadList, <-readlist_ch)
				// fmt.Println(ReadList)
				if len(ReadList) > len(server.Reader)/2 {
					log.Printf("Main Server -> Majority Reached for Read Request!")
					for _, s := range ReadList {
						if s.Timestamp > maxts {
							maxts = s.Timestamp
							read_partial_val = s.FinalVal
							read_low = s.Low
							read_mid = s.Mid
							read_high = s.High
							read_name = s.Name
						}
					}
					ReadList := []readlist{}
					ReadList = append(ReadList)
					for _, server := range server.Reader {
						go startBroadcast(read_partial_val, read_name, read_high, read_mid, read_low, len(ReadList), maxts, server, reading_flag, &cnt, cnt_ch, in.GetId())
						if read_flag == 1 {
							read_flag = 0
							return &pb.WriteResponse{CreateWriteResponse: read_partial_val}, nil
						}
					}
					log.Printf("Main Server -> Token Read Completed!")
					break
				}

			}
		}
	}
	log.Printf("Server Response -> Token Read Completed!")
	return &pb.WriteResponse{CreateWriteResponse: read_partial_val}, nil
}

func startReadBroadcast(token_id int32, wts string, server string, readlist_ch chan readlist) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Not able to Connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTokenManagerClient(conn)
	// Starting read broadcast
	write_req, err := c.ReadBroadcast(context.Background(), &pb.ReadBroadcastRequest{TokenId: token_id})
	var FinalVal uint64
	var ts string
	FinalVal = write_req.FinalVal
	ts = write_req.GetWts()
	// ReadList = append(ReadList, readlist{ts, FinalVal})
	readlist_ch <- readlist{ts, FinalVal, write_req.Low, write_req.Mid, write_req.High, write_req.Name}
	fmt.Println(ReadList)
	if err != nil {
		log.Fatalf("Couldn't Write Token: %v", err)
	}
}

func (s *TokenManagerServer) ReadBroadcast(ctx context.Context, in *pb.ReadBroadcastRequest) (*pb.ReadBroadcastResponse, error) {
	fmt.Println("Read Broadcast Started")
	// fmt.Println(TokenList)
	var partial_val, low, mid, high uint64
	var ts, name string
	for _, i := range TokenList {
		if i.id == in.GetTokenId() {
			partial_val = i.state[0]
			name = i.name
			low = i.domain[0]
			mid = i.domain[1]
			high = i.domain[2]
			ts = i.timestamp
		}
	}
	log.Printf(`Response from the Server: Read token send`)
	return &pb.ReadBroadcastResponse{Wts: ts, FinalVal: partial_val, Name: name, Low: low, Mid: mid, High: high}, nil
}

func main() {
	fmt.Println("Reading YAML Configuration")
	yfile, err := ioutil.ReadFile("token_config.yml")
	if err != nil {
		log.Fatal(err)
	}

	// servers := []string{}
	// reader_server := ""
	// writer_server := ""
	port_input := flag.Int("port", 50051, "Port Number")
	flag.Parse()
	var port = ":" + strconv.Itoa(*port_input)

	// r_yml := &read_yaml{}

	err2 := yaml.Unmarshal(yfile, &ReadConfig)
	fmt.Println(ReadConfig)
	fmt.Println("Configuration Read Complete")
	if err2 != nil {
		log.Fatal(err2)
	}
	for _, config := range ReadConfig {
		if config.Writer == "127.0.0.1"+port {
			created_token := token{id: config.Token, timestamp: time.Now().String(), mutex: sync.RWMutex{}}
			TokenList = append(TokenList, created_token)
			fmt.Println(created_token)
		}
		for _, server := range config.Reader {
			if server != config.Writer {
				if server == "127.0.0.1"+port {
					created_token := token{id: config.Token, timestamp: time.Now().String(), mutex: sync.RWMutex{}}
					TokenList = append(TokenList, created_token)
					fmt.Println(created_token)
				}
			}
		}
	}

	lis, err := net.Listen("tcp", "127.0.0.1"+port)
	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTokenManagerServer(s, &TokenManagerServer{})
	log.Printf("Server Listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve: %v", err)
	}
}
