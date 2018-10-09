package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func getConn() (*sql.DB,error)  {
	db,err := sql.Open("mysql","root:123456@tcp(www.wsj0819.cn:3306)/test")
	//db.SetMaxOpenConns(2)
	if err != nil {
		log.Fatalln(err)
		return nil,err
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
		return nil,err
	}
	log.Println("connection success")
	return db,nil
}




func main()  {
	db,err := getConn()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()


	rows,err := db.Query("select * from user")

	columns,_ := rows.Columns()
	values := make([]sql.RawBytes,len(columns))
	valuesAddr := make([]interface{},len(columns))
	results := []map[string]string{}

	//false
	//for i,addr := range values {
	//	valuesAddr[i] = &addr
	//}

	for i,_ := range values {
		valuesAddr[i] = &values[i]
	}

	for rows.Next(){
		err := rows.Scan(valuesAddr...)
		if err != nil {
			log.Fatalln(err)
		}
		result := make(map[string]string)
		for i,v := range values{
			result[columns[i]] = string(v)
		}
		results = append(results,result)
	}
	rows.Close()

	for k,v := range results{
		fmt.Println(k,v)
	}

}