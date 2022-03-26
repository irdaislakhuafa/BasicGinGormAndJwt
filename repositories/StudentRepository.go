package repositories

import (
	"github.com/irdaislakhuafa/BasicGinGormAndJwt/database"
	"github.com/irdaislakhuafa/BasicGinGormAndJwt/entities"
	"gorm.io/gorm"
)

type StudentRepository struct {
	db *gorm.DB
}

func (s *StudentRepository) SetDB(db *gorm.DB) {
	s.db = db
}

func (s *StudentRepository) FindAll() (listStudent []entities.Student) {
	database.GetDB().Find(&listStudent)
	return
}
func (s *StudentRepository) Save() {

}
