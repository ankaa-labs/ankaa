package engine

// Catcher ...
type Catcher struct {
	// +Required
	// >= 1
	// States.ALL must the last and only elem
	ErrorEquals []string

	// +Required
	Next string

	// +Optional
	// default $
	ResultPath *string
}
