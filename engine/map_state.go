package engine

// Context Object will has .Map.Item.{Index,Value}
type MapState struct {
	// It must be "Map"
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

	// Reference Path, JSON array
	// defaults $
	ItemsPath *string `json:"ItemsPath"`

	// >= 0
	// defaults 0, no limit
	MaxConcurrency *int `json:"MaxConcurrency"`

	// $.States
	// $.StartAt
	Iterator States `json:"Iterator"`
}
