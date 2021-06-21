package model_test

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"

	"github.com/ShunyaNagashige/ca-tech-dojo-golang/config"
	"github.com/ShunyaNagashige/ca-tech-dojo-golang/model"
)

func openTestdb(t *testing.T) {
	t.Helper()

	//テスト用のdsnを生成
	testDsn := strings.Replace(model.CreateDsn(),
		"/"+config.Config.DbName, "/"+config.Config.TestdbName, 1)

	fmt.Println(model.CreateDsn())
	fmt.Println(testDsn)

	var err error
	
	//DBに接続(DBのオープン)
	model.DbConn, err = sql.Open(config.Config.SqlDriver, testDsn)
	if err != nil {
		t.Fatalf("failed to open a database: %v", err)
	}
}

func isUserEqual(t *testing.T, expected, result *model.User) bool {
	t.Helper()

	if result.User_id != expected.User_id {
		return false
	}

	if result.User_name != expected.User_name {
		return false
	}

	if result.Token != expected.Token {
		return false
	}

	return true
}

func TestGetUser(t *testing.T) {
	cases := []struct {
		name     string
		token    string
		expected *model.User
		err      error
	}{
		{
			name:     "ok",
			token:    "5555",
			expected: &model.User{User_id: 1, User_name: "test", Token: "5555"},
			err:      nil,
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			openTestdb(t)

			actual, err := model.GetUser(tt.token)

			if err != tt.err {
				t.Errorf("want %#v, got %#v", tt.err, err)
			}

			if *tt.expected != *actual {
				t.Errorf("want %#v, got %#v", *tt.expected, *actual)
			}
		})
	}
}

func TestUser_CreateUser(t *testing.T) {
	cases := []struct {
		name string
		u    *model.User
		err  error
	}{
		{
			name: "ok",
			u:    &model.User{User_name: "Taro"},
			err:  nil,
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			openTestdb(t)

			//トランザクションを貼って，テストが終わればロールバックする
			//これにより，このテストが他のテストに影響を及ぼさないようにする．
			tx, err := model.DbConn.Begin()
			if err != nil {
				t.Fatalf("failed to DbConn.Begin: %v", err)
			}

			if err := tt.u.CreateUser(); err != tt.err {
				t.Errorf("want %v, got %v", tt.err, err)
			}

			actual, err := model.GetUser(tt.u.Token)
			if err != nil {
				t.Fatalf("failed to model.GetUser: %v", err)
			}

			if tt.u.User_name != actual.User_name {
				t.Errorf("want %#v, got %#v", *tt.u, *actual)
			}

			tx.Rollback()
		})
	}
}

func TestGetAllUser(t *testing.T) {
	cases := []struct {
		name     string
		expected []model.User
		err      error
	}{
		{
			name: "ok",
			expected: []model.User{
				{User_id: 1, User_name: "test", Token: "5555"},
			},
			err: nil,
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			openTestdb(t)

			actual, err := model.GetAllUser()

			if err != tt.err {
				t.Errorf("want %#v, got %#v", tt.err, err)
			}

			if len(actual) != len(tt.expected) {
				t.Errorf("len(actual)=%d is different from len(tt.expected)=%d", len(actual), len(tt.expected))
			}

			for i := 0; i < len(tt.expected); i++ {
				if actual[i] != tt.expected[i] {
					t.Errorf("want %#v, got %#v", tt.expected[i], actual[i])
				}
			}
		})
	}
}
