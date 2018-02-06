package main

import (
	"net/http"
	"github.com/XMatrixStudio/IceCream/router"
)

func main() {
	http.ListenAndServe(":30020", router.Router())
}
