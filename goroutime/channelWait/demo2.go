package main

import "fmt"

func main()  {
	exit := make(chan struct{})

	go func() {
		fmt.Println(11)
		close(exit)   //关闭通道
	}()
	<- exit  //如果通道关闭，立即解除阻塞
}
