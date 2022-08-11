package engine

type States = map[string]State

type State struct {
	*ChoiceState
	*FailState
	*SucceedState
	*TaskState
	*ParallelState
	*PassState
	*WaitState
	*MapState
}

type StateMachine struct {
	Comment *string
	StartAt string
	States  States
}
