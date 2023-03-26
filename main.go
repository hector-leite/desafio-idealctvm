package main

import (
	application "github.com/hector-leite/desafio-idealctvm/application"
	server "github.com/hector-leite/desafio-idealctvm/server"
)

func main() {
	server.Run(application.BuildApp())
}
