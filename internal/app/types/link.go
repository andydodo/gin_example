package types

import (
	"github.com/jinzhu/gorm"
)

type Link struct {
	gorm.Model
	//UserId   int
	UserName string
	Url      string `gorm:"not null"`
}

type LinkRepository interface {
	Store(*Link) error
	Find(string) (*Link, error)
	FindByUserName(string) (*Link, error)
	Update(*Link) error
	Delete(string) (*Link, error)
}

type LinkService interface {
	CreateLink(*Link, string) (*Link, error)
	GetLink(id string) (*Link, error)
	UpdateLink(*Link) error
	DeleteLink(id string) (*Link, error)
}
