package main

import (
	"context"
	"fmt"
	"go-with-fiber/internal/config"
	"go-with-fiber/internal/config/di"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	//loading cfg file i.e. env
	cfg := config.LoadConfig()

	//initialise container
	container := di.InitContainer(cfg)

	//fiber instance
	app := fiber.New()

	// Recover middleware (prevents server crash on panic)
	app.Use(recover.New())

	//register api routes
	api := app.Group("/api/v1")

	//register route
	container.RHandeler.RegisterRoute(api)

	//login route
	container.LHandler.LoginRoute(api)

	go gracefullShutdown(app)
	fmt.Println("Server started on :8080 ✅")
	if err := app.Listen(":8080"); err != nil {
		log.Fatal("failed to start server ❌:", err)
	}

}

func gracefullShutdown(app *fiber.App) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutdown signal receive")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	if err := app.ShutdownWithContext(ctx); err != nil {
		fmt.Println("failed to close the server gracefully ❌")
	}

	fmt.Println("server shutdown gracelly ✅")
}
