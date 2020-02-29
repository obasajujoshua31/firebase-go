package main

import (
	"firebase-go/server"
	"log"
)

func main() {

	err := server.StartServer()
	log.Println("Server starting ")
	if err != nil {
		log.Fatal(err)
	}
}
