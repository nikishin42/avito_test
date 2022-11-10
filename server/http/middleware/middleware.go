package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
	"server/server/logger"
)

type Middleware struct {
	l logger.Logger
}

func New(l logger.Logger) *Middleware {
	return &Middleware{l: l}
}

func (mw *Middleware) RecoverMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			rec := recover()
			if rec != any(nil) {
				stacktrace := string(debug.Stack())
				switch t := rec.(type) {
				case string:
					err = fmt.Errorf("panic: %s, stacktrace: %s", t, stacktrace)
				case error:
					err = fmt.Errorf("panic: %v, stacktrace: %s", t, stacktrace)
				default:
					err = errors.New("unknown panic")
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}()

		h.ServeHTTP(w, r)
	})
}
