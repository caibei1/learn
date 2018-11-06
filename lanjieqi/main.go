package main

import (
	"learn/lanjieqi/account"
	"learn/lanjieqi/proxy"
)

func main()  {
	id := "1"
	a := account.New(id,"zhangsan",100)
	a.Query(id)
	a.Update(id,500)
}

func init()  {
	account.New = func(id, name string, value int) account.Account {
		a := &account.AccountImpl{id,name,value}
		p := &proxy.Proxy{a}
		return p
	}
}
