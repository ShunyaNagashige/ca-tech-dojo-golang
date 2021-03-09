package main

import (
	"flag"
	"fmt"
	"strings"
)

//設定される変数のポインタを取得
var msg=flag.String("msg","デフォルト値","説明")
var n int
func init(){
	//ポインタを指定して設定を予約
	flag.IntVar(&n,"n",1,"回数")
}

// func f()string{
// 	return "errorです"
// }

// func readFile(fn string)error{
// 	f,err:=os.Open(fn)
// 	if err!=nil{
// 		return err
// 	}
// 	defer f.Close()
// 	return nil
// }

func main(){
	//forの中でのdeferを避けるために
	//forの中を関数に分ける
	

	// // //標準エラー出力に表示
	// fmt.Fprintf(os.Stderr,"エラー\n")
	// os.Exit(1)
	//ここで実際にフラグが設定される
	flag.Parse()
	fmt.Println(flag.Args())
	fmt.Println(strings.Repeat(*msg,n))
}