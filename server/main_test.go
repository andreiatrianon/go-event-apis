package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"bytes"
	"code-test/apis"
	"time"
	"sync"
)

type MockEventsApiResponseHandler struct{}
var requestMethod string
var expectedStatusCodeResponse int

func TestHandleHttpResponse(t *testing.T) {
	t.Run("Methods", func(t *testing.T) {

		t.Run("OPTIONS", func(t *testing.T) {

			requestMethod = http.MethodOptions
			expectedStatusCodeResponse = http.StatusOK

			runTestForHandlingHttpResponse(t, requestMethod, expectedStatusCodeResponse)
		})

		t.Run("POST", func(t *testing.T) {
			requestMethod = http.MethodPost
			expectedStatusCodeResponse = http.StatusCreated

			runTestForHandlingHttpResponse(t, requestMethod, expectedStatusCodeResponse)
		})

		t.Run("Not allowed method", func(t *testing.T) {
			requestMethod = http.MethodGet
			expectedStatusCodeResponse = http.StatusMethodNotAllowed

			runTestForHandlingHttpResponse(t, requestMethod, expectedStatusCodeResponse)
		})

	})
}

func TestHandleHttpResponseConcurrency(t *testing.T) {
	apis.PostEventAPI = MockEventsApiResponseHandler{}.PostEvent

	t.Run("Multiple requests for delayed post events", func(t *testing.T) {

		t.Run("1", func(t *testing.T) {
			requestMethod = http.MethodPost
			expectedStatusCodeResponse = http.StatusCreated

			runTestForHandlingHttpResponse(t, requestMethod, expectedStatusCodeResponse)
		})

		t.Run("2", func(t *testing.T) {
			requestMethod = http.MethodPost
			expectedStatusCodeResponse = http.StatusCreated

			runTestForHandlingHttpResponse(t, requestMethod, expectedStatusCodeResponse)
		})

		t.Run("3", func(t *testing.T) {
			requestMethod = http.MethodPost
			expectedStatusCodeResponse = http.StatusCreated

			runTestForHandlingHttpResponse(t, requestMethod, expectedStatusCodeResponse)
		})
	})
}

func (m MockEventsApiResponseHandler) PostEvent(w http.ResponseWriter, r *http.Request, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	w.WriteHeader(http.StatusCreated)
	wg.Done()
}

func runTestForHandlingHttpResponse(t *testing.T, requestMethod string, expectedStatusCodeResponse int) {
	// given
	requestJson := []byte(`{
		"eventType": "timeTaken",
		"websiteUrl": "https://ravelin.com",
		"sessionId": "123123-123123-123123123",
		"timeTaken": 72
	}`)
	req, _ := http.NewRequest(requestMethod, "/", bytes.NewBuffer(requestJson))
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(handleHttpResponse)

	// when
	handler.ServeHTTP(w, req)

	// then
	statusCodeResponse := w.Code

	if statusCodeResponse != expectedStatusCodeResponse {
		t.Errorf("Handler returned unexpected response status code:\nActual:\n%v\nExpected:\n%v",
		statusCodeResponse, expectedStatusCodeResponse)
	}
}
