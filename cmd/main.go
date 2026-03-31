package main

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/vic-eco/go_ecom_rest_api/internal/env"
	"log/slog"
	"os"
)

func main() {

	ctx := context.Background()

	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "host=localhost user=user password=pass dbname=ecom sslmode=disable"),
		},
	}

	//Structured Logging
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		slog.Error("database failed to connect", "error", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	slog.Info("database connected", "dsn", cfg.db.dsn)

	api := application{
		config: cfg,
		db:     conn,
	}

	err = api.run(api.mount())
	if err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}

}
