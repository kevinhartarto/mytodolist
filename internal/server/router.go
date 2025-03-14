package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kevinhartarto/mytodolist/internal/controllers"
	"github.com/kevinhartarto/mytodolist/internal/database"
	"github.com/kevinhartarto/mytodolist/internal/utils"
)

func NewHandler(db database.Service) *fiber.App {

	app := fiber.New()
	app.Use(healthcheck.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	// API groups
	api := app.Group("/api")

	// Version 1
	v1 := api.Group("/v1")

	// Tasks
	list := controllers.NewTaskController(db)
	listAPI := v1.Group("/list")

	listAPI.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
	listAPI.Get("/tasks", func(c *fiber.Ctx) error {
		return list.GetTasks(c)
	})
	listAPI.Get("/task/:uuid", func(c *fiber.Ctx) error {
		uuid := utils.ParseUUID(c.Params("uuid"))
		return list.GetTaskByUuid(uuid, c)
	})
	listAPI.Put("/task", func(c *fiber.Ctx) error {
		return list.UpdateTask(c)
	})
	listAPI.Delete("/task", func(c *fiber.Ctx) error {
		return list.TaskFinished(c)
	})

	return app
}
