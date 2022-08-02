package engine

// Context Object will has .Map.Item.{Index,Value}
type MapState struct {
	// It must be "Map"
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

	// Reference Path, JSON array
	// defaults $
	ItemsPath *string

	// >= 0
	// defaults 0, no limit
	MaxConcurrency *int

	// $.States
	// $.StartAt
	Iterator *string
}
