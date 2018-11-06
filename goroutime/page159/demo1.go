package main

import (
	"sync"
	"fmt"
)

func main()  {
	var wg sync.WaitGroup
	var gs [5]struct{      //用于实现类似TLS(局部存储)功能
		id int            //编号
		result int		  //返回值
	}

	for i := 0; i < len(gs); i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			gs[id].id = id
			gs[id].result = (id +1) * 100
		}(i)
	}

	wg.Wait()
	fmt.Println(gs)
}
