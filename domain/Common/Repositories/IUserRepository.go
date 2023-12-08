package Repositories

import "github.com/kerimcetinbas/partizor.domain/User"

type IUserRepository interface {
	Create(user User.User) error
}
