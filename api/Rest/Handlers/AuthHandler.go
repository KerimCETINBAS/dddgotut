package Handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kerimcetinbas/partizor.contracts/Auth/Register"
	"github.com/kerimcetinbas/partizor.domain/Common/Services"
	"github.com/kerimcetinbas/partizor.domain/User"
	"go.uber.org/dig"
)

type authHandler struct {
	userService Services.IUserService
}

func newAuthHandler(userService Services.IUserService) *authHandler {
	return &authHandler{
		userService: userService,
	}
}

func authRouter(handler *authHandler, app *fiber.App) {

	r := app.Group("/auth")

	r.Post("/login", handler.Login)
	r.Post("/register", handler.Register)
}

func UseAuthHandler(c *dig.Container) {
	c.Provide(newAuthHandler)
	c.Invoke(authRouter)
}

func (h *authHandler) Login(ctx *fiber.Ctx) error {
	return nil
}

func (h *authHandler) Register(ctx *fiber.Ctx) error {

	body := Register.RegisterRequest{}

	if err := ctx.BodyParser(&body); err != nil {

		ctx.Status(fiber.StatusUnprocessableEntity)
		return ctx.JSON(
			fiber.Map{
				"Error": err.Error(),
			}, "application/problem+json")
	}

	h.userService.Create(User.User{
		Email:    body.Email,
		Password: body.Password,
	})
	return nil
}
