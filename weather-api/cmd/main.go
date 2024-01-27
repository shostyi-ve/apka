package main

import (
	"errors"
	"io/fs"
	"log"

	"github.com/joho/godotenv"
	"github.com/shostyi-ve/apka/weather/container"
)

func main() {
	if err := godotenv.Load("./config/.env"); err != nil && !errors.Is(err, fs.ErrNotExist) {
		log.Fatal(err)
	}

	container.Build().Run()
}
