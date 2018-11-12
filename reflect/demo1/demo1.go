package main

import (
	"reflect"
	"fmt"
)

type X int

//typeOf  valueOf 会将任何传入的对象转换为接口类型

func main()  {
	var a X = 100
	t := reflect.TypeOf(a)
	fmt.Println(t.Name(),t.Kind())

	t1 := reflect.ValueOf(a)
	fmt.Println(t1.Type(),t1.Kind())

	//区分基类型和指针类型，他们不是属于同一种类型
	tx,tp := reflect.TypeOf(a),reflect.TypeOf(&a)
	fmt.Println(tx,tp)
}
