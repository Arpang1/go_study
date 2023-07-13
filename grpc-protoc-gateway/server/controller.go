package server

import (
	"context"
	"fmt"
	"grpc-protoc-gateway/pb/book"
	"sync"
)

type BookService struct {
	book.UnimplementedBookServerServer
}

func NewBookService() *BookService {
	l := localStorage{
		Count: 0,
		DB:    make(map[string]*book.Book),
	}
	Db = &l
	return &BookService{}
}

func (s *BookService) CreateBook(ctx context.Context, req *book.CreateBookRequest) (*book.CreateBookResponse, error) {
	Db.mux.Lock()
	defer Db.mux.Unlock()
	createbook := book.Book{
		Id:          Db.GetId(),
		Name:        req.Book.Name,
		Author:      req.Book.Author,
		Description: req.Book.Description,
	}
	err := Db.Store(&createbook)
	if err != nil {
		return nil, fmt.Errorf("本地储存失败：", err.Error())
	}
	return &book.CreateBookResponse{Name: req.Book.Name}, nil
}

func (s *BookService) GetBook(ctx context.Context, req *book.GetBookRequest) (*book.GetBookResponse, error) {
	Db.mux.Lock()
	defer Db.mux.Unlock()
	fmt.Println(Db)
	if v, err := Db.Load(req.Name); err != nil {
		return nil, fmt.Errorf("Get 失败：", err)
	} else {
		return &book.GetBookResponse{Book: v}, nil
	}

}

var Db *localStorage

type localStorage struct {
	Count int32
	DB    map[string]*book.Book
	mux   sync.Mutex
}

func (l *localStorage) GetId() int32 {
	l.Count = l.Count + 1
	return l.Count
}

func (l *localStorage) Store(b *book.Book) error {
	if b == nil {
		return fmt.Errorf("传入的值为空")
	}
	if b.Id <= 0 {
		return fmt.Errorf("id不能为0")
	}
	// 检查 DB 是否已经被初始化
	if l == nil {
		l.DB = make(map[string]*book.Book) // 如果 DB 未初始化，则使用 make 初始化它
	}
	l.DB[b.Name] = b
	return nil
}

func (l *localStorage) Load(name string) (*book.Book, error) {
	if v, ok := l.DB[name]; !ok {
		return nil, fmt.Errorf("找不到此书")
	} else {
		return v, nil
	}
}
