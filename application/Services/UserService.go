package Services

import (
	"github.com/kerimcetinbas/partizor.domain/Common/Repositories"
	"github.com/kerimcetinbas/partizor.domain/Common/Services"
	"github.com/kerimcetinbas/partizor.domain/User"
)

type userService struct {
	userRepository Repositories.IUserRepository
}

func UserService(userRepository Repositories.IUserRepository) Services.IUserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) Create(user User.User) error {

	s.userRepository.Create(user)
	return nil
}
