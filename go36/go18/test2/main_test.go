package test

//1、文件名须以"_test.go"结尾
//2、方法名须以"Test"打头，并且形参为 (t *testing.T)

import (
	"testing"
	"fmt"
)

func BenchmarkMainf(t *testing.B)  {
	main1()
}




func TestHello(t *testing.T) {

	fmt.Println("TestHello")

}



func TestWorld(t *testing.T) {

	fmt.Println("TestWorld")

}
