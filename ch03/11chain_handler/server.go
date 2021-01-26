// リスト3.11
package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

//ハンドラいれてハンドラだす
func log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("ハンドラが呼び出されました - %T\n", h)
		h.ServeHTTP(w, r)
	})
}

//ハンドラいれてハンドラだす
func protect(h http.Handler) http.Handler {
	//
	//ここで型変換して関数を関数型にしてハンドラ関数からハンドラにする
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// some code to make sure the user is authorized

		//ハンドラのメソッドを呼び出す
		//多分h を見た瞬間に、hってなにってなり、logにいく
		h.ServeHTTP(w, r)
	})
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	hello := HelloHandler{}
	http.Handle("/hello", protect(log(hello)))
	server.ListenAndServe()
}
