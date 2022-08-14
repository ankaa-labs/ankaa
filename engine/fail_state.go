package engine

type FailState struct {
	Type string `json:"Type"` // Fail

	Comment *string `json:"Comment"`

	// required
	Error string `json:"Error"`
	Cause string `json:"Cause"`
}
