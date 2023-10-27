package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	userService "totality_corp_assignment/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

func main() {
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Connect to ProtoBuf Service
	c := userService.NewUserServiceClient(conn)

	// Contact the server and fetch all users.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetUsers(ctx, &userService.Nil{})

	if err != nil {
		fmt.Println("\nerror occurred while fetching users", err)
	}
	fmt.Println("\nlist of users", r)

	println("\ntrying to fetch user with id", 3)

	// Fetch User with id 3
	id := userService.UserRequest{Id: 3}

	user, err := c.GetUserById(ctx, &id)

	if err != nil {
		fmt.Println("\nerror occurred while fetching user with id", id.Id, err)
	}

	fmt.Println("\nuser details for id", id.Id, user)

	println("\ntrying to fetch user with id", 10)

	// Fetch User with id 10
	id = userService.UserRequest{Id: 10}

	user, err = c.GetUserById(ctx, &id)

	if err != nil {
		fmt.Println("\nerror occurred while fetching user with id", id.Id, err)
	}

	fmt.Println("\nuser details for id", id.Id, user)

}
