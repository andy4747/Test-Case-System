package repository

import (
	"github.com/angeldhakal/testcase-ms/models"
	"gorm.io/gorm"
)

type CaseRepository interface {
    GetCase(int) (models.TestCaseModel, error)
    GetAllCases() ([]models.TestCaseModel, error)
    AddCase(models.TestCaseModel) (models.TestCaseModel, error)
    UpdateCase(models.TestCaseModel) (models.TestCaseModel, error)
    DeleteCase(models.TestCaseModel) (models.TestCaseModel, error)
}

type caseRepository struct {
    connection *gorm.DB
}

func NewCaseRepository() *caseRepository {
    return &caseRepository{
        connection: models.Connect(),
    }
}


func (db *caseRepository) GetCase(id int) (models.TestCaseModel, error) {return models.TestCaseModel{}, nil}

func (db *caseRepository) GetAllCases() ([]models.TestCaseModel, error)  {return []models.TestCaseModel{}, nil}

func (db *caseRepository) AddCase(testCase models.TestCaseModel) (models.TestCaseModel, error) {return models.TestCaseModel{}, nil}

func (db *caseRepository) UpdateCase(testCase models.TestCaseModel) (models.TestCaseModel, error)  {return models.TestCaseModel{}, nil}

func (db *caseRepository) DeleteCase(testCase models.TestCaseModel) (models.TestCaseModel, error)  {return models.TestCaseModel{}, nil}

