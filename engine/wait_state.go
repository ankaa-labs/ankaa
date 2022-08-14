package engine

import "time"

type WaitState struct {
	Type string `json:"Type"` // Wait

	// One of (Next, End)
	Next *string `json:"Next"`
	End  *bool   `json:"End"`

	Comment *string `json:"Comment"`

	InputPath  *string `json:"InputPath"`
	OutputPath *string `json:"OutputPath"`

	// One of (sub 4)
	Seconds     *int    `json:"Seconds"`
	SecondsPath *string `json:"SecondsPath"`

	Timestamp     *time.Time `json:"Timestamp"`
	TimestampPath *string    `json:"TimestampPath"`
}
