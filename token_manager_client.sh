go build client/tokenclient.go
./tokenclient -create -id 1234 -host localhost -port 50051
./tokenclient -write -id 1234 -name abc -low 0 -mid 10 -high 100 -host localhost -port 50051
./tokenclient -read -id 1234 -host localhost -port 50051
./tokenclient -drop -id 1234 -host localhost -port 50051
