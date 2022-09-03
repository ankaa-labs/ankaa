package engine

import (
	"context"
	"fmt"
	"log"
	"sync/atomic"
)

type FailState struct {
	Type string `json:"Type"` // Fail

	Comment *string `json:"Comment"`

	// required
	Error string `json:"Error"`
	Cause string `json:"Cause"`
}

func (s *FailState) Execute(ctx context.Context, executeFn func(context.Context, string) error) error {
	log.Printf("execute[%v] FailState(%v)\n", atomic.AddInt32(&index, 1), s.Comment)
	return fmt.Errorf("end in fail state %v", s.Comment)
}
