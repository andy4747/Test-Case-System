package models

import "gorm.io/gorm"

type TestCaseModel struct {
	gorm.Model
	Title          string	`json:"title"`
	Date           string	`json:"date"`
	TestedBy       string	`json:"tested_by"`
	Functionality  string	`json:"functionality"`
	Summary        string	`json:"summary"`
	Description    string	`json:"description"`
	Data           string	`json:"data"`
	URL            string	`json:"url"`
	ExpectedResult string	`json:"expected_result"`
	ActualResult   string	`json:"actual_result"`
	Environment    string	`json:"environment"`
	Device         string	`json:"device"`
}

//TableName for the TestCaseModel
func (TestCaseModel) TableName() string {
	return "test_cases"
}