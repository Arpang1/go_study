package main

import (
	"client/pb/book"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//WithTransportCredentials  insecure.NewCredentials()
	//创建客户端到服务端的连接
	conn, err := grpc.Dial("127.0.0.1:31107", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic("failed to close: " + err.Error())
		}
	}(conn)

	//创建一个grpc客户端
	client := book.NewBookServiceClient(conn)
	//调用服务端的方法
	//通过name查询书籍
	resp, err := client.GetBook(context.Background(), &book.GetBookRequest{Name: "安徒生"})
	if err != nil {
		panic("failed to get book: " + err.Error())
	}
	fmt.Println("Get书籍：")
	fmt.Println(resp.Book)
	//查询所有书籍
	resp2, err := client.GetBooks(context.Background(), &book.GetBooksRequest{})
	if err != nil {
		panic("failed to get books: " + err.Error())
	}
	fmt.Println("所有书籍：")
	for _, v := range resp2.Book {
		fmt.Println(v)
	}
	//创建一本书籍
	resp3, err := client.CreateBook(context.Background(), &book.CreateBookRequest{Book: &book.Book{
		Name:        "格林童话",
		Author:      "格林",
		Price:       "26.00",
		Description: "格林童话故事",
	}})
	if err != nil {
		panic("failed to create book: " + err.Error())
	}
	fmt.Println()
	fmt.Println("创建书籍：" + resp3.Name)
}
