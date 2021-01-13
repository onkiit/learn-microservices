package main

import (
	_ "github.com/onkiit/learn-microservices/user-service/controllers/auth"
	sapp "github.com/onkiit/sapp"
	"github.com/onkiit/sapp/config"
)

var (
	cfg *config.Configuration
)

func Run() {
	sapp.RunServer()
}

func main() {
	Run()
}
