package main

import (
	_ "github.com/onkiit/learn-microservices/user-service/controllers/auth"
	"github.com/onkiit/oserver"
	"github.com/onkiit/oserver/config"
)

var (
	cfg *config.Configuration
)

func Run() {
	oserver.Run()
}

func main() {
	Run()
}
