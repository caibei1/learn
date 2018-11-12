package main

import (
	"reflect"
	"fmt"
)

type Y int

func (Y) String() string  {
	return ""
}

func main()  {
	var a Y
	t := reflect.TypeOf(a)

	st := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	fmt.Println(t.Implements(st))//实现接口

	it := reflect.TypeOf(0)
	fmt.Println(t.ConvertibleTo(it)) //可转换？ t是否可以转换成it的类型

	fmt.Println(t.AssignableTo(st),t.AssignableTo(it)) //可归属？
}
