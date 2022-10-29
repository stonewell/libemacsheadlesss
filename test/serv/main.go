package main

import (
	log "github.com/sirupsen/logrus"
	ehs "github.com/stonewell/emacsheadless/pkg/service"
)

func main() {
	log.SetLevel(log.DebugLevel)

	ehs.StartServer(&ehs.ServerConfig{
		Port:                3000,
		Addr:                "127.0.0.1",
	})
}
