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
func (s *StudentRepository) Save(entity interface{}) (student entities.Student, err error) {
	err = database.GetDB().Save(entity).Error
	student, _ = s.FindById(int(entity.(*entities.Student).ID))
	return student, err
}
func (s *StudentRepository) FindById(id int) (student entities.Student, err error) {
	err = database.GetDB().Where("id = ?", id).Find(&student).Error
	return
}
