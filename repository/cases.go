package repository

import (
	"github.com/angeldhakal/testcase-ms/models"
	"gorm.io/gorm"
)

type CaseRepository interface {
	GetCase(int, string) (models.TestCaseModel, error)
	GetAllCases(string) ([]models.TestCaseModel, error)
	AddCase(models.TestCaseModel) (models.TestCaseModel, error)
	UpdateCase(models.TestCaseModel) (models.TestCaseModel, error)
	DeleteCase(int) (models.TestCaseModel, error)
}

type caseRepository struct {
	connection *gorm.DB
}

func NewCaseRepository() CaseRepository {
	return &caseRepository{
		connection: models.Connect(),
	}
}

func (db *caseRepository) GetCase(id int, user string) (testCase models.TestCaseModel, err error) {
	if err := db.connection.Where(&testCase, id).Error; err != nil {
		return testCase, err
	}
	return testCase, db.connection.Where("id=? and user=?", id, user).Error
}

func (db *caseRepository) GetAllCases(user string) (cases []models.TestCaseModel, err error) {
	return cases, db.connection.Where("user=?", user).Find(&cases).Error
}

func (db *caseRepository) AddCase(testCase models.TestCaseModel) (models.TestCaseModel, error) {
	return testCase, db.connection.Create(&testCase).Error
}

func (db *caseRepository) UpdateCase(testCase models.TestCaseModel) (models.TestCaseModel, error) {
	if err := db.connection.First(&testCase, testCase.ID).Error; err != nil {
		return testCase, err
	}
	return testCase, db.connection.Model(&testCase).Updates(&testCase).Error
}

func (db *caseRepository) DeleteCase(id int) (testCase models.TestCaseModel, err error) {
	if err := db.connection.First(&testCase, id).Error; err != nil {
		return testCase, err
	}
	return testCase, db.connection.Delete(&testCase).Error
}
