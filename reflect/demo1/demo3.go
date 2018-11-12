package main

import (
	"reflect"
	"fmt"
)

type user struct {
	name string
	age int
}


type manager struct {
	user
	title string
	name int
}

func main()  {
	var m manager
	t := reflect.TypeOf(&m)

	if t.Kind() == reflect.Ptr{
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++{  //NumField 字段数
		f := t.Field(i)
		fmt.Println(f.Name,f.Type,f.Offset) //结构内的偏移量，以字节为单位
		if f.Anonymous{			//输出匿名结构字段
			for x := 0; x < f.Type.NumField(); x++ {
				af := f.Type.Field(x)
				fmt.Println(af.Name,af.Type)
			}
		}
	}

	//按照名字查找   如果user里的name和外面结构体的name  名字一样 查找的是外面的
	name,ok := t.FieldByName("name")
	if ok {
		fmt.Println("按照名字查找")
		fmt.Println(name.Name,name.Type)
	}

	//按照多级索引查找   0,1代表查找user里的age
	age := t.FieldByIndex([]int{0,1})
	fmt.Println("按照多级索引查找")
	fmt.Println(age.Name,age.Type)
}

