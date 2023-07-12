package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"
	"time"
)

type respData struct {
	resp *http.Response
	err  error
}

func doCall(ctx context.Context) {
	transport := http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			default:
				return net.Dial(network, addr)
			}
		},
	}
	client := http.Client{
		Transport: &transport,
	}
	respChan := make(chan *respData, 1)
	req, err := http.NewRequest("GET", "http://localhost:8080/hello", nil)
	if err != nil {
		panic("请求失败" + err.Error())
		return
	}
	req = req.WithContext(ctx)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		resp, err := client.Do(req)
		pack := &respData{
			resp: resp,
			err:  err,
		}
		respChan <- pack
		wg.Done()
	}()
	select {
	case <-ctx.Done():
		req.WithContext(context.Background())
		fmt.Println("timeout")
	case resp := <-respChan:
		if resp.err != nil {
			fmt.Println("Client Error:", resp.err)
			return
		} else {
			b, err := io.ReadAll(resp.resp.Body)
			if err != nil {
				panic(err)
			}
			fmt.Println("Client :", string(b))
		}
		defer resp.resp.Body.Close()
		fmt.Println("Server Code:", resp.resp.StatusCode)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	doCall(ctx)
}
