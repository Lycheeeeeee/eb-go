package main

import (
	"net/http"
	"test"
)

func main() {
	fs := test.Object
	http.Handle("/", fs)
	http.ListenAndServe(":5000", nil)
}
