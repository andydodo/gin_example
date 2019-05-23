package postgres

import (
	"errors"
	"fmt"

	"github.com/LIYINGZHEN/ginexample/internal/app/types"
	"github.com/jinzhu/gorm"
)

type LinkRepository struct {
	db *gorm.DB
}

func newLinkRepository(db *gorm.DB) *LinkRepository {
	return &LinkRepository{
		db: db,
	}
}

// Store creates a link record in the table
func (l *LinkRepository) Store(link *types.Link) error {
	return l.db.Create(link).Error
}

func (l *LinkRepository) Find(id string) (*types.Link, error) {
	var link types.Link

	db := l.db.Where("id = ?", id)
	err := first(db, &link)
	if err != nil {
		return nil, err
	}

	return &link, nil
}

func (l *LinkRepository) FindByUserName(userName string) (*types.Link, error) {
	if userName == "" {
		return &types.Link{}, errors.New("not found")
	}
	return l.findBy("username", userName)
}

func (l *LinkRepository) findBy(key, value string) (*types.Link, error) {
	link := types.Link{}

	db := l.db.Where(fmt.Sprintf("%s = ?", key), value)
	err := first(db, &link)

	return &link, err
}

func (l *LinkRepository) Update(link *types.Link) error {
	return l.db.Save(link).Error
}

func (l *LinkRepository) Delete(id string) error {
	var link types.Link

	err := l.db.Where("id = ?", id).Delete(&link).Error
	if err != nil {
		return err
	}
	return nil
}

func (l *LinkRepository) FindAll(links *[]types.Link) error {
	return l.db.Find(links).Error
}
