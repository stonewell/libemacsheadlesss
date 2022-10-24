package main

import (
	ehs "github.com/stonewell/emacsheadless/pkg/service"
)

func main() {
	ehs.StartServer(&ehs.ServerConfig{
		Port:                3000,
		Addr:                "127.0.0.1",
	})
}
