package main

import (
	"time"
	"fmt"
)

var i = 1

func sleep()  {
	time.Sleep(time.Second*5)
	i++
	fmt.Println(i)
}
