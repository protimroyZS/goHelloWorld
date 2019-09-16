package hello

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//HTTPPOSTTestCase is the data format for passing data
type HTTPPOSTTestCase struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// POSTHandler for HelloWorld in server with name and age
func POSTHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var data HTTPPOSTTestCase
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	name := data.Name
	age := data.Age
	log.Print("Dekh!!")
	log.Print(name, age)
	if name != "" && age != 0 {
		fmt.Fprintf(w, "Hello %v! \n Your Age is %v", name, age)
	} else {
		fmt.Fprint(w, "Hello HTTP!")
	}

}
