package main

import (
	"karting-challenge/internal/config"
	"karting-challenge/internal/router"
	"log"
)

func main() {
	config.LoadConfig()
	r := router.SetupRouter()
	log.Println("🚀 Server running on :8080")
	log.Fatal(r.Run(":8080"))
}
