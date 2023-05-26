package main

import (
	"log"

	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/config"
	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/di"
)

func main() {

	cfg, err := config.LoadEnvs()
	if err != nil {
		log.Fatalf("failed to load envs error:%s", err.Error())
	}
	server, err := di.InitializeService(cfg)
	if err != nil {
		log.Fatalf("failed to initialize service error:%s", err.Error())
	}
	server.Start()
}
