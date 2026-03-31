package main

import (
	"log"
	"os"
)

func main() {
	cfg := config{
		addr: ":8080",
		db:   dbConfig{},
	}
	api := application{
		config: cfg,
	}

	err := api.run(api.mount())
	if err != nil {
		log.Printf("server failed to start, err %s", err)
		os.Exit(1)
	}

}
