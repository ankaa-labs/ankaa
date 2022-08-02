package engine

type FailState struct {
	Type string // Fail

	Comment *string

	// required
	Error string
	Cause string
}
