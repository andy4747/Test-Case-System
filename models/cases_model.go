package models

type TestCaseModel struct {
	ID             uint64	`json:"id"`
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


func GetCaseByTitle(title string) (TestCaseModel, error) {
	conn := Connect()
	defer conn.Close()
	sql := `SELECT * FROM test_case WHERE title=$1`
	rs, err := conn.Query(sql, title)
	if err != nil {
		return TestCaseModel{}, err
	}
	defer rs.Close()
	var tc TestCaseModel
	if rs.Next() {
		err := rs.Scan(&tc.ID,&tc.Title,&tc.Date,&tc.TestedBy,&tc.Functionality,&tc.Summary,&tc.Description,&tc.Data,&tc.URL,&tc.ExpectedResult,&tc.ActualResult,&tc.Environment,&tc.Device)
		if err != nil {
			return TestCaseModel{}, err
		}
	}
	return tc, nil
}

func GetCaseByID(id uint32) (TestCaseModel, error) {
	conn := Connect()
	defer conn.Close()
	sql := `SELECT * FROM test_case WHERE id=$1`
	rs, err := conn.Query(sql, id)
	if err != nil {
		return TestCaseModel{}, err
	}
	defer rs.Close()
	var tc TestCaseModel
	if rs.Next() {
		err := rs.Scan(&tc.ID,&tc.Title,&tc.Date,&tc.TestedBy,&tc.Functionality,&tc.Summary,&tc.Description,&tc.Data,&tc.URL,&tc.ExpectedResult,&tc.ActualResult,&tc.Environment,&tc.Device)
		if err != nil {
			return TestCaseModel{}, err
		}
	}
	return tc, nil
}