package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	db,err := sql.Open("mysql","root:123456@tcp(www.wsj0819.cn:3306)/test")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	r,err := db.Exec("insert into num (num) values (?)",7)

	if err != nil {
		log.Fatalln(err)
	}

	if lastInsertId,err := r.LastInsertId();err == nil {
		//返回插入的id
		fmt.Println(lastInsertId)
	}

	if rowsAffecter,err := r.RowsAffected();err == nil {
		fmt.Println(rowsAffecter)
	}



}