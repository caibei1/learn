package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main()  {
	db,err := sql.Open("mysql","root:123456@tcp(www.wsj0819.cn:3306)/test")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	type num struct {
		id int
		num int
	}

	stmt,err := db.Prepare("select * from num where id = ?")
	if err != nil {
		log.Fatalln(err)
	}
	rows,err := stmt.Query(`1' or '1' = '1`)

	for rows.Next(){
		n := num{}
		err = rows.Scan(&n.id,&n.num)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(n)
	}
}
