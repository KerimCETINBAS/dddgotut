package main

import (
	"github.com/gofiber/fiber/v2"
	Api "github.com/kerimcetinbas/partizor.api"
	Application "github.com/kerimcetinbas/partizor.application"
	Infrastructure "github.com/kerimcetinbas/partizor.infrastructure"
	"go.uber.org/dig"
)

func main() {

	container := dig.New()

	Application.UseApplication(container)
	Infrastructure.UseInfastructure(container)
	Api.UseRest(container)

	container.Invoke(func(app *fiber.App) {
		app.Listen(":8080")
	})
}
