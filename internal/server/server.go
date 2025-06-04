package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/AD12-codes/type-ninjas/db"
	"github.com/AD12-codes/type-ninjas/internal/users"
	"github.com/AD12-codes/type-ninjas/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Uptime    string    `json:"uptime"`
	Version   string    `json:"version"`
	Hostname  string    `json:"hostname"`
	GoVersion string    `json:"goVersion"`
	Env       string    `json:"env"`
	DBStatus  string    `json:"dbStatus"`
}

var startTime = time.Now()

func Run() {
	// context
	ctx := context.Background()

	// database connection
	pool := db.DbConnection(ctx)

	// sqlc queries setup
	queries := db.New(pool)

	// echo server
	e := echo.New()

	// CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// validation middleware
	e.Validator = utils.NewValidator()

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

	v1 := e.Group("/api/v1")

	// all packages route registrations
	users.RegisterRoutes(v1, queries)

	// start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "7070"
	}

	if err := e.Start(":" + port); err != nil {
		log.Fatal(err)
	}
}
