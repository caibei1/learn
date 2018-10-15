package main

import "fmt"

type A1 struct {
	a1 string
	a2 string
	B1
}

type B1 struct {
	b1 string
	b2 string
}

func (b B1)String() string {
	return fmt.Sprintf("%s %s",b.b1,b.b2)
}

func main()  {


	b := B1{
		"b1",
		"b2",
	}

	a := A1{
		"a1",
		"a2",
		b,
	}

	fmt.Println(a.B1.b1)
	fmt.Println(a.b1)
	fmt.Println(a.String())
	fmt.Println(a.B1.String())
}
