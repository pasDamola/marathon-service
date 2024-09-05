package main

import "log"

func main() {
	log.Println("Starting Runners App")
	log.Println("Initializing Configuration")
	config := config.InitConfig("runners")

	log.Println("Initializing database")
	dbHandler := server.InitDatabase(config)

	log.Println("Initializing HTTP server")
	httpServer := server.InitHttpServer(config, dbHandler)

	httpServer.Start()
}
