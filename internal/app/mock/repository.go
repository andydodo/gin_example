package mock

import "github.com/LIYINGZHEN/ginexample/internal/app/types"

type UserRepository struct {
	StoreFn        func(user *types.User) error
	StoreFnInvoked bool

	UpdateFn        func(user *types.User) error
	UpdateFnInvoked bool

	FindFn        func(id string) (*types.User, error)
	FindFnInvoked bool

	FindByEmailFn        func(email string) (*types.User, error)
	FindByEmailFnInvoked bool

	FindBySessionIDFn        func(sessionID string) (*types.User, error)
	FindBySessionIDFnInvoked bool
}

func (uRM *UserRepository) Store(user *types.User) error {
	uRM.StoreFnInvoked = true
	return uRM.StoreFn(user)
}

func (uRM *UserRepository) Update(user *types.User) error {
	uRM.UpdateFnInvoked = true
	return uRM.UpdateFn(user)
}

func (uRM *UserRepository) Find(id string) (*types.User, error) {
	uRM.FindFnInvoked = true
	return uRM.FindFn(id)
}

func (uRM *UserRepository) FindByEmail(email string) (*types.User, error) {
	uRM.FindByEmailFnInvoked = true
	return uRM.FindByEmailFn(email)
}

func (uRM *UserRepository) FindBySessionID(sessionID string) (*types.User, error) {
	uRM.FindBySessionIDFnInvoked = true
	return uRM.FindBySessionIDFn(sessionID)
}
