package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main()  {
	db,err := sql.Open("mysql","root:123456@tcp(www.wsj0819.cn:3306)/test")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	tx,_ := db.Begin()
	//db.Exec("insert into num (num) values (?)",777)  //不能用db调用  应该需要tx调用

	//Tx建立的时候就和一个连接绑定了，所以这些操作内部共用一个TX内部的连接。
	tx.Exec("insert into num (num) values (?)",777)
	tx.Rollback()
}
