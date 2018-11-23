package server

import (
	"context"

	"github.com/jeepli/ichat/app/user-grpc/db"
	userpb "github.com/jeepli/ichat/proto/user"
	"github.com/pkg/errors"
)

func (s *UserServer) CreateUser(ctx context.Context, in *userpb.CreateUserRequest) (*userpb.CreateUserReply, error) {

	toInsert := PbUserToDbUser(in.GetUser())
	inserted, err := s.userDb.InsertUser(*toInsert)
	if err != nil {
		return nil, err
	}

	if inserted == nil {
		return nil, errors.Errorf("inserted failed")
	}

	return &userpb.CreateUserReply{
		User: &userpb.User{
			Id:       inserted.Id,
			Name:     inserted.Name,
			Password: inserted.Password,
			Email:    inserted.Email,
		},
	}, nil
}

func (s *UserServer) GetUsers(ctx context.Context, in *userpb.GetUsersRequest) (*userpb.GetUsersReply, error) {
	users, err := s.userDb.SelectUsersByIds(in.GetIds())
	if err != nil {
		return nil, err
	}

	reply := userpb.GetUsersReply{}
	for _, u := range users {
		pbu := DbUserToPbUser(&u)
		reply.Users = append(reply.Users, pbu)
	}
	return &reply, nil
}

func (s *UserServer) GetUserByEmail(ctx context.Context, in *userpb.GetUserByEmailRequest) (*userpb.GetUsersReply, error) {
	users, err := s.userDb.SelectUserByEmail(in.GetEmail())
	if err != nil {
		return nil, err
	}

	reply := userpb.GetUsersReply{}
	for _, u := range users {
		pbu := DbUserToPbUser(&u)
		reply.Users = append(reply.Users, pbu)
	}
	return &reply, nil
}

func DbUserToPbUser(from *db.User) *userpb.User {
	if from == nil {
		return nil
	}

	to := userpb.User{
		Id:       from.Id,
		Name:     from.Name,
		Email:    from.Email,
		Password: from.Password,
	}

	return &to
}

func PbUserToDbUser(from *userpb.User) *db.User {
	if from == nil {
		return nil
	}

	to := db.User{
		Id:       from.GetId(),
		Name:     from.GetName(),
		Email:    from.GetEmail(),
		Password: from.GetPassword(),
	}

	return &to
}
