package api

import (
	"fmt"
	"net"

	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/config"
	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/pb"
	"google.golang.org/grpc"
)

type Server struct {
	server *grpc.Server
	lis    net.Listener
}

func SetupAuthServer(server pb.AuthServiceServer, cfg *config.Config) (*Server, error) {

	grpcServer := grpc.NewServer()

	lis, err := net.Listen("tcp", cfg.ServicePort)
	if err != nil {
		return nil, fmt.Errorf("failed to create net listener error:%s", err.Error())
	}
	pb.RegisterAuthServiceServer(grpcServer, server)

	return &Server{
		server: grpcServer,
		lis:    lis,
	}, nil
}

func (c *Server) Start() error {
	fmt.Println("auth service listening...")
	return c.server.Serve(c.lis)
}
