## Simple GRPC service in Golang for transferring data between client and server

If you want to generate protocol buffer for the project go to *grpc* official website and follow instructions to download and install `protoc` to your machine and run the command

`protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./user_service.proto`

### Steps to reciprocate to run the repository
  - Install Go Compiler with version greater than 1.20
  - `cd` to the repository's directory.
  - run `go mod download` to download all the necessary package.
  - Open first terminal instance and run `go run ./server/main.go`, and keep the terminal instance running
  - Open another terminal instance and run `go run ./client/main.go`

*Following the above steps should yield you output similar to this in the client's terminal instance*

```shell
list of users users:{id:4  fname:"Jackson"  city:"LO"  phone:1234567890  height:5.8}  users:{id:4  fname:"Jackson"  city:"LO"  phone:1234567890  height:5.8}  users:{id:4  fname:"Jackson"  city:"LO"  phone:1234567890  height:5.8}  users:{id:4  fname:"Jackson"  city:"LO"  phone:1234567890  height:5.8}

trying to fetch user with id 3

user details for id 3 id:3  fname:"Jake"  city:"NY"  phone:1234567890  height:5.8  Married:true

trying to fetch user with id 10

error occurred while fetching user with id 10 rpc error: code = Unknown desc = user with ID 10 not found

user details for id 10 <nil>
```