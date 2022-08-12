package engine

type ParallelState struct {
	// It must be "Parallel"
	Type string `json:"Type"`

	Comment *string `json:"Comment"`

	InputPath  *string `json:"InputPath"`
	OutputPath *string `json:"OutputPath"`

	ResultPath *string `json:"ResultPath"`

	Retry []Retrier `json:"Retry"`
	Catch []Catcher `json:"Catch"`

	// One of (Next, End)
	Next *string `json:"Next"`
	End  *bool   `json:"End"`

	// Payload Template
	Parameters     interface{} `json:"Parameters"`
	ResultSelector interface{} `json:"ResultSelector"`

	// $.States
	// $.StartAt
	Branches []InnerStateMachine `json:"Branches"`
}
