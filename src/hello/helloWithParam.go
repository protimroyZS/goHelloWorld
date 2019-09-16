package hello

import (
	"fmt"
	"net/http"
)

// NameHandler for HelloWorld in server with name
func NameHandler(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")

	if name != "" {
		fmt.Fprintf(w, "Hello %v!", name)
	} else {
		fmt.Fprint(w, "Hello HTTP!")
	}

}
