package apis

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
	"code-test/data"
	"code-test/helper"
	"sync"
)

type EventsApiResponseHandler struct{}
var PostEventAPI = EventsApiResponseHandler{}.PostEvent
var ConvertRequestBodyToStructFunc = EventsApiResponseHandler{}.ConvertRequestBodyToStruct

func (e EventsApiResponseHandler) PostEvent(w http.ResponseWriter, r *http.Request, wg *sync.WaitGroup) {
	request, err := ConvertRequestBodyToStructFunc(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		wg.Done()
		return
	}
	fmt.Printf("Request struct: %v\n\n", request)

	event := helper.BuildNewEvent(request)

	fmt.Printf("completed Event struct: %v\n\n", event)

	helper.BuildResponse(w, event)

	fmt.Println("--------------------------------------------------------------------------")
	wg.Done()
}

func (e EventsApiResponseHandler) ConvertRequestBodyToStruct(w http.ResponseWriter, r *http.Request) (data.Request, error) {
	var requestStruct data.Request

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		return requestStruct, err
	}
	fmt.Printf("Request body: %s\n\n", requestBody)

	json.Unmarshal(requestBody, &requestStruct)
	return requestStruct, nil
}
