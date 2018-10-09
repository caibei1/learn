package main

import (
	"fmt"
	"net/http"
	"time"
)

func main()  {

	fmt.Println(1)

	http.HandleFunc("/hello",Hello)
	http.ListenAndServe(":8080",nil)

}

func Hello(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("hello")
	time.Sleep(100*time.Second)
}