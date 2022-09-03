package engine

import (
	"context"
	"log"
	"sync"
	"sync/atomic"
)

type ParallelState struct {
	// It must be "Parallel"
	Type string `json:"Type"`

	Comment string `json:"Comment"`

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

	// $.States
	// $.StartAt
	Branches []InnerStateMachine `json:"Branches"`
}

func (s *ParallelState) Execute(ctx context.Context, executeFn func(context.Context, string) error) error {
	log.Printf("execute[%v] ParallelState(%v)\n", atomic.AddInt32(&index, 1), s.Comment)

	var wg sync.WaitGroup
	wg.Add(len(s.Branches))

	output := make([]interface{}, len(s.Branches))

	for i := range s.Branches {
		go func(i int) {
			defer wg.Done()

			err := s.Branches[i].Execute(ctx)
			output[i] = err
		}(i)
	}

	wg.Wait()

	log.Printf("ParallelState(%v) output: %v\n", s.Comment, output)

	if s.End != nil && *s.End {
		return nil
	}

	return executeFn(ctx, *s.Next)
}
