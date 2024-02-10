package initializers

import (
	"os"

	"github.com/ToffaKrtek/self-executor/controllers"
	"github.com/gofiber/fiber/v2"
)

var APP *fiber.App

func Run() {
	APP = fiber.New()
	routes()
	APP.Listen(":" + os.Getenv("PORT"))
}

func routes() {
	APP.Get("/", controllers.Index)
}

func Static() {
	APP.Static("/", "./public")
}
