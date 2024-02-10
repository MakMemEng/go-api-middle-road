package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
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

	http.HandleFunc("/", helloHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	/* 以下と同じ意味
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
	*/
}
