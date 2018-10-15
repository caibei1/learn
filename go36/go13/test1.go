package main

import "fmt"

type A struct {
	a1 string
	a2 string
}

type B struct {
	b1 string
	b2 string
}

func (b B)String() string {
	return fmt.Sprintf("%s %s",b.b1,b.b2)
}

func main()  {
	a := A{
		"a1",
		"a2",
	}

	b := B{
		"b1",
		"b2",
	}

	fmt.Println(a)
	fmt.Println(b)
}
