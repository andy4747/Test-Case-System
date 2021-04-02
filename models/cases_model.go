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