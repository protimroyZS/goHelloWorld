package hello

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestPOSTHandler(t *testing.T) {

	requestBodies := []HTTPPOSTTestCase{
		{"Pratim", 23},
		{"Roy", 24},
	}

	for _, v := range requestBodies {

		requestBody, _ := json.Marshal(&HTTPPOSTTestCase{
			Name: v.Name,
			Age:  v.Age,
		})
		request, _ := http.NewRequest("POST", "/hello", bytes.NewBuffer(requestBody))
		response := httptest.NewRecorder()

		POSTHandler(response, request)

		if response.Code != http.StatusOK {
			t.Fatalf("Expected status code %v, but received: %v", "200", response.Code)
		}

		expectedString := "Hello " + v.Name + "! \n Your Age is " + strconv.Itoa(v.Age)
		if response.Body.String() != expectedString {
			t.Fatalf("Expected body: %v, but received: %v", expectedString, response.Body)
		}
	}

}
