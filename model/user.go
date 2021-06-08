package model

import (
	"crypto/rand"
	"fmt"
	"unsafe"
)

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

const (
	tokenLen = 10
)

func createToken() (string, error) {
	b := make([]byte, tokenLen)

	//ランダムなバイト列の生成
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	//stringにして返す
	return *(*string)(unsafe.Pointer(&b)), nil
}

func (u *User) Create() error {
	var err error

	u.Token, err = createToken()
	if err != nil {
		return &UserError{u, err}
	}

	cmd := fmt.Sprintf("INSERT INTO %s(user_name,token) VALUES(?,?)", tableNameUsers)
	if _, err := dbConn.Exec(cmd, u.User_name, u.Token); err != nil {
		return &DbError{cmd, err}
	}

	return nil
}

func (u *User) Update() error {
	cmd := fmt.Sprintf("UPDATE %s(user_name) VALUES(?,?)", tableNameUsers)
	if _, err := dbConn.Exec(cmd, u.User_name, u.Token); err != nil {
		return &DbError{cmd, err}
	}

	return nil
}

func Get(token string) (*User, error) {
	var u *User

	cmd := fmt.Sprintf("SELECT * FROM %s WHERE token=?", tableNameUsers)
	row:= dbConn.QueryRow(cmd, u.Token)
	
	if err:=row.Scan(&u.User_id,&u.User_name,&u.Token){
		return nil,&UserError{}
	}

	
}

func GetAll() ([]User, error) {
	cmd := fmt.Sprintf("SELECT * FROM %s", tableNameUsers)
	rows, err := dbConn.Query(cmd)
	if err != nil {
		return nil, &DbError{Cmd: cmd, Err: err}
	}

	users := make([]User, 0)

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.User_id, &u.User_name, &u.Token); err != nil {
			return nil, &DbError{cmd,err}
		}
		users = append(users, u)
	}

	return users, nil
}
