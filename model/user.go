package model

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"unsafe"

	"github.com/pkg/errors"
)

type User struct {
	User_id   int
	User_name string
	Token     string
}

const (
	tokenLen = 10
)

func createToken() (string, error) {
	b := make([]byte, tokenLen)

	//ランダムなバイト列の生成
	if _, err := rand.Read(b); err != nil {
		return "", errors.Wrap(err,"failed to rand.Read")
	}

	//stringにして返す
	return *(*string)(unsafe.Pointer(&b)), nil
}

func (u *User) CreateUser() error {
	var err error

	u.Token, err = createToken()
	if err != nil {
		return errors.Wrapf(err,"failed to createToken. u.Token=%s",u.Token)
	}

	cmd := fmt.Sprintf("INSERT INTO %s(user_name,token) VALUES(?,?)", tableNameUsers)
	if _, err := dbConn.Exec(cmd, u.User_name, u.Token); err != nil {
		return NewDbError(cmd,err)
	}

	return nil
}

func (u *User) UpdateUser() error {
	cmd := fmt.Sprintf("UPDATE %s(user_name) VALUES(?,?)", tableNameUsers)
	if _, err := dbConn.Exec(cmd, u.User_name, u.Token); err != nil {
		return NewDbError(cmd,err)
	}

	return nil
}

func GetUser(token string) (*User, error) {
	var u User

	cmd := fmt.Sprintf("SELECT * FROM %s WHERE token=?", tableNameUsers)
	row:= dbConn.QueryRow(cmd, u.Token)
	
	if err:=row.Scan(&u.User_id,&u.User_name,&u.Token);err!=nil{
		if err==sql.ErrNoRows{
			return nil,NewDbError(cmd,err)
		}else{
			return nil,errors.Wrapf(err,"failed to Scan. u=%#v",u)
		}
	}

	return &u,nil
}

func GetAllUser() ([]User, error) {
	cmd := fmt.Sprintf("SELECT * FROM %s", tableNameUsers)
	rows, err := dbConn.Query(cmd)
	if err != nil {
		return nil, NewDbError(cmd,err)
	}

	users := make([]User, 0)

	for rows.Next() {
		var u User

		rows.Scan(&u.User_id, &u.User_name, &u.Token)
		users = append(users, u)
	}

	if err:=rows.Err();err!=nil{
		return nil,errors.Wrap(err,"failed to Scan.")
	}

	return users, nil
}
