package server

import (
	"log"
	"net"

	"github.com/jeepli/ichat/app/user-grpc/config"
	"github.com/jeepli/ichat/app/user-grpc/db"
	"github.com/jeepli/ichat/database"
	userpb "github.com/jeepli/ichat/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type UserServer struct {
	addr   string
	userDb *db.UserDB
}

func NewServer(conf *config.Config) *UserServer {
	dh := database.NewDBHolder(&conf.Db)
	return &UserServer{
		userDb: db.NewUserDb(dh),
		addr:   conf.Service.Addr,
	}
}

func (s *UserServer) Start() {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	ss := grpc.NewServer()
	userpb.RegisterUserServiceServer(ss, s)

	// Register reflection service on gRPC server.
	reflection.Register(ss)
	if err := ss.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
