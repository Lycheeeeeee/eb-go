package main

import (
	"net/http"
	"github.com/Lycheeeeeee/eb-go/test"
)

func main() {
	fs := test.Object
	http.Handle("/", fs)
	http.ListenAndServe(":5000", nil)
}
