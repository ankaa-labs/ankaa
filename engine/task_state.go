package engine

type TaskState struct {
	// It must be "Task"
	Type string `json:"Type"`

	Resource string `json:"Resource"`

	Comment *string `json:"Comment"`

	InputPath  *string `json:"InputPath"`
	OutputPath *string `json:"OutputPath"`

	ResultPath *string `json:"ResultPath"`

	Retry []Retrier `json:"Retry"`
	Catch []Catcher `json:"Catch"`

	// One of (Next, End)
	Next *string `json:"Next"`
	End  *bool   `json:"End"`

	// Payload Template, it could be any JSON object
	Parameters     interface{} `json:"Parameters"`
	ResultSelector interface{} `json:"ResultSelector"`

	// States.Timeout
	TimeoutSeconds   *int `json:"TimeoutSeconds"`   // defaults 60
	HeartbeatSeconds *int `json:"HeartbeatSeconds"` // < TimeoutSeconds

	// Reference Paths
	TimeoutSecondsPath   *string `json:"TimeoutSecondsPath"`   // One of (TimeoutSeconds)
	HeartbeatSecondsPath *string `json:"HeartbeatSecondsPath"` // One of (HeartbeatSeconds)
}
