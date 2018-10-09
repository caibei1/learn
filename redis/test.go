package main
//www.wsj0819.cn:6379

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

func main()  {
	//连接redis
	c,err := redis.Dial("tcp","www.wsj0819.cn:6379")
	if err != nil {
		log.Fatalln("redis connection false")
	}
	log.Println("redis connection success")
	defer c.Close()

	//读写
	r,err := c.Do("SET","a","a不会过期")
	if err != nil {
		log.Fatalln("set failed:",err)
	}
	fmt.Println("写",r)

	value,err := redis.String(c.Do("GET","a"))

	if err != nil{
		log.Println(err)
	} else {
		fmt.Println(value)
	}

	//删除数据
	r,err = c.Do("DEL","a")
	if err != nil{
		log.Println(err)
	} else {
		fmt.Println("删除数据",r)
	}

	//设置过期
	r,err = c.Do("SET","b","b会过期","EX",1)
	if err != nil {
		log.Fatalln("set failed:",err)
	}
	fmt.Println(r)

	exist,err := redis.Bool(c.Do("EXISTS","b"))

	fmt.Println("b是否存在",exist)

	//time.Sleep(time.Second)
	value,err = redis.String(c.Do("GET","b"))
	if err != nil{
		log.Println(err)
	} else {
		fmt.Println(value)
	}

	time.Sleep(time.Second)
	exist,err = redis.Bool(c.Do("EXISTS","b"))
	fmt.Println("b是否存在",exist)
}
