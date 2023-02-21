package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	pb "example.com/token_client_server_rpc/token_management"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
)

const (
	address = "127.0.0.1:50051"
)

type read_yaml struct {
	Token  int32    `yaml:"token_id"`
	Writer string   `yaml:"writer"`
	Reader []string `yaml:"reader,flow"`
}

var ReadConfig []read_yaml

func main() {
	var access_check int
	port_input := flag.Int("port", 50051, "Port Number")
	host_input := flag.String("host", "127.0.0.1", "Host Name")
	create_input := flag.Bool("create", false, "Create Operation")
	drop_input := flag.Bool("drop", false, "Drop Operation")
	write_input := flag.Bool("write", false, "Write Operation")
	read_input := flag.Bool("read", false, "Read Operation")
	name_input := flag.String("name", "abc", "Token Name")
	low_input := flag.Int("low", 0, "Domain Low")
	mid_input := flag.Int("mid", 10, "Domain Mid")
	high_input := flag.Int("high", 100, "Domain High")
	id_input := flag.Int("id", 1001, "Token ID")
	flag.Parse()

	// Setting up the client side to connect from the server.
	var address = *host_input + ":" + strconv.Itoa(*port_input)
	var conn *grpc.ClientConn
	fmt.Println(address)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Not able to Connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTokenManagerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fmt.Println("Reading YAML Configuration")
	yfile, err := ioutil.ReadFile("token_config.yml")
	if err != nil {
		log.Fatal(err)
	}
	err2 := yaml.Unmarshal(yfile, &ReadConfig)
	fmt.Println(ReadConfig)
	fmt.Println("Configuration Read Complete")
	if err2 != nil {
		log.Fatal(err2)
	}

	if *create_input {
		create_req, err := c.CreateNewToken(ctx, &pb.Token{Id: int32(*id_input)})
		if err != nil {
			log.Fatalf("Couldn't Create Token: %v", err)
		}
		log.Printf(`Response from the Server: %v`, create_req.GetCreateResponse())
	}

	if *drop_input {
		drop_req, err := c.DropToken(ctx, &pb.Token{Id: int32(*id_input)})
		if err != nil {
			log.Fatalf("Couldn't Drop Token: %v", err)
		}
		log.Printf(`Response from the Server: %v`, drop_req.GetCreateResponse())
	}

	if *write_input {
		fmt.Println(address)
		for _, config := range ReadConfig {
			if config.Token == int32(*id_input) && config.Writer == address {
				write_req, err := c.WriteToken(ctx, &pb.WriteTokenMsg{Id: int32(*id_input), Name: *name_input, Low: uint64(*low_input), Mid: uint64(*mid_input), High: uint64(*high_input)})
				access_check = 1
				if err != nil {
					log.Fatalf("Couldn't Write Token: %v", err)
				}
				log.Printf(`Response from the Server: %v`, write_req.GetCreateWriteResponse())
			}
		}
		if access_check != 1 {
			log.Printf("Unauthorized to do write request")
		}
	}
	if *read_input {
		for _, config := range ReadConfig {
			if config.Token == int32(*id_input) {
				for _, server := range config.Reader {
					if server == address {
						access_check = 1
					}
				}
			}
		}
		if access_check == 1 {
			read_req, err := c.ReadToken(ctx, &pb.Token{Id: int32(*id_input)})
			if err != nil {
				log.Fatalf("Couldn't Read Token: %v", err)
			}
			log.Printf(`Response from the Server: %v`, read_req.GetCreateWriteResponse())
		} else {
			log.Printf("Unauthorized to do read request")
		}
	}
}
