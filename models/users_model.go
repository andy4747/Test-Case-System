package models

type Users struct {
	ID          uint64 `json:"id"`
	Username	string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	CreatedAt   string `json:"created_at"`
}

func GetUserByEmail(email string) (Users, error) {
	conn := Connect()
	defer conn.Close()
	query, err := conn.Query(`SELECT * FROM users WHERE email=$1`,email)
	if err != nil {
		return Users{}, err
	}
	defer query.Close()
	var user Users
	if query.Next() {
		err := query.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
		if err != nil {
			return Users{}, err
		}
	}
	return user, nil
}

func GetUserByID(id uint32) (Users, error) {
	conn := Connect()
	defer conn.Close()
	query, err := conn.Query(`SELECT * FROM users WHERE id=$1`,id)
	if err != nil {
		return Users{}, err
	}
	defer query.Close()
	var user Users
	if query.Next() {
		err := query.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
		if err != nil {
			return Users{}, err
		}
	}
	return user, nil
}
