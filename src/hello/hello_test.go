package hello

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	request, _ := http.NewRequest("GET", "/hello", nil)
	response := httptest.NewRecorder()

	type HTTPTestCase struct {
		output string
	}

	testBody := []string{
		"Hello HTTP!",
	}

	HelloHandler(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Expected status code %v, but received: %v", "200", response.Code)
	}

	for _, v := range testBody {
		if response.Body.String() != v {
			t.Fatalf("Expected body: %v, but received: %v", v, response.Body)
		}
	}

}
