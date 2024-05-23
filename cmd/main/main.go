package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	"github.com/diplom-pam/edu/internal/application"
	"github.com/diplom-pam/edu/internal/config"
	"github.com/diplom-pam/edu/internal/logger"
	"github.com/diplom-pam/edu/internal/storage"
)

var (
	version = "undefined"
)

func main() {
	cfg := config.MustLoad()

	logger.Init(cfg.Debug)

	errRun := run(cfg)
	if errRun != nil {
		log.Printf("run error: %v", errRun)
		os.Exit(1)
	}
}

func run(cfg *config.Config) error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	logger.Info("start", zap.String("version", version))

	ln, errLn := net.Listen("tcp", cfg.Address)
	if errLn != nil {
		return fmt.Errorf("listen error: %w", errLn)
	}
	defer ln.Close()

	PgPool, errDBMain := pgxpool.New(ctx, buildPGConnectionString(cfg.Pg))
	if errDBMain != nil {
		return fmt.Errorf("db main error: %w", errDBMain)
	}
	defer PgPool.Close()

	store := storage.New(PgPool)

	wg := sync.WaitGroup{}

	app := application.New(store)
	wg.Add(1)
	go app.Run(ctx, cancel, &wg, ln)

	<-ctx.Done()

	wg.Wait()

	logger.Info("stop")

	return nil
}

func buildPGConnectionString(cfg config.Postgres) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s&sslrootcert=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.SSLMode,
		cfg.SSLCertPath)
}
