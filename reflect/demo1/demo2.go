package main

import (
	"reflect"
	"fmt"
)

func main()  {
	//构造基础复合类型
	a := reflect.ArrayOf(10,reflect.TypeOf(byte(0)))
	b := reflect.MapOf(reflect.TypeOf(""),reflect.TypeOf(0))
	fmt.Println(a,b)

	//方法elem返回指针，数组，切片，字典，活通道的基类型
	fmt.Println(reflect.TypeOf(map[string]int{}).Elem())
	fmt.Println(reflect.TypeOf([]int32{}).Elem())

	i := 1
	fmt.Println(reflect.TypeOf(i))
	//fmt.Println(reflect.TypeOf(i).Elem())  //报错 不是以上类型
	fmt.Println(reflect.TypeOf(&i))
	fmt.Println(reflect.TypeOf(&i).Elem())
}
