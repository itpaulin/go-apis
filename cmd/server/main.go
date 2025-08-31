package main

import (
	"github.com/itpaulin/go-apis/configs"
)

func main() {
	config, err := configs.LoadConfig("./cmd/server")
	if err != nil {
		panic(err)
	}
	println(config.DBDriver)
}
