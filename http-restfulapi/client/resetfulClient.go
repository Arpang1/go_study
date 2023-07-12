package main

import (
	"bytes"
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

func main() {
	//通过HTTP请求调用远程函数
	url := "http://localhost:31107/add"
	param := addParam{
		X: 110,
		Y: 20,
	}
	jsonParam, err := json.Marshal(param)
	if err != nil {
		log.Printf("json.Marshal failed,err:%v\n", err)
		return
	}
	//调用post请求
	resp, _ := http.Post(url, "application/json", bytes.NewBuffer(jsonParam))
	if err != nil {
		log.Printf("http.Post failed,err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	//获取post响应，进行解析
	b, _ := io.ReadAll(resp.Body)
	var result addResult
	err = json.Unmarshal(b, &result)
	if err != nil {
		log.Printf("json.Unmarshal failed,err:%v\n", err)
		return
	}
	log.Printf("result:%#v\n", result)
}
