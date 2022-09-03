package engine

import (
	"context"
	"log"
	"sync/atomic"
)

type PassState struct {
	Type    string  `json:"Type"` // Pass
	Next    *string `json:"Next"`
	End     *bool   `json:"End"`
	Comment string  `json:"Comment"`

	InputPath  *string `json:"InputPath"`
	OutputPath *string `json:"OutputPath"`

	// Payload Template
	Parameters interface{} `json:"Parameters"`

	Result     interface{} `json:"Result"`
	ResultPath *string     `json:"ResultPath"`
}

func (s *PassState) Execute(ctx context.Context, executeFn func(context.Context, string) error) error {
	log.Printf("execute[%v] PassState(%v)\n", atomic.AddInt32(&index, 1), s.Comment)
	if s.End != nil && *s.End {
		return nil
	}

	return executeFn(ctx, *s.Next)
}
