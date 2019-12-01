package test

import (
	"net/http"
)

var Object := http.FileServer(http.Dir("public"))