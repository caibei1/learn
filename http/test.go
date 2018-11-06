package main

import (
	"fmt"
	"net/http"
)

func main()  {

	fmt.Println(1)

	http.HandleFunc("/hello",Hello)
	http.ListenAndServe(":8080",nil)

}

func Hello(w http.ResponseWriter,r *http.Request)  {
	sleep()
}