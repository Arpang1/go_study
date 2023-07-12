package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Println("get method")
	}
	n := rand.Intn(2)
	fmt.Println("n", n)
	if n == 0 {
		time.Sleep(time.Second * 3)
		_, err := fmt.Fprintf(w, "hello,slow response  path: "+r.URL.Path)
		if err != nil {
			return
		}
		return
	}
	fmt.Fprintf(w, "hello,quick response  path: "+r.URL.Path)
}

func main() {
	http.HandleFunc("/hello", Hello)
	log.Printf("start server port:8080 ...\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
