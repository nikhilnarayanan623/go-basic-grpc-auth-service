package api

import (
	"fmt"
	"net"

	service "github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/api/service/interfaces"
	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/config"
	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/pb/authpb"
	"google.golang.org/grpc"
)

type Server struct {
	server *grpc.Server
	lis    net.Listener
}

func SetupAuthServer(server service.ServiceServer, cfg *config.Config) (*Server, error) {

	grpcServer := grpc.NewServer()

	lis, err := net.Listen("tcp", cfg.ServicePort)
	if err != nil {
		return nil, fmt.Errorf("faild to start tcp connection error:%s", err.Error())
	}
	authpb.RegisterAuthServiceServer(grpcServer, server)
	return &Server{
		server: grpcServer,
		lis:    lis,
	}, nil
}

func (c *Server) Start() (err error) {
	fmt.Println("auth service listening...")
	err = c.server.Serve(c.lis)
	return
}
