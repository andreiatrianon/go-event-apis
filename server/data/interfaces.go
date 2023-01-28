package data

import (
	"net/http"
	"sync"
)

type EventsAPI interface {
	PostEvent(w http.ResponseWriter, r *http.Request, wg *sync.WaitGroup)
}

type HttpResponseHandler interface {
	ConvertRequestBodyToStruct(w http.ResponseWriter, r *http.Request) (Request, error)
}
