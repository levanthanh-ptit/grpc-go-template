package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "OK")
	})
	fmt.Println("server start: 127.0.0.1:8010")
	http.ListenAndServe("127.0.0.1:8010", nil)
}
