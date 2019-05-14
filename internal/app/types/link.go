package types

import (
	"github.com/jinzhu/gorm"
)

type Link struct {
	gorm.Model
	//UserId   int
	UserName string `gorm:"column:username" form:"username" json:"username"`
	Url      string `gorm:"column:url" form:"url" json:"url" "not null"`
}

type LinkRepository interface {
	Store(*Link) error
	Find(string) (*Link, error)
	FindByUserName(string) (*Link, error)
	Update(username string, url string) error
	Delete(string) (*Link, error)
}

type LinkService interface {
	CreateLink(*Link, string) (*Link, error)
	GetLink(id string) (*Link, error)
	UpdateLink(username string, url string) error
	DeleteLink(id string) (*Link, error)
}
