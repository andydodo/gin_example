package postgres

import (
	"errors"
	"fmt"

	"github.com/LIYINGZHEN/ginexample/internal/app/types"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func newUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Stroe creates a user record in the table
func (u *UserRepository) Store(user *types.User) error {
	return u.db.Create(user).Error
}

func (u *UserRepository) Find(id string) (*types.User, error) {
	var user types.User

	db := u.db.Where("id = ?", id)
	err := first(db, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) FindByEmail(email string) (*types.User, error) {
	if email == "" {
		return &types.User{}, errors.New("not found")
	}
	return u.findBy("email", email)
}

func (u *UserRepository) FindBySessionID(sessionID string) (*types.User, error) {
	if sessionID == "" {
		return nil, errors.New("not found")
	}
	return u.findBy("session_id", sessionID)
}

func (u *UserRepository) findBy(key string, value string) (*types.User, error) {
	user := types.User{}

	db := u.db.Where(fmt.Sprintf("%s = ?", key), value)
	err := first(db, &user)

	return &user, err
}

func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return errors.New("resource not found")
	}
	return err
}

func (u *UserRepository) Update(user *types.User) error {
	return u.db.Save(user).Error
}
