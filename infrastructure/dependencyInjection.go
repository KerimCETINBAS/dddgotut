package Infrastructure

import (
	Memory "github.com/kerimcetinbas/partizor.infrastructure/memory"
	"go.uber.org/dig"
)

func UseInfastructure(c *dig.Container) {

	c.Provide(Memory.UseUserRepository)
}
