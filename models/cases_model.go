package models

type TestCaseModel struct {
	ID             uint64
	Title          string
	Date           string
	TestedBy       string
	Functionality  string
	Summary        string
	Description    string
	Data           string
	URL            string
	ExpectedResult string
	ActualResult   string
	Environment    string
	Device         string
}


