package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("helo from chi"))
	})

	fmt.Println("hi")
	//группировка запросов
	r.Route("/articles", func(r chi.Router) {
		r.Post("/", createArticle)
		r.Get("/search", searchArticles)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", getOneArticle)
			r.Delete("/", deleteOneArticle)
		})
	})

	http.ListenAndServe(":8080", r)
}
