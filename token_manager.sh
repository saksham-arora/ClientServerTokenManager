# Running server in the background
go build server/tokenserver.go 
./tokenserver -port 50051 &

# Taking process id of last background job
server_pid=$!

# Running client in background
go build client/tokenclient.go 

# Testing concurrency by running an expensive operation and then showing that when
# cheap operation runs it will run concurrently take the cheap operation request.

#Cheap operation
./tokenclient -create -id 1 -host localhost -port 50051 &

#Expensive operation - 1
i=1
while [ $i -le 100 ]
do
./tokenclient -write -id 1 -name abc -low 0 -mid 10 -high 100 -host localhost -port 50051 &
((i++))
done

#Cheap operations
#Result -> Token would be deleted while operation 1 would be still running.
./tokenclient -drop -id 1 -host localhost -port 50051 &
./tokenclient -create -id 2 -host localhost -port 50051 &
./tokenclient -write -id 2 -name abc -low 0 -mid 10 -high 100 -host localhost -port 50051 &

read=1
while [ $read -le 100 ]
do
./tokenclient -read -id 2 -host localhost -port 50051 &
((read++))
done
./tokenclient -drop -id 2 -host localhost -port 50051 &

# echo $server_pid


# go build client/tokenclient.go
# ./tokenclient -create -id 1 -host localhost -port 50051
# ./tokenclient -write -id 1 -name abc -low 0 -mid 10 -high 100 -host localhost -port 50051
# ./tokenclient -read -id 1 -host localhost -port 50051
# ./tokenclient -drop -id 1 -host localhost -port 50051

# Waiting for other jobs to be completed
sleep 10
# Killing server once the job is completed.
kill $server_pid