package pet

import (
	"github.com/gin-gonic/gin"
	"github.com/ysdy/go_rest/db"
	"github.com/ysdy/go_rest/entity"
)

// Service procides pet's behavior
type Service struct{}

// User is alias of entity.Pet struct
type Pet entity.Pet

// GetAll is get all Pet
func (s Service) GetAll() ([]Pet, error) {
	db := db.GetDB()
	var u []Pet

	if err := db.Find(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

// CreateModel is create Pet model
func (s Service) CreateModel(c *gin.Context) (Pet, error) {
	db := db.GetDB()
	var u Pet

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	if err := db.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

// GetByID is get a User
func (s Service) GetByID(id string) (Pet, error) {
	db := db.GetDB()
	var u Pet

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}
