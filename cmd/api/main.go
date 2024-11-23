package main

import (
	"gopher-social/internal/env"
	"log"
)

func main() {
	// Load the configuration
	envVar := env.LoadEnvVar()

	port := envVar.PORT

	cfg := config{
		addr: port,
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
