package main

import (
	"net/http"
	"./router"
)

func main() {
	http.ListenAndServe(":30020", router.Router())
}
