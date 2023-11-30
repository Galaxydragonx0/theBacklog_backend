/*
	TODO:
	Setup env with variables needed
		- api keys
		- database conn string
		- api connection strings



	Write handlers to handle the logic of endpoints

*/

package main

import (
	"fmt"
	"log"
	"os"
	"theBacklog/backend/api"
)

func main() {

	//Server - of type api.Server
	server := api.NewServer()

	err := server.Start()

	if err != nil {
		log.Fatal("The server could not be started:", err)
	}

	fmt.Println("running the server on http:localhost:", os.Getenv("PORT"))
}
