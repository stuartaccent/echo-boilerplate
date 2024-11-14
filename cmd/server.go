package cmd

import (
	"context"
	"echo.go.dev/pkg/auth"
	"echo.go.dev/pkg/home"
	"echo.go.dev/pkg/static"
	"echo.go.dev/pkg/transport/middleware"
	"echo.go.dev/pkg/transport/validate"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"time"
)

var cmdServer = &cobra.Command{
	Use:   "server",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		runServer()
	},
}

func initPool() *pgxpool.Pool {
	ctx := context.Background()
	dbPool, err := pgxpool.New(ctx, cfg.Database.URL().String())
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}

	if err := dbPool.Ping(ctx); err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}

	return dbPool
}

func runServer() {
	dbPool := initPool()
	defer dbPool.Close()

	engine := echo.New()
	engine.Debug = cfg.Server.Debug
	engine.Validator = validate.New()

	engine.Use(
		echomiddleware.Recover(),
		middleware.Logger(),
		middleware.SecureMiddleware(cfg.Security),
		middleware.CSRFMiddleware(cfg.Session),
		middleware.CORSMiddleware(cfg.Security),
		middleware.RateLimitMiddleware(),
		echomiddleware.Gzip(),
		middleware.SessionMiddleware(cfg.Session),
		middleware.Context(dbPool),
	)

	static.Router(engine)
	home.Router(engine)
	auth.Router(engine)

	server := http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Server.Port),
		ReadHeaderTimeout: 10 * time.Second,
		Handler:           engine,
	}

	log.Printf("Starting server on port %d", cfg.Server.Port)
	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}
}
