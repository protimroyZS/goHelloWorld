package databaseTrial

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

func TestDBReqHandler(t *testing.T) {
	r := readCSV("/home/raramuri/Documents/dbHandlerTests.csv")
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("Question: %s Answer %s\n", record[0], record[1])
		age, _ := strconv.Atoi(record[3])
		requestBody, _ := json.Marshal(&HTTPPOSTTestCase{
			Name: record[2],
			Age:  age,
		})
		request, _ := http.NewRequest(record[0], "/hello", bytes.NewBuffer(requestBody))
		response := httptest.NewRecorder()

		DBReqHandler(response, request)

		if response.Code != http.StatusOK {
			t.Fatalf("Expected status code %v, but received: %v", "200", response.Code)
		}
		fmt.Println(response.Body)
	}
}

func TestPUTHandler(t *testing.T) {
	r := readCSV("/home/raramuri/Documents/putHandlertests.csv")
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("Question: %s Answer %s\n", record[0], record[1])
		age, _ := strconv.Atoi(record[3])
		requestBody, _ := json.Marshal(&HTTPPOSTTestCase{
			Name: record[2],
			Age:  age,
		})
		requestURL := "/hello/" + record[1]
		fmt.Println(requestURL)
		request, _ := http.NewRequest(record[0], requestURL, bytes.NewBuffer(requestBody))
		response := httptest.NewRecorder()

		PUTHandler(response, request)

		if response.Code != http.StatusOK {
			t.Fatalf("Expected status code %v, but received: %v", "200", response.Code)
		}
		fmt.Println(response.Body)
	}
}

func readCSV(file string) *csv.Reader {
	// Open the file
	csvfile, err := os.Open(file)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	// Parse the file
	r := csv.NewReader(csvfile)
	// fmt.Printf("%T", r)
	return r
}
