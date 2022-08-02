package engine

type ParallelState struct {
	// It must be "Parallel"
	Type string

	Comment *string

	InputPath  *string
	OutputPath *string

	ResultPath *string

	Retry []Retrier
	Catch []Catcher

	// One of (Next, End)
	Next *string
	End  *bool

	// Payload Template
	Parameters     *string
	ResultSelector *string

	// $.States
	// $.StartAt
	Branches []*string
}
