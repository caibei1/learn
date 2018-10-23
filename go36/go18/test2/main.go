package test

import "fmt"

func main1()  {
	number := [...]int{1,2,3,4,5,6}
	maxIndex := len(number)-1
	for i,e := range number{
		if i == maxIndex{
			number[0] += e
		} else {
			number[i+1] += e
		}
	}
	fmt.Println(number)
}
