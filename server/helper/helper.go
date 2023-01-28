package helper

import (
	"fmt"
	"net/http"
	"encoding/json"
	"code-test/data"
)

func EnableCors(w *http.ResponseWriter) {
	// See more details about enabling CORS in Go:
	// https://www.stackhawk.com/blog/golang-cors-guide-what-it-is-and-how-to-enable-it/

	// TODO: fix the header Access-Control-Allow-Origin to accept localhost
	(*w).Header().Set("Access-Control-Allow-Origin", "*")

	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func BuildNewEvent(request data.Request) data.Event {
	newEvent := data.Event{
		WebsiteUrl: request.WebsiteUrl,
		SessionId: request.SessionId,
	}
	fmt.Printf("partial Event struct: %v\n\n", newEvent)

	switch request.EventType {
		case "copyAndPaste":
			newEvent.CopyAndPaste = map[string]bool{
				"pasted": true,
				request.FormId: true,
			}
			fmt.Println("CopyAndPaste event")

		case "screenResize":
			newEvent.ResizeFrom = request.ResizeFrom
			newEvent.ResizeTo = request.ResizeTo
			fmt.Println("ScreenResize event")

		case "timeTaken":
			newEvent.FormCompletionTime = request.TimeTaken
			fmt.Println("TimeTaken event")

	}

	return newEvent
}

func BuildResponse(w http.ResponseWriter, event data.Event) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
}
