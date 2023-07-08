package main

import (
	"context"
	"fmt"
	"fx-postgres/config"
	"fx-postgres/db"
	"fx-postgres/handler"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func NewFiberServer(lc fx.Lifecycle, handler *handler.Handler) *fiber.App {
	app := fiber.New()

	app.Get("/hello", handler.Hello)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("Starting fiber server on port 8080")
			go app.Listen(":8080")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})

	return app
}

func main() {
	app := fx.New(
		config.EnvModule,
		db.DatabaseModule,
		handler.HandlerModule,
		fx.Invoke(NewFiberServer),
	)

	app.Run()
}
