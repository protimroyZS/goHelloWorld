package main

import (
	"net/http"
	"regexp"

	"./databaseTrial"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/", route) // Match everything
	http.ListenAndServe(":8081", nil)
}

func route(w http.ResponseWriter, r *http.Request) {

	var rNum = regexp.MustCompile(`hello/\d`) // Has digit(s)
	var rAbc = regexp.MustCompile(`hello/`)   // Contains "abc"
	switch {
	case rNum.MatchString(r.URL.Path):
		databaseTrial.PUTHandler(w, r)
	case rAbc.MatchString(r.URL.Path):
		databaseTrial.DBReqHandler(w, r)
	default:
		w.Write([]byte("Unknown Pattern"))
	}
}
