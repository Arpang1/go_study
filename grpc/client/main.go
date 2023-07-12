package main

import (
	"client/pb/book"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

//type PerRPCCredentials interface {
//	GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
//	RequireTransportSecurity() bool
//}

type ClientTokenAuth struct {
}

func (c *ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"AK": "123456",
		"SK": "abcdef",
	}, nil
}
func (c *ClientTokenAuth) RequireTransportSecurity() bool {
	return false
}

func main() {
	//WithTransportCredentials  insecure.NewCredentials()
	//创建客户端到服务端的连接
	options := make([]grpc.DialOption, 0)
	options = append(options, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))
	options = append(options, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial("127.0.0.1:31107", options...)
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
	//resp3, err := client.CreateBook(context.Background(), &book.CreateBookRequest{Book: &book.Book{
	//	Name:        "格林童话",
	//	Author:      "格林",
	//	Price:       "26.00",
	//	Description: "格林童话故事",
	//}})
	resp3, err := client.CreateBook(context.Background(), &book.CreateBookRequest{Book: &book.Book{
		Name:        "平凡的世界",
		Author:      "路遥",
		Price:       "28.93",
		Description: "一个村庄的兴衰史",
	}})
	//resp3, err := client.CreateBook(context.Background(), &book.CreateBookRequest{Book: &book.Book{
	//	Name:        "红楼梦",
	//	Author:      "曹雪芹",
	//	Price:       "28.93",
	//	Description: "红楼旧梦",
	//}})
	//resp3, err := client.CreateBook(context.Background(), &book.CreateBookRequest{Book: &book.Book{
	//	Name:        "三国演义",
	//	Author:      "罗贯中",
	//	Price:       "28.93",
	//	Description: "三国鼎立局面",
	//}})
	if err != nil {
		log.Println("failed to create book: " + err.Error())
	} else {
		fmt.Println("创建书籍：")
		fmt.Println(resp3.Name)
	}
}
