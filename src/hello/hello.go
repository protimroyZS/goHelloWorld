package hello

import (
	"fmt"
	"net/http"
)

// HelloHandler for HelloWorld in server
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello HTTP!")
}
