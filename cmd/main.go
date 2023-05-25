package main

import (
	"log"

	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/config"
)

func main() {

	cfg, err := config.LoadEnvs()
	if err != nil {
		log.Fatalf("faild to load envs error:%s", err.Error())
	}

}
