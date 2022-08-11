package engine

type States = map[string]State

type State struct {
	*ChoiceState
	*FailState
	*SucceedState
	*TaskState
	// *ParallelState
	*PassState
	*WaitState
	*MapState
}

type StateMachine struct {
	Comment        *string `json:"Comment"`
	StartAt        string  `json:"StartAt"`
	Version        string  `json:"Version"`
	TimeoutSeconds *int    `json:"TimeoutSeconds"`

	States States `json:"States"`
}
