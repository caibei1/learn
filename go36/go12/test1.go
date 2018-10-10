package main

import "fmt"


//只要两个函数的参数列表和结果列表中的元素顺序及其类型是一致的,们就可以说它们是一样的函数，或者说是实现了同一个函数类型的函数。



type Printer func(content string) (n int, err error)

func printToStd(content string) (byteNum int, err error) {
	return fmt.Println(content)
}

func main()  {
	var p  Printer
	p = printToStd
	p("something")

	//printToStd("s")
}
