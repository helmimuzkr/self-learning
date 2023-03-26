package main

import (
	"context"
	"flag"
	"log"
	"net"

	pb "github.com/helmimuzkr/02-unary/user"
	"google.golang.org/grpc"
)

type userServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func (uss *userServiceServer) ListUser(ctx context.Context, req *pb.Empty) (*pb.ListUserResponse, error) {
	return &pb.ListUserResponse{}, nil
}

func (uss *userServiceServer) Register(ctx context.Context, req *pb.RegisterUserRequest) (*pb.Empty, error) {
	log.Println("register process on server...")
	// return empty and nil error
	return new(pb.Empty), nil
}

var (
	host    = flag.String("host", "localhost", "type the user server host")
	port    = flag.String("port", "5001", "type the user server port")
	address = *host + ":" + *port
)

func main() {
	// Parse the flag
	flag.Parse()
	// Set up server
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen port %s : %s", listener.Addr(), err)
	}
	server := grpc.NewServer()
	// Register the user service server
	userServiceServer := &userServiceServer{}
	pb.RegisterUserServiceServer(server, userServiceServer)
	// Start listen
	log.Printf("server listening at %v", listener.Addr())
	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve grpc server : %s", err)
	}

}
