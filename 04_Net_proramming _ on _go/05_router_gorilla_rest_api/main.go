package main

//doesnot work..  pofixit...

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Post struct {
	ID    string `json: "id"`
	Title string `json: "title"`
	Body  string `json: "body"`
}

var posts []Post

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("in get postS metod")
	json.NewEncoder(w).Encode(posts)

}

func getPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in get ONE post metod")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//fmt.Println(params)
	for _, item := range posts {
		fmt.Println(item)
		if item.ID == params["id"] {
			fmt.Println("in get ONE ")
			json.NewEncoder(w).Encode(item)
		}
	}
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("in Create post metod 1")
	var post Post
	_ = json.NewDecoder(r.Body).Decode(post)
	post.ID = strconv.Itoa(rand.Intn(1000000))
	posts = append(posts, post)
	fmt.Println("in Create post metod 2")
	json.NewEncoder(w).Encode(&post)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range posts {
		if item.ID == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
			var post Post
			_ = json.NewDecoder(r.Body).Decode(post)
			post.ID = params["id"]
			posts = append(posts, post)
			fmt.Println("in Update post metod")
			json.NewEncoder(w).Encode(&post)
			return
		}
	}
	json.NewEncoder(w).Encode(posts)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range posts {
		if item.ID == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
			break
		}
	}
	fmt.Println("in Delete post metod")
	json.NewEncoder(w).Encode(&posts)
}

func main() {

	router := mux.NewRouter()
	//have to crate end points

	posts = append(posts, Post{ID: "1", Title: "My first post", Body: "This is the content of my first post"})
	posts = append(posts, Post{ID: "2", Title: "My second post", Body: "This is the content of my second post"})

	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", createPost).Methods("POST")
	router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

	http.ListenAndServe(":8080", router)

	fmt.Println("online..	")

}
