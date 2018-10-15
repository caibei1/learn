package main

import "fmt"

type  S struct {
	age int
}

func (s S)change(i int)  {
	s.age = i
}


func (s *S)change1(i int)  {
	s.age = i
}

func main (){
	s := S{
		1,
	}

	s.change(2)
	fmt.Println(s)

	s.change1(2)
	fmt.Println(s)
}