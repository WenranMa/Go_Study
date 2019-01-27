package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func main() {
	db := database{"shoes": 50, "socks": 5}

	/*
		第一种写法：

		db.list是一个方法值，这个类型的值是
		func(w http.ResponseWriter, req *http.Request)

		db.list是一个实现类似handler接口中ServerHTTP行为的函数，
		但却不满足http.Handler接口并且不能直接传给mux.Handle。

		语句http.HandlerFunc(db.list)是一个转换而非一个函数调用，
		因为http.HandlerFunc是一个类型。它有如下的定义:

		package http
		type HandlerFunc func(w ResponseWriter, r *Request)

		func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
			f(w, r)
		}

		HandlerFunc显示了在Go语言接口机制中一些不同寻常的特点。
		HandlerFunc是一个有实现了接口http.Handler方法的函数类型。
		ServeHTTP方法的行为调用了它本身的函数。
		因此HandlerFunc是一个让函数值满足一个接口的适配器。
		这个技巧让一个单一的类型例如database以多种方式满足http.Handler接口:
		一种通过它的list方法，一种通过它的price方法。
	*/
	// mux := http.NewServeMux()
	// mux.Handle("/list", http.HandlerFunc(db.list))
	// mux.Handle("/price", http.HandlerFunc(db.price))
	// log.Fatal(http.ListenAndServe("localhost:8000", mux))

	/*
		第二种写法：
		因为handler通过这种方式注册非常普遍，ServeMux有一个方便的HandleFunc方法:
	*/

	// mux := http.NewServeMux()
	// mux.HandleFunc("/list", db.list)
	// mux.HandleFunc("/price", db.price)
	// log.Fatal(http.ListenAndServe("localhost:8000", mux))

	/*
		第三种写法：
		net/http包提供了一个全局的ServeMux实例DefaultServerMux和包级别的http.Handle和http.HandleFunc函数。
	*/

	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}
