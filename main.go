package main

import (
	"log"
	"net/http"

	"github.com/MakMemEng/go-api-middle-road/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HelloHandler)

	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.ArticleListHandler)
	http.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	http.HandleFunc("/article/nice", handlers.PostNiceHandler)

	http.HandleFunc("/comment", handlers.PostCommentHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	/* 以下と同じ意味
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
	*/
}
