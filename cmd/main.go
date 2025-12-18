package main

import (
	"fmt"
	"os"

	"github.com/Yanisssssse/vidego/internal/server"
	"github.com/Yanisssssse/vidego/internal/storage"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading env file: " + err.Error())
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	updir := os.Getenv("UPDIR")

	if host == "" {
		panic("Missing required environment variable: HOST")
	}
	if port == "" {
		panic("Missing required environment variable: PORT")
	}
	if updir == "" {
		panic("Missing required environment variable: UPDIR")
	}

	s, err := storage.NewLocalStorage(updir)
	if err != nil {
		panic("Failed to initialize storage: " + err.Error())
	}

	c := server.NewConfig(host, port, false, s)
	sv := server.NewServer(c)

	fmt.Println("Listening on http://" + host + ":" + port)
	err = sv.Serve()
	fmt.Println(err)
}
