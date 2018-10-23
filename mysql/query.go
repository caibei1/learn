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

	//ns := []num{}
	//rows,_ := db.Query("select * from num")
	////rows,_ := db.Query("select * from num")
	//for rows.Next(){
	//	n := num{}
	//	rows.Scan(&n.id,&n.num)
	//	ns = append(ns,n)
	//}
	//rows.Close()
	//fmt.Println(ns)

	row1 := db.QueryRow("select count(*) from num where num = ?",7)
	var i int
	err = row1.Scan(&i)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)

}
