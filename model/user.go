package model

import "fmt"

type User struct {
	User_id   int
	User_name string
	Token     string
}

type UserError struct {
	U   *User
	Err error
}

func (err *UserError) Error() string {
	return err.Err.Error()
}

func GetAllUser() ([]User, error) {
	cmd := fmt.Sprintf("SELECT * FROM %s", tableNameUsers)
	rows, err := dbConn.Query(cmd)
	if err != nil {
		return nil, &DbError{Cmd: cmd, Err: err}
	}

	users := make([]User, 0)

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.User_id, &u.User_name, &u.Token); err != nil {
			return nil, &UserError{U: &u, Err: err}
		}
		users = append(users, u)
	}

	return users, nil
}
