package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	// "github.com/AD12-codes/go-template/internal/users"
	"github.com/AD12-codes/go-template/db"
	"github.com/labstack/echo/v4"
)

type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Uptime    string    `json:"uptime"`
	Version   string    `json:"version"`
	Hostname  string    `json:"hostname"`
	GoVersion string    `json:"go_version"`
	Env       string    `json:"env"`
	DBStatus  string    `json:"db_status"`
}

var startTime = time.Now()

func Run() {
	// context
	ctx := context.Background()

	// database connection
	pool := db.DbConnection(ctx)

	// sqlc queries setup
	// queries := db.New(pool)

	// echo server
	e := echo.New()

	// health route
	e.GET("/health", func(ec echo.Context) error {
		hostname, _ := os.Hostname()

		resp := HealthResponse{
			Status:    "OK",
			Timestamp: time.Now(),
			Uptime:    time.Since(startTime).String(),
			Version:   "1.0.0",
			Hostname:  hostname,
			GoVersion: runtime.Version(),
			Env:       os.Getenv("APP_ENV"),
			DBStatus: func() string {
				dbCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
				defer cancel()
				if err := pool.Ping(dbCtx); err != nil {
					return "DOWN"
				}
				return "UP"
			}(),
		}
		return ec.JSON(http.StatusOK, resp)
	})

	// all packages route registrations
	// users.RegisterRoutes(e, queries)

	// start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "7070"
	}

	if err := e.Start(":" + port); err != nil {
		log.Fatal(err)
	}
}
