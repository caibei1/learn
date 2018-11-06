package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main()  {
	defer func() {
		fmt.Println("enter defer")
		if p := recover();p!= nil {
			fmt.Printf("panic:%s\n",p)
		}
		fmt.Println("exit defer")
	}()
	panic(errors.New("11"))
	fmt.Println("last")
}
