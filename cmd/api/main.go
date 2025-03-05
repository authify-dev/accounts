package main

import (
	"accounts/cmd/queue"
	"accounts/internal/api/server"
	"accounts/internal/core/settings"
	"fmt"
)

func main() {
	fmt.Println("accounts v0.0.1")

	settings.LoadDotEnv()

	settings.LoadEnvs()

	go queue.InitQueue()

	server.Run()
}
