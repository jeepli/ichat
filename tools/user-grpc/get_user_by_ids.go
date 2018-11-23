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
	id = flag.String("id", "", "id")
)

func main() {
	flag.Parse()
	if *id == "" {
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
	r, err := c.GetUsers(ctx, &userpb.GetUsersRequest{
		Ids: []string{*id},
	})
	if err != nil {
		log.Fatalf("get users err: %v", err)
	}
	log.Printf("get users: %s", r.GetUsers())
}
