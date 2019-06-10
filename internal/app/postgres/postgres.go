package postgres

import (
	"fmt"

	"github.com/LIYINGZHEN/ginexample/internal/app/types"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DBConfig contains the environment varialbes requirements to initialize postgres.
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func (c DBConfig) connectionInfo() string {
	if c.Password == "" {
		return fmt.Sprintf("host=%s port=%s user=%s dbname=%s "+
			"sslmode=disable", c.Host, c.Port, c.User, c.Name)
	}
	return fmt.Sprintf("host=%s port=%s user=%s password=%s "+
		"dbname=%s sslmode=disable", c.Host, c.Port, c.User,
		c.Password, c.Name)
}

// Repository contains information for every repositories.
type Repository struct {
	UserRepository *UserRepository
	LinkRepository *LinkRepository
	DB             *gorm.DB
}

// Initialize the postgres database.
func Initialize(c DBConfig) *Repository {
	fmt.Printf("c.connectionInfo() %+v", c.connectionInfo())
	db, err := gorm.Open("postgres", c.connectionInfo())
	if err != nil {
		panic(err)
	}

	return &Repository{
		UserRepository: newUserRepository(db),
		LinkRepository: newLinkRepository(db),
		DB:             db,
	}
}

// AutoMigrate will attempt to automatically migrate all tables
func (r *Repository) AutoMigrate() error {
	err := r.DB.AutoMigrate(&types.User{}, &types.Link{}).Error
	if err != nil {
		return err
	}
	return nil
}
