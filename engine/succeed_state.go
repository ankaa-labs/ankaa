package engine

type SucceedState struct {
	Type string `json:"Type"` // Succeed

	Comment *string `json:"Comment"`

	InputPath  *string `json:"InputPath"`
	OutputPath *string `json:"OutputPath"`
}
