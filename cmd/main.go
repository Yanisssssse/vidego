package main

import (
	"fmt"
	"os"

	"github.com/Yanisssssse/vidego/internal/server"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	c := server.NewConfig(os.Getenv("HOST"), os.Getenv("PORT"), false)
	s := server.NewServer(c)

	fmt.Println("Listening on http://" + os.Getenv("HOST") + ":" + os.Getenv("PORT"))
	err = s.Serve()
	fmt.Println(err)
}
