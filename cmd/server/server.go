package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/auth_service/internal/apiserver"
	"github.com/auth_service/internal/config"
	"github.com/auth_service/internal/store"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", "configs/config.toml", "Path to server config")
}

func main() {
	flag.Parse()

	c := config.NewConfig()
	if _, err := toml.DecodeFile(configPath, c); err != nil {
		log.Fatal(err)
	}

	stroe := store.New(c.Store)

	server := apiserver.New(c.Server, stroe)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
