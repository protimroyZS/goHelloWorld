package databaseTrial

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

//HTTPPOSTTestCase is the data format for passing data
type HTTPPOSTTestCase struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//HTTPDBTestCase is the struct from csv
type HTTPDBTestCase struct {
	Method string `json:"method"`
	ID     string `json:"id"`
	Name   string `json:"name"`
	Age    string `json:"age"`
}

//ResponseObject for response
type ResponseObject struct {
	Message string `json:"message"`
}

//Customer for response
type Customer struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// DBReqHandler for HelloWorld in server with database connectivity
func DBReqHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)

		var data HTTPPOSTTestCase
		err := decoder.Decode(&data)
		if err != nil {
			panic(err)
		}
		name := data.Name
		age := data.Age
		if name != "" && age != 0 {
			// fmt.Fprintf(w, "Hello %v! \n Your Age is %v", name, age)

			res, _ := insertCustomer(name, age)
			log.Print(res)
			json.NewEncoder(w).Encode(res)
			return
		}
		fmt.Fprint(w, "Hello HTTP!")
	} else if r.Method == "GET" {
		data, err := getCustomers()
		if err != nil {
			panic(err)
		}
		json.NewEncoder(w).Encode(data)
		return
	} else {
		res := ResponseObject{
			Message: "Wrong Method!",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

}

//PUTHandler for put requests
func PUTHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		urlParts := strings.Split(r.URL.Path, "/")
		id, _ := strconv.Atoi(urlParts[len(urlParts)-1])
		log.Print(id)
		decoder := json.NewDecoder(r.Body)

		var data HTTPPOSTTestCase
		err := decoder.Decode(&data)
		if err != nil {
			panic(err)
		}
		name := data.Name
		age := data.Age

		res, _ := SetCustomers(id, name, age)
		json.NewEncoder(w).Encode(res)
		return
	}
	res := ResponseObject{
		Message: "Wrong Method!",
	}
	json.NewEncoder(w).Encode(res)
	return

}

func insertCustomer(name string, age int) (ResponseObject, error) {
	db, err := sql.Open("mysql", "protim:password@tcp(127.0.0.1:3306)/test2")
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}
	_, err = db.Query("CREATE TABLE IF NOT EXISTS customers (id int NOT NULL AUTO_INCREMENT, name varchar(255) NOT NULL, age int, PRIMARY KEY (id));")
	if err != nil {
		panic(err.Error())
	}
	_, err = db.Query("INSERT INTO customers(`name`, `age`) VALUES(?, ?)", name, age)
	if err != nil {
		panic(err.Error())
	}
	return ResponseObject{
		Message: "Customer Added!",
	}, nil
}

func getCustomers() ([]Customer, error) {
	db, err := sql.Open("mysql", "protim:password@tcp(127.0.0.1:3306)/test2")
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}

	results, err := db.Query("SELECT * FROM customers")
	if err != nil {
		panic(err.Error())
	}

	var allCustomers []Customer

	for results.Next() {
		var f Customer
		err = results.Scan(&f.Id, &f.Name, &f.Age)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		// data, err := json.Marshal(&Customer{
		// 	Id:   f.Id,
		// 	Name: f.Name,
		// 	Age:  f.Age,
		// })
		// log.Print(f.Id, f.Name, f.Age)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		allCustomers = append(allCustomers, f)
		// log.Printf(f.item_name)
	}
	return allCustomers, nil
}

//SetCustomers to edit customers
func SetCustomers(customerID int, name string, age int) (ResponseObject, error) {
	if customerID < 1 {
		return ResponseObject{
			Message: "ID cannot be less than 1!",
		}, nil
	}
	db, err := sql.Open("mysql", "protim:password@tcp(127.0.0.1:3306)/test2")
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Query("UPDATE customers set name = ?, age = ? WHERE id = ?", name, age, customerID)
	if err != nil {
		panic(err.Error())
	}
	return ResponseObject{
		Message: "Customer Updated!",
	}, nil
}
