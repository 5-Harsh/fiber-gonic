package server

import (
	"github.com/gofiber/fiber/v2"

	"fiber-mongo-rest/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "fiber-mongo-rest",
			AppName:      "fiber-mongo-rest",
		}),

		db: database.New(),
	}

	return server
}
