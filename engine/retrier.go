package engine

// Retrier ...
type Retrier struct {
	// +Required
	// >= 1
	// States.ALL must the last and only elem
	ErrorEquals []string `json:"ErrorEquals"`

	// +Optional
	// > 0
	// default 1
	IntervalSeconds int `json:"IntervalSeconds"`

	// +Optional
	// >= 0
	// default 3
	MaxAttempts int `json:"MaxAttempts"`

	// +Optional
	// >= 1.0
	// default 2.0
	BackoffRate float32 `json:"BackoffRate"`
}
