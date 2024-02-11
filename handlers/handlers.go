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
	 if req.Method == http.MethodGet {
		 io.WriteString(w, "Hello, World!\n")
	 } else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	 }
}

// /article のハンドラ
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Article...\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// /article/list のハンドラ
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "Article List\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// /article/1 のハンドラ
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		articleID := 1
		resString := fmt.Sprintf("Article No.%d\n", articleID)
		io.WriteString(w, resString)
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// /article/nice のハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Nice...\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// /comment のハンドラ
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Comment...\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}
