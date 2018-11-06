package main

func main()  {
	m := make(map[int]int)

	for i := 0;i<1000;i++{
		go writeMap(m,i,i)
		go readMap(m,i)
	}
}

func readMap(m map[int]int,key int) int {
	return m[key]
}

func writeMap(m map[int]int,key int,value int)  {
	m[key] = value
}