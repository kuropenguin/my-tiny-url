package middleware

import (
	"context"
	"net/http"

	"github.com/kuropenguin/my-tiny-url/app/config"
	"github.com/kuropenguin/my-tiny-url/app/mysql"
)

func Transaction(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), config.TxKey, mysql.GetSQLCQueries)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
