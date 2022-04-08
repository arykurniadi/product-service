package repository

import (
	userInterface "dbo.id/product-service/app/user"
	"dbo.id/product-service/models"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	ConnDB *gorm.DB
}

func NewUserRepository(ConnDB *gorm.DB) userInterface.IUserRepository {
	return &UserRepository{ConnDB}
}

func (m *UserRepository) GetUserById(id int) (user models.User, err error) {
	tx := m.ConnDB.Begin()

	if err = tx.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	tx.Commit()
	return user, nil
}

func (m *UserRepository) GetListUser(page int, perPage int) ([]models.User, *models.Pagination, error) {
	tx := m.ConnDB

	var users []models.User

	if page < 1 {
		page = 1
	}

	if perPage < 1 {
		perPage = 25
	}

	offset := (page * perPage) - perPage
	_ = tx.Find(&users).Error
	total := len(users)

	err := tx.Limit(perPage).Offset(offset).Find(&users).Error
	if err != nil {
		return nil, nil, err
	}

	pagination := models.BuildPagination(total, page, perPage, len(users))

	return users, pagination, nil
}

func (m *UserRepository) Create(user models.User) (models.User, error) {
	tx := m.ConnDB.Begin()
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return user, err
	}

	tx.Commit()
	return user, nil
}

func (m *UserRepository) Update(id int, value models.User) (user models.User, err error) {
	tx := m.ConnDB.Begin()

	if err := tx.Model(&user).Where("id = ?", id).Updates(value).Error; err != nil {
		tx.Rollback()
		return user, err
	}

	tx.Commit()
	return user, nil
}

func (m *UserRepository) Delete(id int) (err error) {
	var user models.User

	tx := m.ConnDB.Begin()

	if err := tx.Where("id = ?", id).Delete(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
