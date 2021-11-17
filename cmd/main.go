package main

import (
	"client/internal/sender"
)

func main() {
	var cmdService sender.IService = &sender.Service{}
	var cmd sender.ISender = cmdService.CreateCommander()
	cmd.ConnectToServer("localhost:8080")
}
