package main

import (
	"context"
	"fmt"
	"log"
	"net"
	userService "totality_corp_assignment/grpc"

	"google.golang.org/grpc"
)

type Server struct {
	userService.UnimplementedUserServiceServer

	userDatabase map[int32]userService.UserResponse
}

func (s *Server) GetUserById(ctx context.Context, req *userService.UserRequest) (*userService.UserResponse, error) {
	user, ok := s.userDatabase[req.Id]

	if !ok {
		return nil, fmt.Errorf("user with ID %d not found", req.Id)
	}

	return &user, nil
}

func (s *Server) GetUsers(ctx context.Context, req *userService.Nil) (*userService.UsersResponse, error) {
	usersList := userService.UsersResponse{Users: []*userService.UserResponse{}}

	for _, user := range s.userDatabase {
		usersList.Users = append(usersList.Users, &user)
	}

	return &usersList, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	usersDB := map[int32]userService.UserResponse{
		1: {
			Id:      1,
			Fname:   "Steve",
			City:    "LA",
			Phone:   1234567890,
			Height:  5.8,
			Married: true,
		},
		2: {
			Id:      2,
			Fname:   "Stephen",
			City:    "BLR",
			Phone:   1234567890,
			Height:  5.8,
			Married: false,
		},
		3: {
			Id:      3,
			Fname:   "Jake",
			City:    "NY",
			Phone:   1234567890,
			Height:  5.8,
			Married: true,
		},
		4: {
			Id:      4,
			Fname:   "Jackson",
			City:    "LO",
			Phone:   1234567890,
			Height:  5.8,
			Married: false,
		},
	}

	userService.RegisterUserServiceServer(s, &Server{userDatabase: usersDB})

	fmt.Println("gRPC server is running on :50051")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
