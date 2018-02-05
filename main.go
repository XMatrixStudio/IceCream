package main

import (
	"net/http"
	"./router"
)

func main() {
	http.ListenAndServe(":8080", router.Router())
}
