package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
	"fmt"
	"time"
	"io/ioutil"
	"encoding/json"
)

//var Router *mux.Router

func main()  {

	Router := mux.NewRouter()
	//配置路由
	Router.HandleFunc("/hello",hello).Methods("GET")  //可以不定参
	Router.HandleFunc("/hello1",hello1).Methods("POST")

	//设置端口 路由
	server := http.Server{
		Addr:":8080",
		ReadTimeout:time.Second,
		WriteTimeout:time.Second,
		Handler:Router,
	}
	//启动监听
	server.ListenAndServe()
}


//get
func hello(w http.ResponseWriter,r *http.Request)  {
	//参数解析
	params,_ := url.ParseQuery(r.URL.RawQuery)
	msg := "hello"+params["msg"][0]
	fmt.Println(msg)
	w.Write([]byte(msg))
}

//post
func hello1(w http.ResponseWriter,r *http.Request)  {
	//参数解析
	body,_ := ioutil.ReadAll(r.Body)
	var params map[string]string
	json.Unmarshal(body, &params)
	fmt.Println(params["msg"])
	resp := "hello" + params["msg"]
	w.Write([]byte(resp))
}

