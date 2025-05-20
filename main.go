package main

import (
	"karting-challenge/config"
	"karting-challenge/router"
	"log"
)

func main() {
	config.LoadConfig()
	r := router.SetupRouter()
	log.Println("🚀 Server running on :8080")
	log.Fatal(r.Run(":8080"))
}
