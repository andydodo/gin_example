package types

import (
	"github.com/jinzhu/gorm"
)

type Link struct {
	gorm.Model
	//UserId   int
	Name string `gorm:"column:name" form:"name" json:"name"`
	Url  string `gorm:"column:url" form:"url" json:"url" "not null"`
}

type LinkRepository interface {
	Store(*Link) error
	Find(string) (*Link, error)
	FindByName(string) (*Link, error)
	Update(*Link) error
	Delete(string) error
	FindAll() ([]Link, error)
}

type LinkService interface {
	CreateLink(*Link, string) (*Link, error)
	GetLink(string) (*Link, error)
	UpdateLink(*Link) error
	DeleteLink(string) error
	GetAllLink() ([]Link, error)
}
