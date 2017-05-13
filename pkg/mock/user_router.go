package mock

import "go_rest_api/pkg"

type UserService struct {
	CreateUserFn      func(u *root.User) error
	CreateUserInvoked bool

	GetUserByUsernameFn      func(username string) (error, root.User)
	GetUserByUsernameInvoked bool
}

func (us *UserService) CreateUser(u *root.User) error {
	us.CreateUserInvoked = true
	return us.CreateUserFn(u)
}

func (us *UserService) GetUserByUsername(username string) (error, root.User) {
	us.GetUserByUsernameInvoked = true
	return us.GetUserByUsernameFn(username)
}
