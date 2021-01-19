package main

import (
	"ci-playground/server"
)

func main() {
	err := server.Init()

	if err != nil {
		panic("The server failed to start up")
	}
}
