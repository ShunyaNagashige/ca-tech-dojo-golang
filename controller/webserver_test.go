package controller_test

import (
	"net/http/httptest"
	"testing"
)

func TestUser_CreateUser(t *testing.T){
	w:=httptest.NewRecorder()
	r:=httptest.NewRequest("POST","/user/create",nil)

}