package main

import "fmt"

func main()  {
	number := []int{1,2,3,4,5,6}
	for i := range number{
		if i == 3 {
			number[i] |= i
		}
	}
	fmt.Println(number)
}
