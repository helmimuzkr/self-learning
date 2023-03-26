package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	pb "github.com/helmimuzkr/02-unary/user"
	"google.golang.org/grpc"
)

var (
	host    = flag.String("host", "localhost", "type the user server host")
	port    = flag.String("port", "5001", "type the user server port")
	address = *host + ":" + *port
)

func main() {
	// Parse the flag
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial client: %v", err)
	}
	defer conn.Close()

	// Create user client
	client := pb.NewUserServiceClient(conn)

	// Contact the user server and print out its response
	ctx := context.TODO()
	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancel()
	// Create request for register
	var firstName, lastName, dateOfBirth string
	fmt.Printf("Insert first name: ")
	fmt.Scanln(&firstName)
	fmt.Printf("Insert last name: ")
	fmt.Scanln(&lastName)
	fmt.Printf("Insert date of birth: ")
	fmt.Scanln(&dateOfBirth)
	request := &pb.RegisterUserRequest{
		FirstName:   firstName,
		LastName:    lastName,
		DateOfBirth: dateOfBirth,
	}
	// Execute the method register
	_, err = client.Register(ctx, request)
	if err != nil {
		log.Fatalf("cloud not register user: %v", err)
	}
	log.Printf("user with first name %s has been registered\n", request.GetFirstName())
}
