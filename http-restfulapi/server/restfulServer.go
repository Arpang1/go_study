package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type addParam struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type addResult struct {
	Code int `json:"code"`
	Data int `json:"data"`
}

func add(x, y int) int {
	return x + y
}
func addHandler(w http.ResponseWriter, r *http.Request) {
	//解析参数
	b, _ := io.ReadAll(r.Body)
	var p addParam
	err := json.Unmarshal(b, &p)
	if err != nil {
		return
	}
	//查看获取client端的值
	log.Printf("get param:%#v\n", p)
	//业务逻辑
	ret := add(p.X, p.Y)
	//返回结果
	resByte, _ := json.Marshal(addResult{
		Code: 0,
		Data: ret,
	})
	_, err = w.Write(resByte)
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/add", addHandler)
	log.Printf("server start poet 31107 ...\n")
	err := http.ListenAndServe(":31107", nil)
	if err != nil {
		return
	}
}
