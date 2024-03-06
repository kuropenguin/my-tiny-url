package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/kuropenguin/my-tiny-url/app/config"
	"github.com/kuropenguin/my-tiny-url/app/mysql"
)

type responseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
}

func newResponseWriterWrapper(w http.ResponseWriter) *responseWriterWrapper {
	return &responseWriterWrapper{ResponseWriter: w, statusCode: 0}
}

func (rww *responseWriterWrapper) WriteHeader(code int) {
	rww.statusCode = code
	rww.ResponseWriter.WriteHeader(code)
}

func Transaction(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tx, err := mysql.GetDB().Begin()
		if err != nil {
			panic(err)
		}
		qTx := mysql.GetSQLCQueries().WithTx(tx)
		ctx := context.WithValue(r.Context(), config.TxKey, qTx)
		r = r.WithContext(ctx)

		wrr := newResponseWriterWrapper(w)
		next.ServeHTTP(wrr, r)

		// tx.Rollback()
		if http.StatusOK <= wrr.statusCode && wrr.statusCode < http.StatusBadRequest {
			tx.Rollback()
			// tx.Commit()
			log.Println("commit")
		} else {
			tx.Rollback()
			log.Println("rollback")
		}
	})
}
