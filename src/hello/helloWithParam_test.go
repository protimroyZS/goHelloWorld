package hello

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNameHandler(t *testing.T) {

	type HTTPTestCase struct {
		input  string
		output string
	}

	testBody := []HTTPTestCase{
		{"", "Hello HTTP!"},
		{"Pratim", "Hello Pratim!"},
	}

	for _, v := range testBody {

		log.Print("/hello?name=" + v.input)
		request, _ := http.NewRequest("GET", "/hello?name="+v.input, nil)
		response := httptest.NewRecorder()

		NameHandler(response, request)

		if response.Code != http.StatusOK {
			t.Fatalf("Expected status code %v, but received: %v", "200", response.Code)
		}

		if response.Body.String() != v.output {
			t.Fatalf("Expected body: %v, but received: %v", v.output, response.Body)
		}
	}

}
