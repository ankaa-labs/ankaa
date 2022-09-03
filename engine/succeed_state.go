package engine

import (
	"context"
	"log"
	"sync/atomic"
)

type SucceedState struct {
	Type string `json:"Type"` // Succeed

	Comment string `json:"Comment"`

	InputPath  *string `json:"InputPath"`
	OutputPath *string `json:"OutputPath"`
}

func (s *SucceedState) Execute(ctx context.Context, executeFn func(context.Context, string) error) error {
	log.Printf("execute[%v] SucceedState(%v)\n", atomic.AddInt32(&index, 1), s.Comment)
	return nil
}
