package mock

import "github.com/LIYINGZHEN/ginexample"

type UserRepository struct {
	StoreFn        func(user *ginexample.User) error
	StoreFnInvoked bool

	UpdateFn        func(user *ginexample.User) error
	UpdateFnInvoked bool

	FindFn        func(id string) (*ginexample.User, error)
	FindFnInvoked bool

	FindByEmailFn        func(email string) (*ginexample.User, error)
	FindByEmailFnInvoked bool

	FindBySessionIDFn        func(sessionID string) (*ginexample.User, error)
	FindBySessionIDFnInvoked bool
}

func (uRM *UserRepository) Store(user *ginexample.User) error {
	uRM.StoreFnInvoked = true
	return uRM.StoreFn(user)
}

func (uRM *UserRepository) Update(user *ginexample.User) error {
	uRM.UpdateFnInvoked = true
	return uRM.UpdateFn(user)
}

func (uRM *UserRepository) Find(id string) (*ginexample.User, error) {
	uRM.FindFnInvoked = true
	return uRM.FindFn(id)
}

func (uRM *UserRepository) FindByEmail(email string) (*ginexample.User, error) {
	uRM.FindByEmailFnInvoked = true
	return uRM.FindByEmailFn(email)
}

func (uRM *UserRepository) FindBySessionID(sessionID string) (*ginexample.User, error) {
	uRM.FindBySessionIDFnInvoked = true
	return uRM.FindBySessionIDFn(sessionID)
}
