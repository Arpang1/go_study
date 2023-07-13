package middleware

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/proto"
	"net/http"
)

// 会拦截响应，给响应加上一些东西
func Foeward(ctx context.Context, w http.ResponseWriter, pro proto.Message) error {
	w.Header().Set("result-id", "000-success")
	w.WriteHeader(http.StatusOK)
	return nil
}

// 当路由发送错误时，会调用这个函数，我们可以处理
func RoutingErrorHandler(ctx context.Context, smux *runtime.ServeMux, mar runtime.Marshaler, w http.ResponseWriter, r *http.Request, n int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	_, err := w.Write([]byte(`{"code":404,"url_path":"` + r.URL.Path + `","msg":"路径不存在"}`))
	if err != nil {
		return
	}
	return
}

// 保留请求头信息,保存到上下文
func IncomingHeaderMatcher(key string) (string, bool) {
	if key == "X-User-Id" {
		return key, true
	}
	return runtime.DefaultHeaderMatcher(key)
}
