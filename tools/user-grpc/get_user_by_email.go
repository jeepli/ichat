package main

import (
	"context"
	"flag"
	"log"
	"time"

	userpb "github.com/jeepli/ichat/proto/user"
	"google.golang.org/grpc"
)

const (
	address = "127.0.0.1:5050"
)

var (
	email = flag.String("email", "", "email")
)

func main() {
	flag.Parse()
	if *email == "" {
		log.Fatal("bad args")
	}
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := userpb.NewUserServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetUserByEmail(ctx, &userpb.GetUserByEmailRequest{
		Email: *email,
	})
	if err != nil {
		log.Fatalf("get user err: %v", err)
	}
	log.Printf("get user: %s", r.GetUsers())
}
