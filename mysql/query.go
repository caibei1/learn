package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type num struct {
	id int
	num int
}


func main()  {
	db,err := sql.Open("mysql","root:123456@tcp(www.wsj0819.cn:3306)/test")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	ns := []num{}
	rows,_ := db.Query("select * from num where id = ?",`1' or '1' = '1`)
	//rows,_ := db.Query("select * from num")
	for rows.Next(){
		n := num{}
		rows.Scan(&n.id,&n.num)
		ns = append(ns,n)
	}
	rows.Close()
	fmt.Println(ns)
}
