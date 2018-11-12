package main

import (
	"reflect"
	"fmt"
)

//tag的使用  常用于orm映射 或者数据格式校验

type user1 struct {
	name string `field1:"name1" type:"varchar(50)"`
	age int     `field1:"age1"  type:"int"`
}

func main()  {
	var u user1
	t := reflect.TypeOf(u)

	for i := 0;i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("%s: %s %s\n",f.Name,f.Tag.Get("field1"),f.Tag.Get("type"))
	}
}
