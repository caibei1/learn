package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"strconv"
)

var db *leveldb.DB

//打开数据库
func init()  {
	var err error
	//数据存储路径和一些初始文件
	db,err = leveldb.OpenFile("./levelDB/db",nil)
	if err != nil {
		log.Fatalln(err)
	}
}


func main() {
	save("1","111")
	save("2","222")
	save("3","333")
	printStrVal("1000")
	delete("2")

	var kv = make(map[string]string)
	for i := 4; i < 20 ; i++{
		istr := strconv.Itoa(i)
		kv[istr] = istr+istr+istr
	}
	saveAll(kv)

	//getAll()
	iter := db.NewIterator(nil, nil)
	for ok := iter.Seek([]byte("1010")); ok; ok = iter.Next() {
		fmt.Printf("查找数据:%s, value:%s\n", iter.Key(), iter.Value())
	}
	db.Close()
}

//存数据
func save(key string,value string)  {
	db.Put([]byte(key),[]byte(value),nil)
}

//获取数据
func get(key string) []byte  {

	value,err := db.Get([]byte(key),nil)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return value
}

func printStrVal(key string)  {
	value := get(key)
	fmt.Println(string(value))
}

// 遍历所有数据
func getAll()  {
	ite := db.NewIterator(nil,nil)
	for ite.Next(){
		fmt.Println(string(ite.Key()),string(ite.Value()))
	}
	ite.Release()
}

//删除
func delete(key string)  {
	db.Delete([]byte(key),nil)
}

////批量写入数据
func saveAll(kv map[string]string)  {
	batch := new(leveldb.Batch)
	for k,v := range kv{
		batch.Put([]byte(k),[]byte(v))
	}
	db.Write(batch,nil)
}