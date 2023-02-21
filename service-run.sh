# Approach

# Take servers in one array
# Boot all the servers
# Run client request
# kill servers

servers=("50051" "50052" "50053" "50055" "50056" "50057" "50058" "50059" "50060")

test_servers=("50051" "50052" "50053")
test_servers_pid=()

go build server/tokenserver.go 

for server in ${test_servers[@]}; do
  ./tokenserver -port $server > $server.txt & 
  server_pid=$!
  test_servers_pid+=($server_pid)
done

# echo "${test_servers_pid[@]}" 

go build client/tokenclient.go
./tokenclient -write -id 1 -name abc -low 5 -mid 25 -high 100 -host 127.0.0.1 -port 50051 > write.txt

./tokenclient -read -id 1 -host 127.0.0.1 -port 50052  > read.txt

sleep 90

for pid in ${test_servers_pid[@]}; do
    # echo $pid
    kill $pid
done
