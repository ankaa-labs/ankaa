package engine

// Catcher ...
type Catcher struct {
	// +Required
	// >= 1
	// States.ALL must the last and only elem
	ErrorEquals []string `json:"ErrorEquals"`

	// +Required
	Next string `json:"Next"`

	// +Optional
	// default $
	ResultPath *string `json:"ResultPath"`
}
