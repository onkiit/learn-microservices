package main

import (
	"github.com/onkiit/oserver"
	"github.com/onkiit/oserver/config"
)

var (
	cfg *config.Configuration
)

func Run() {
	oserver.RunWithConfigs()
}

func main() {
	Run()
}
