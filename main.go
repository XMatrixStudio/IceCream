package main

import (
	"net/http"

	"github.com/XMatrixStudio/IceCream.Server/router"
)

func main() {
	http.ListenAndServe(":8080", router.Router())
}
