package Api

import (
	"github.com/kerimcetinbas/partizor.api/Rest"
	"github.com/kerimcetinbas/partizor.api/Rest/Handlers"
	"go.uber.org/dig"
)

func UseRest(c *dig.Container) {

	c.Provide(Rest.Rest)
	Handlers.UseAuthHandler(c)
}
