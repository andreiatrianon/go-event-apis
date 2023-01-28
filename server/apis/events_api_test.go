package apis

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"bytes"
	"code-test/data"
	"encoding/json"
	"errors"
	"sync"
)

type MockEventsApiResponseHandler struct{}
var requestJson []byte
var expectedEvent data.Event

func TestPostEvent(t *testing.T) {
	t.Run("Events", func(t *testing.T) {

		t.Run("CopyAndPaste", func(t *testing.T) {
			requestJson = []byte(`{
				"eventType": "copyAndPaste",
				"websiteUrl": "https://ravelin.com",
				"sessionId": "123123-123123-123123123",
				"pasted": true,
				"formId": "inputCardNumber"
			}`)

			expectedEvent = data.Event{
				WebsiteUrl: "https://ravelin.com",
				SessionId: "123123-123123-123123123",
				CopyAndPaste: map[string]bool{
					"pasted": true,
					"inputCardNumber": true,
				},
			}

			runTestForPostEvent(t, requestJson, expectedEvent)
		})

		t.Run("ScreenResize", func(t *testing.T) {
			requestJson = []byte(`{
				"eventType": "screenResize",
				"websiteUrl": "https://ravelin.com",
				"sessionId": "123123-123123-123123123",
				"resizeFrom": {
					"width": "1920",
					"height": "1080"
				},
				"resizeTo": {
					"width": "1280",
					"height": "720"
				}
			}`)

			expectedEvent = data.Event{
				WebsiteUrl: "https://ravelin.com",
				SessionId: "123123-123123-123123123",
				ResizeFrom: data.Dimension{
					Width: "1920",
					Height: "1080",
				},
				ResizeTo: data.Dimension{
					Width: "1280",
					Height: "720",
				},
			}

			runTestForPostEvent(t, requestJson, expectedEvent)
		})

		t.Run("TimeTaken", func(t *testing.T) {
			requestJson = []byte(`{
				"eventType": "timeTaken",
				"websiteUrl": "https://ravelin.com",
				"sessionId": "123123-123123-123123123",
				"timeTaken": 72
			}`)

			expectedEvent = data.Event{
				WebsiteUrl: "https://ravelin.com",
				SessionId: "123123-123123-123123123",
				FormCompletionTime: 72,
			}

			runTestForPostEvent(t, requestJson, expectedEvent)
		})

	})
}

func TestErrorOnConvertingRequestBodyToStruct(t *testing.T) {
	// given
	ConvertRequestBodyToStructFunc = MockEventsApiResponseHandler{}.ConvertRequestBodyToStruct
	requestJson = []byte(`{}`)
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(requestJson))
	w := httptest.NewRecorder()
	wg := sync.WaitGroup{}
	wg.Add(1)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		EventsApiResponseHandler{}.PostEvent(w, req, &wg)
		wg.Wait()
	})

	// when
	handler.ServeHTTP(w, req)

	// then
	statusCodeResponse := w.Code
	expectedStatusCode := http.StatusInternalServerError

	if statusCodeResponse != expectedStatusCode {
		t.Errorf("Handler returned unexpected response status code:\nActual:\n%v\nExpected:\n%v",
		statusCodeResponse, expectedStatusCode)
	}
}

func (m MockEventsApiResponseHandler) ConvertRequestBodyToStruct(w http.ResponseWriter, r *http.Request) (data.Request, error) {
	var requestStruct data.Request
	mockResponseError := errors.New("Error on converting request body to a struct")
	return requestStruct, mockResponseError
}

func runTestForPostEvent(t *testing.T, requestJson []byte, expectedEvent data.Event) {
	// given
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(requestJson))
	w := httptest.NewRecorder()
	wg := sync.WaitGroup{}
	wg.Add(1)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		EventsApiResponseHandler{}.PostEvent(w, req, &wg)
		wg.Wait()
	})

	// when
	handler.ServeHTTP(w, req)

	// then
	responseBody := w.Body.String()

	var buffer bytes.Buffer
	json.NewEncoder(&buffer).Encode(&expectedEvent)
	expectedResponse := buffer.String()

	if responseBody != expectedResponse {
		t.Errorf("Handler returned unexpected response body:\n Actual:\n%v\nExpected:\n%v",
		responseBody, expectedResponse)
	}
}
