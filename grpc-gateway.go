syntax = "proto3";

package users;

service UserService {
  rpc GetUser(GetUserRequest) returns (User);
}

message GetUserRequest {
  string user_id = 1;
}

message User {
  string name = 1;
  string occuptation = 2;
}

package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
    pb "path/to/users"
)

type server struct{}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
    // Retrieve user information from database or some other source
    user := &pb.User{
        Name: "John Doe",
        Occuptation:  "Software Engineer",
    }
    return user, nil
}

func main() {
    lis, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterUserServiceServer(s, &server{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}


package main

import (
    "context"
    "log"
    "net/http"

    "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
    "google.golang.org/grpc"
    pb "path/to/users"
)

func main() {
    ctx := context.Background()
    mux := runtime.NewServeMux()

    opts := []grpc.DialOption{grpc.WithInsecure()}
    err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "localhost:8080", opts)
    if err != nil {
        log.Fatalf("failed to register gateway: %v", err)
    }

    if err := http.ListenAndServe(":8081", mux); err != nil {
        log.Fatalf("failed to serve gateway: %v", err)
    }
}

