package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

type DB struct {
	dict map[string]interface{}
	lock sync.Mutex
}

const DB_PATH = "./data/dump.db"

var db DB

// get操作
func get(w http.ResponseWriter, r *http.Request) {
	url_list := strings.Split(r.URL.Path, "/")
	db.dict = Load()
	if v, ok := db.dict[url_list[2]]; ok {
		fmt.Fprintf(w, "Succeeded to read key: %s, value: %s", url_list[2], v)
	} else {
		fmt.Fprintf(w, "Key %s does not exist!", url_list[2])
	}
}

// set操作
func set(w http.ResponseWriter, r *http.Request) {
	url_list := strings.Split(r.URL.Path, "/")
	db.dict = make(map[string]interface{})
	if len(url_list) == 4 {
		db.dict = Load()
		db.lock.Lock()
		defer db.lock.Unlock()
		db.dict[url_list[2]] = url_list[3]
		Dump(db.dict)
		fmt.Fprintf(w, "Succeeded to write key: %s, value: %s", url_list[2], url_list[3])
	}
}

// 持久化数据
func Dump(data map[string]interface{}) {
	buf, err := json.Marshal(data)
	handleErr(err, "Marshall error")
	ioutil.WriteFile(DB_PATH, buf, 0644)
}

// 从存储介质加载数据
func Load() map[string]interface{} {
	f, err := ioutil.ReadFile(DB_PATH)
	handleErr(err, "Failed to load data from "+DB_PATH)
	db.dict = make(map[string]interface{})
	err = json.Unmarshal(f, &db.dict)
	handleErr(err, "Can't decode json message")
	return db.dict
}

// 创建存储路径
func init() {
	if _, err := os.Stat(DB_PATH); err != nil {
		_, err1 := os.Create(DB_PATH)
		handleErr(err1, "Failed to create path: "+DB_PATH)
	}
}

// 错误处理
func handleErr(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
	}
}

func main() {
	http.HandleFunc("/get/", get)
	http.HandleFunc("/set/", set)
	err := http.ListenAndServe(":7070", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
