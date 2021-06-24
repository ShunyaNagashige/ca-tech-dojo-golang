package model_test

import (
	"database/sql"
	"os"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/ShunyaNagashige/ca-tech-dojo-golang/config"
	"github.com/ShunyaNagashige/ca-tech-dojo-golang/model"
	"github.com/google/uuid"
)

func TestMain(m *testing.M) {
	testDsn := createTestDsn()

	txdb.Register(config.Config.TestSqlDriver, config.Config.SqlDriver, testDsn)

	// 以下の２行がなぜ必要なのか，まだわかっていない
	// ただ，以下の２行を消すと，デバッグできなくなることは分かっている
	code := m.Run()
	os.Exit(code)
}

func createTestDsn() (testDsn string) {
	//テスト用のdsnを生成
	testDsn = strings.Replace(model.CreateDsn(),
		"/"+config.Config.DbName, "/"+config.Config.TestdbName, 1)

	return testDsn
}

// func openTestdb() (testDsn string){
// 	//テスト用のdsnを生成
// 	testDsn = strings.Replace(model.CreateDsn(),
// 		"/"+config.Config.DbName, "/"+config.Config.TestdbName, 1)

// 	var err error

// 	//DBに接続(DBのオープン)
// 	model.DbConn, err = sql.Open(config.Config.SqlDriver, testDsn)
// 	if err != nil {
// 		log.Fatalf("failed to open a database: %v", err)
// 	}

// 	return testDsn
// }

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
			dbConn, err := sql.Open("txdb", uuid.New().String())
			if err != nil {
				t.Fatalf("failed to Open (failed to start transaction): %v", err)
			}

			//トランザクション内で実行したクエリを全てロールバック
			defer dbConn.Close()

			if err := tt.u.CreateUser(dbConn); err != tt.err {
				t.Errorf("want %v, got %v", tt.err, err)
			}

			actual, err := model.GetUser(dbConn, tt.u.Token)
			if err != nil {
				t.Fatalf("failed to model.GetUser: %v", err)
			}

			if tt.u.User_name != actual.User_name {
				t.Errorf("want %#v, got %#v", *tt.u, *actual)
			}
		})
	}
}

func TestUser_UpdateUser(t *testing.T) {
	cases := []struct {
		name string
		u    *model.User
		err  error
	}{
		{
			name: "ok",
			u:    &model.User{User_name: "Taro",Token: "5555"},
			err:  nil,
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			dbConn, err := sql.Open("txdb", uuid.New().String())
			if err != nil {
				t.Fatalf("failed to Open (failed to start transaction): %v", err)
			}

			//トランザクション内で実行したクエリを全てロールバック
			defer dbConn.Close()

			if err := tt.u.UpdateUser(dbConn); err != tt.err {
				t.Errorf("want %v, got %v", tt.err, err)
			}

			actual, err := model.GetUser(dbConn, tt.u.Token)
			if err != nil {
				t.Fatalf("failed to model.GetUser: %v", err)
			}

			if tt.u.User_name != actual.User_name {
				t.Errorf("want %#v, got %#v", tt.u, actual)
			}
		})
	}
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
			dbConn, err := sql.Open("txdb", uuid.New().String())
			if err != nil {
				t.Fatalf("failed to Open (failed to start transaction): %v", err)
			}

			//トランザクション内で実行したクエリを全てロールバック
			defer dbConn.Close()

			actual, err := model.GetUser(dbConn, tt.token)

			if err != tt.err {
				t.Errorf("want %#v, got %#v", tt.err, err)
			}

			if *tt.expected != *actual {
				t.Errorf("want %#v, got %#v", *tt.expected, *actual)
			}
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

			dbConn, err := sql.Open("txdb", uuid.New().String())
			if err != nil {
				t.Fatalf("failed to Open (failed to start transaction): %v", err)
			}

			//トランザクション内で実行したクエリを全てロールバック
			defer dbConn.Close()

			actual, err := model.GetAllUser(dbConn)

			if err != tt.err {
				t.Errorf("want %#v, got %#v", tt.err, err)
			}

			if len(actual) != len(tt.expected) {
				t.Fatalf("want %#v, got %#v", tt.expected, actual)
			}

			for i := 0; i < len(tt.expected); i++ {
				if actual[i] != tt.expected[i] {
					t.Errorf("want %#v, got %#v", tt.expected[i], actual[i])
				}
			}
		})
	}
}
