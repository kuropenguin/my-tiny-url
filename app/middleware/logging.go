package middleware

import (
	"log"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// リクエストに関する情報をログに記録
		log.Printf("Request received: %s %s", r.Method, r.URL.Path)
		// 次の処理（次のミドルウェアまたは最終的なハンドラ）へリクエストを渡す
		next.ServeHTTP(w, r)
	})
}
