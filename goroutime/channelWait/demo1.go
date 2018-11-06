package main

import "fmt"

func main()  {
	exit := make(chan struct{})

	go func() {
		fmt.Println(111)
		exit <- struct{}{}
	}()
	<- exit
	close(exit)
}

