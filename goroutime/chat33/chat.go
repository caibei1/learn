package main

import "sync"

var (
	UserMap = make(map[string]*User)
)

type User struct {
	Id       string
	Level    int
	clientsMutex sync.Mutex
	Clients  map[string]string
}

func (u *User)Lock()  {
	u.clientsMutex.Lock()
}


func (u *User)UnLock()  {
	u.clientsMutex.Unlock()
}

func (u *User)Get(target string) string {
	u.Lock()
	a := u.Clients[target]
	u.UnLock()
	return a
}

func (u *User)Set()  {

}