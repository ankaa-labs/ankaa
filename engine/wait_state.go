package engine

import (
	"fmt"
	"time"
)

// WaitState causes the interpreter to delay the machine from continuing for a specified time.
type WaitState struct {
	// Type is the type of WaitState
	// +Required
	Type string `json:"Type"`

	// Next is the state to run next
	// +Optional
	Next *string `json:"Next"`

	// End represent where the state is Terminal State
	// +Optional
	// Defaults to false
	End bool `json:"End"`

	// Comment provided for human-readable description of the machine
	// +Optional
	Comment string `json:"Comment"`

	// InputPath MUST be a Path, which is applied to a State's raw input to select some or all of it
	// +Optional
	InputPath *string `json:"InputPath"`

	// OutputPath MUST be a Path, which is applied to the state's output after the application of ResultPath
	// +Optional
	OutputPath *string `json:"OutputPath"`

	// Seconds represent the wait duration in seconds
	// +Optional
	Seconds *int `json:"Seconds"`

	// SecondsPath be a Reference Path to Seconds value
	// +Optional
	SecondsPath *string `json:"SecondsPath"`

	// Timestamp represent the wait absolute expiry time
	// +Optional
	Timestamp *time.Time `json:"Timestamp"`

	// TimestampPath be a Reference Path to Timestamp value
	// +Optional
	TimestampPath *string `json:"TimestampPath"`
}

// Validate will validate the WaitState configuration
func (s *WaitState) Validate() error {
	if err := ValidateTerminalState(s.Next, s.End); err != nil {
		return err
	}

	if err := ValidatePath(s.InputPath); err != nil {
		return err
	}

	if err := ValidatePath(s.OutputPath); err != nil {
		return err
	}

	counts := []bool{
		s.Seconds != nil,
		s.SecondsPath != nil,
		s.Timestamp != nil,
		s.TimestampPath != nil,
	}
	v := 0
	for _, b := range counts {
		if b {
			v++
		}
	}
	if v != 1 {
		return fmt.Errorf("the path must be 1, but got %v", v)
	}

	if s.Seconds != nil && *(s.Seconds) < 0 {
		return fmt.Errorf("seconds must greater than or equal to 0")
	}

	return nil
}
