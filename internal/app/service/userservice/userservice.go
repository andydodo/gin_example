package userservice

import (
	"github.com/LIYINGZHEN/ginexample/internal/app/auth"
	"github.com/LIYINGZHEN/ginexample/internal/app/types"
	"github.com/pkg/errors"
)

type UserService struct {
	r types.UserRepository
	a Authenticator
}

type Authenticator interface {
	Hash(password string) (string, error)
	CompareHash(hashedPassword string, plainPassword string) error
}

// New returns the UserService.
func New(userRepository types.UserRepository) *UserService {
	return &UserService{
		r: userRepository,
		a: &auth.Authenticator{},
	}
}

func (uS *UserService) CreateUser(user *types.User, password string) (*types.User, error) {
	_, err := uS.r.FindByEmail(user.Email)
	if err == nil {
		return &types.User{}, errors.New("email already exists")
	}

	if len(password) < 8 {
		return &types.User{}, errors.New("password too short")
	}

	hashedPassword, err := uS.a.Hash(password)
	if err != nil {
		return &types.User{}, errors.Wrap(err, "error hashing password")
	}

	user.PasswordHash = hashedPassword

	err = uS.r.Store(user)
	if err != nil {
		return &types.User{}, errors.Wrap(err, "error storing user")
	}
	return user, nil
}

func (uS *UserService) Login(email string, password string) (*types.User, error) {
	user, err := uS.r.FindByEmail(email)
	if err != nil {
		return nil, errors.Wrap(err, "error finding user by email")
	}

	err = uS.a.CompareHash(user.PasswordHash, password)
	if err != nil {
		return nil, errors.Wrap(err, "error comparing hash")
	}

	err = uS.r.Update(user)
	if err != nil {
		return nil, errors.Wrap(err, "error updating sessionID")
	}

	return user, nil
}

func (uS *UserService) CheckAuthentication(sessionID string) (*types.User, error) {
	user, err := uS.r.FindBySessionID(sessionID)
	if err != nil {
		return nil, errors.Wrap(err, "error finding by sessionID")
	}

	return user, nil
}

func (uS *UserService) GetUser(id string) (*types.User, error) {
	return uS.r.Find(id)
}

func (uS *UserService) ChangePasswd(user *types.User, oldPw, newPw string) (*types.User, error) {
	err := uS.a.CompareHash(user.PasswordHash, oldPw)
	if err != nil {
		return nil, errors.Wrap(err, "error change password failed comparing hash")
	}

	user.PasswordHash, err = uS.a.Hash(newPw)
	if err != nil {
		return &types.User{}, errors.Wrap(err, "error change password failed hashing password")
	}
	err = uS.r.Update(user)
	if err != nil {
		return nil, errors.Wrap(err, "error change password failed updating user info")
	}

	return user, nil
}
