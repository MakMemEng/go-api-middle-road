package handlers

import (
	"fmt"
	"io"
	"net/http"
)

// /hello のハンドラ
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	/* 	// io.Writer interface型の変数
		var w io.Writer

		// Write(p []byte) (n int, err, error)を持つ構造体
		type MyType1 struct{}
		func (t MyType1) Write(p []byte) (n int, err, error) {
			//
		}
		w = MyType1{} // 代入OK

		// Write(p []byte) (n int, err, error)を持たない構造体
		type MyType2 struct{}
		w = MyType2{} // 代入NG(コンパイルエラー)
	 */
	 io.WriteString(w, "Hello, World!\n")
}

// /article のハンドラ
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article...\n")
}

// /article/list のハンドラ
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Article List\n")
}

// /article/1 のハンドラ
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID := 1
	resString := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resString)
}

// /article/nice のハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}

// /comment のハンドラ
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
}
