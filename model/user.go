package model

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type User struct {
	UserId   int `json:"user_id"`
	UserName string `json:"user_name"`
	Token     string
}

func NewUser(userId int,userName,token string)*User{
	return &User{UserId: userId,UserName: userName,Token: token}
}

const (
	tokenLen = 10
)

func createToken() (string, error) {
	uuidObj, err := uuid.NewRandom()

	return uuidObj.String(), err

	// b := make([]byte, tokenLen)

	// //ランダムなバイト列の生成
	// if _, err := rand.Read(b); err != nil {
	// 	return "", errors.Wrap(err, "failed to rand.Read")
	// }

	// //stringにして返す
	// return *(*string)(unsafe.Pointer(&b)), nil
}

func (u *User) CreateUser(DbConn *sql.DB) (string,error) {
	var err error

	u.Token, err = createToken()
	if err != nil {
		return "",errors.Wrapf(err, "failed to createToken. u.Token=%s", u.Token)
	}

	cmd := fmt.Sprintf("INSERT INTO %s(user_name,token) VALUES(?,?)", tableNameUsers)
	if _, err := DbConn.Exec(cmd, u.UserName, u.Token); err != nil {
		return "",NewDbError(cmd, err)
	}

	return u.Token,nil
}

func GetUser(DbConn *sql.DB, token string) (*User, error) {
	var u User

	cmd := fmt.Sprintf("SELECT * FROM %s WHERE token = ?", tableNameUsers)
	row := DbConn.QueryRow(cmd, token)

	if err := row.Scan(&u.UserId, &u.UserName, &u.Token); err != nil {
		if err == sql.ErrNoRows {
			return nil, NewDbError(cmd, err)
		} else {
			return nil, errors.Wrapf(err, "failed to Scan.u=%#v", u)
		}
	}

	return &u, nil
}


func (u *User) UpdateUser(DbConn *sql.DB) error {
	cmd := fmt.Sprintf("UPDATE %s SET user_name = ?  WHERE token = ?", tableNameUsers)
	if _, err := DbConn.Exec(cmd, u.UserName,u.Token); err != nil {
		return NewDbError(cmd, err)
	}

	return nil
}

func GetAllUser(DbConn *sql.DB) ([]User, error) {
	cmd := fmt.Sprintf("SELECT * FROM %s", tableNameUsers)
	rows, err := DbConn.Query(cmd)
	if err != nil {
		return nil, NewDbError(cmd, err)
	}

	users := make([]User, 0)

	for rows.Next() {
		var u User

		rows.Scan(&u.UserId, &u.UserName, &u.Token)
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "failed to Scan.")
	}

	return users, nil
}
