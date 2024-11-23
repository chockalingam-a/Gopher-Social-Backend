package main

import (
	"gopher-social/internal/db"
	"gopher-social/internal/env"
	"gopher-social/internal/store"
	"log"
)

func main() {
	// Load the configuration
	envVar := env.LoadEnvVar()

	cfg := config{
		addr: envVar.ADDR,
		db: dbConfig{
			addr:         envVar.DB_ADDR,
			maxOpenConns: envVar.DB_MAX_OPEN_CONNS,
			maxIdleConns: envVar.DB_MAX_IDLE_CONNS,
			maxIdleTime:  envVar.DB_MAX_IDLE_TIME,
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	log.Println("Database connection pool established")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
