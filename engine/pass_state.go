package engine

type PassState struct {
	Type    string  `json:"Type"` // Pass
	Next    *string `json:"Next"`
	End     *bool   `json:"End"`
	Comment *string `json:"Comment"`

	InputPath  *string `json:"InputPath"`
	OutputPath *string `json:"OutputPath"`

	// Payload Template
	Parameters interface{} `json:"Parameters"`

	Result     interface{} `json:"Result"`
	ResultPath *string     `json:"ResultPath"`
}
