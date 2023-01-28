package main

import (
	"fmt"
	"net/http"
	"sync"
	"code-test/apis"
	"code-test/helper"
)

var Wg = sync.WaitGroup{}

func main() {
	http.HandleFunc("/", handleHttpResponse)

	http.ListenAndServe(":8080", nil)
}

func handleHttpResponse(w http.ResponseWriter, r *http.Request) {
	helper.EnableCors(&w)

	switch r.Method {
		case http.MethodOptions:
			w.WriteHeader(http.StatusOK)
			fmt.Println("Status code response OK for OPTIONS request")
			return
		case http.MethodPost:
			Wg.Add(1)
			go apis.PostEventAPI(w, r, &Wg)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Printf("Bad Request: method not allowed\nRequest: %v\nMethod: %v\n", r, r.Method)
			return
	}

	Wg.Wait()
}
