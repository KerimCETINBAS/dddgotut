package Memory

import (
	"fmt"
	"sync"

	"github.com/kerimcetinbas/partizor.domain/Common/Repositories"
	"github.com/kerimcetinbas/partizor.domain/User"
)

type userRepository struct {
	users []User.User
	sync.Mutex
}

func UseUserRepository() Repositories.IUserRepository {
	return &userRepository{
		users: make([]User.User, 0),
	}
}

func (r *userRepository) Create(user User.User) error {

	r.Lock()
	r.users = append(r.users, user)
	r.Unlock()
	fmt.Println(r.users)
	return nil
}
