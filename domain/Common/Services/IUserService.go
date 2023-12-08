package Services

import "github.com/kerimcetinbas/partizor.domain/User"

type IUserService interface {
	Create(user User.User) error
}
