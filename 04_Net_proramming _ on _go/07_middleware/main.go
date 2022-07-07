package main

import (
	"fmt"
	"net/http"
)

// type MiddlewareHandler = func(next http.Handler) http.Handler

func SimpleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "before handler")
		next.ServeHTTP(w, r)
		fmt.Fprintln(w, "after handler")
	})
}

func main() {
	http.Handle("/middleware",
		SimpleMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "in handler")
		})))
	http.ListenAndServe(":8080", nil)
}

// После отправки запроса мы должны увидеть следующую картину.
// → curl localhost:8080/middleware
// Hello, middleware
// in handler
// Hello from middleware!
