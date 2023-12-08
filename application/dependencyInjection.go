package Application

import (
	"github.com/kerimcetinbas/partizor.application/Services"
	"go.uber.org/dig"
)

func UseApplication(c *dig.Container) {

	c.Provide(Services.UserService)
}
