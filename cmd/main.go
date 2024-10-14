package main

import (
	"fib/internal/interface/http/server"
	"fib/internal/interface/http/server/config"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Print(err)
	}

	cfg := config.Config{
		Adr:    os.Getenv("ADR_PATH"),
		Token:  os.Getenv("API_TOKEN"),
		DbConn: os.Getenv("DB_CONNECTION"),
	}

	srv := server.NewServer(
		cfg.Adr,
		cfg.Token,
		cfg.DbConn,
	)

	if err := srv.Run(); err != nil {
		log.Print("run err")
	}

}
