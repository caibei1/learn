package main

import "fmt"


func main (){
	array := [4]int{10, 20, 30, 40}
	slice := array[0:2]
	newSlice := append(append(append(slice, 50), 100), 150)
	newSlice[1] += 1
	fmt.Println(slice)

}