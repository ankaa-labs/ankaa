package engine

type PassState struct {
	Type    string // Pass
	Next    *string
	End     *string
	Comment *string

	InputPath  *string
	OutputPath *string

	// Payload Template
	Parameters *string

	Result     *string
	ResultPath *string
}
