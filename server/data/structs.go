package data

type Event struct {
	WebsiteUrl string `json:"websiteUrl"`
	SessionId string `json:"sessionId"`
	ResizeFrom Dimension `json:"resizeFrom"`
	ResizeTo Dimension `json:"resizeTo"`
	CopyAndPaste map[string]bool `json:"copyAndPaste"` // map[fieldId]true
	FormCompletionTime int `json:"formCompletionTime"` // Seconds
}

type Dimension struct {
	Width string `json:"width"`
	Height string `json:"height"`
}

type Request struct {
	EventType	string `json:"eventType"`
	WebsiteUrl string `json:"websiteUrl"`
	SessionId string `json:"sessionId"`
	Pasted	bool `json:"pasted"`
	FormId	string `json:"formId"`
	ResizeFrom Dimension `json:"resizeFrom"`
	ResizeTo Dimension `json:"resizeTo"`
	TimeTaken int `json:"timeTaken"`
}
