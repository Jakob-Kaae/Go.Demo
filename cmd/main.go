package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/Jakob-Kaae/Go.Demo/internal/env"
	"github.com/jackc/pgx/v5"
)

func main() {
	// Entry point of the application
	ctx := context.Background()
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})))

	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetString("DATABASE_URL", "host=localhost user=postgres password=postgres dbname=ecom sslmode=disable"),
		},
	}

	// database
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		slog.Error("Unable to connect to database", "error", err)
		os.Exit(1)
	}
	slog.Log(ctx, slog.LevelInfo, "Connected to DB", "cfg.db.dsn", cfg.db.dsn)
	defer conn.Close(ctx)

	api := application{
		config: cfg,
		db:     conn,
	}

	h := api.mount()
	if err := api.run(h); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}

}
