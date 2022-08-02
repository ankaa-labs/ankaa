package engine

type SucceedState struct {
	Type string // Succeed

	Comment *string

	InputPath  *string
	OutputPath *string
}
