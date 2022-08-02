package engine

type TaskState struct {
	// It must be "Task"
	Type string

	Resource string

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

	// States.Timeout
	TimeoutSeconds   *int // defaults 60
	HeartbeatSeconds *int // < TimeoutSeconds

	// Reference Paths
	TimeoutSecondsPath   *string // One of (TimeoutSeconds)
	HeartbeatSecondsPath *string // One of (HeartbeatSeconds)
}
