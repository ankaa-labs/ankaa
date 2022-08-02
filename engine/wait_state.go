package engine

import "time"

type WaitState struct {
	Type string // Wait

	// One of (Next, End)
	Next *string
	End  *bool

	Comment *string

	InputPath  *string
	OutputPath *string

	// One of (sub 4)
	Seconds     *int
	SecondsPath *string

	Timestamp     *time.Time
	TimestampPath *string
}
