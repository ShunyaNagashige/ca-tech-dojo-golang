package controller

import (
	"bytes"
	"encoding/json"
	"errors"

	"log"
	"net/http"
	"regexp"

	"github.com/ShunyaNagashige/ca-tech-dojo-golang/model"
)

type JsonError struct{
	Err error `json:"error"`
}

type UserCreateRequest struct{
	Name string `json:"name"`
}

type UserCreateResponse struct{
	Token string `json:"token"`
}

func NewUserCreateRequest(name string)*UserCreateRequest{
	return &UserCreateRequest{Name:name}
}

func NewUserCreateResponse(token string)*UserCreateResponse{
	return &UserCreateResponse{Token:token}
}

func apiError(w http.ResponseWriter,err error,code int){
	w.WriteHeader(code)
	w.Header().Set("Content-Type","application/json; charset=utf-8")

	jsonErr:=&JsonError{err}
	var jsonBuf bytes.Buffer

	enc:=json.NewEncoder(&jsonBuf)
	if err:=enc.Encode(jsonErr);err!=nil{
		w.WriteHeader(http.StatusInternalServerError)

		// enc.Encodeでエラーが発生した場合はプログラムを終了させる
		// 次の場合を回避するため）　enc.Encodeでエラーが発生したときにはその旨をjsonで返す，なのでEncodeを実施するが
		// そのEncodeもエラーハンドリングしなければならない．．．以下ループ
		log.Fatalf("failed to enc.Encode: %v",err)
	}

	if _,err:=w.Write(jsonBuf.Bytes());err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("failed to ResponseWriter.Write: %v",err)
	}
}

var apiValidPath=regexp.MustCompile("^/((user/(create|get|update))|(gacha/draw)|(character/list))$")

func makeHandler(fn func(http.ResponseWriter,*http.Request))http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		m:=apiValidPath.FindStringSubmatch(r.URL.Path)
		if m==nil{
			apiError(w,errors.New("Not found"),http.StatusNotFound)
			return
		}

		fn(w,r)
	}
}

func checkMethod(r *http.Request,allowedMethods []string)error{
	for _,method:=range allowedMethods{
		if r.Method == method{
			return nil
		}
	}
	
	return errors.New("Method Not Allowed")
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	allowedMethods:=[]string{"HEAD","OPTIONS","POST"}

	if err:=checkMethod(r,allowedMethods);err!=nil{
		apiError(w,errors.New("Method Not Allowed"),http.StatusMethodNotAllowed)
		return
	}
	
	if r.Header.Get("Content-Type")!="application/json"{
		apiError(w,errors.New("Unsupported Media Type"),http.StatusUnsupportedMediaType)
		return
	}
	
	var ucreq UserCreateRequest
	dec:=json.NewDecoder(r.Body)

	if err:=dec.Decode(&ucreq);err!=nil{
		// リクエストボディのjson形式が正しくない場合と，dec.Decodeの処理が正しくない場合のどちらかが
		// 考えられるため，とりあえずステータスコード500で返す
		apiError(w,errors.New("Internal Server Error"),http.StatusInternalServerError)
		log.Fatalf("failed to dec.Decode: %v",err)
		return
	}

	u:=&model.User{UserName: ucreq.Name}
	var token string
	var err error

	if token,err=u.CreateUser(model.DbConn);err!=nil{
		apiError(w,errors.New("Internal Server Error"),http.StatusInternalServerError)
		log.Fatalf("failed to u.CreateUser: %v",err)
		return
	}

	if r.Method=="HEAD"{
		w.WriteHeader(http.StatusCreated)
	}

	w.Header().Set("Content-Type","application/json; charset=utf-8")

	ucres:=NewUserCreateResponse(token)
	var jsonBuf bytes.Buffer

	enc:=json.NewEncoder(&jsonBuf)
	if err:=enc.Encode(ucres);err!=nil{
		w.WriteHeader(http.StatusInternalServerError)

		// enc.Encodeでエラーが発生した場合はプログラムを終了させる
		// 次の場合を回避するため）　enc.Encodeでエラーが発生したときにはその旨をjsonで返す，なのでEncodeを実施するが
		// そのEncodeもエラーハンドリングしなければならない．．．以下ループ
		log.Fatalf("failed to enc.Encode: %v",err)
	}

	if _,err:=w.Write(jsonBuf.Bytes());err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("failed to ResponseWriter.Write: %v",err)
	}

	w.WriteHeader(http.StatusCreated)
}

func StartWebServer() {
	http.HandleFunc("/user/create", makeHandler(createUserHandler))
	http.ListenAndServe(":8080", nil)
}
