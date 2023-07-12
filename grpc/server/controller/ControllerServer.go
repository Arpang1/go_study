package controller

import (
	"context"
	"fmt"
	"log"
	"server/pb/book"
	"strconv"
	"sync"
)

var id int
var books = make(map[string]book.Book)

func NewBook(name, author, price, description string) {
	id++
	books[name] = book.Book{
		Id:          strconv.Itoa(id),
		Name:        name,
		Author:      author,
		Price:       price,
		Description: description,
	}
}

type GrpcServer struct {
	sync.RWMutex
	book.UnimplementedBookServiceServer
}

func NewGrpcServer() *GrpcServer {
	return &GrpcServer{}
}

func (s *GrpcServer) GetBook(ctx context.Context, req *book.GetBookRequest) (*book.GetBookResponse, error) {
	log.Println("GetBook Request: ", req.Name)
	if v, ok := books[req.Name]; ok {
		return &book.GetBookResponse{Book: &v}, nil
	} else {
		return &book.GetBookResponse{}, fmt.Errorf("book not found")
	}
}

func (s *GrpcServer) GetBooks(ctx context.Context, req *book.GetBooksRequest) (*book.GetBooksResponse, error) {
	s.RLock()
	defer s.RUnlock()
	log.Println("GetBooks Request: ", req)
	respBooks := make([]*book.Book, 0)
	for _, v := range books {
		b := v
		respBooks = append(respBooks, &b)
	}
	return &book.GetBooksResponse{Book: respBooks}, nil
}
func (s *GrpcServer) CreateBook(ctx context.Context, req *book.CreateBookRequest) (*book.CreateBookResponse, error) {
	s.RLock()
	defer s.RUnlock()
	log.Println("CreateBook Request: ", req)
	id++
	req.Book.Id = strconv.Itoa(id)
	books[req.Book.Name] = *req.Book
	return &book.CreateBookResponse{Name: req.Book.Name}, nil
}
