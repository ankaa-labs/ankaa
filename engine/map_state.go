package engine

import (
	"context"
	"log"
	"sync"
	"sync/atomic"

	"golang.org/x/sync/semaphore"
)

// Context Object will has .Map.Item.{Index,Value}
type MapState struct {
	// It must be "Map"
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

	// Reference Path, JSON array
	// defaults $
	ItemsPath *string `json:"ItemsPath"`

	// >= 0
	// defaults 0, no limit
	MaxConcurrency int `json:"MaxConcurrency"`

	// $.States
	// $.StartAt
	Iterator InnerStateMachine `json:"Iterator"`
}

func (s *MapState) Execute(ctx context.Context, executeFn func(context.Context, string) error) error {
	log.Printf("execute[%v] MapState(%v)\n", atomic.AddInt32(&index, 1), s.Comment)

	input := []string{"1", "2", "3", "4", "5"}
	output := make([]interface{}, len(input))

	maxConcurrency := s.MaxConcurrency
	if s.MaxConcurrency == 0 {
		maxConcurrency = len(input)
	}

	sem := semaphore.NewWeighted(int64(maxConcurrency))
	var wg sync.WaitGroup

	for i, _ := range input {
		if err := sem.Acquire(ctx, 1); err != nil {
			return err
		}

		wg.Add(1)
		go func(i int) {
			defer sem.Release(1)
			defer wg.Done()

			err := s.Iterator.Execute(ctx)
			output[i] = err
		}(i)
	}

	// Wait all have done
	wg.Wait()

	log.Printf("MapState(%v) output: %v\n", s.Comment, output)

	if s.End != nil && *s.End {
		return nil
	}

	return executeFn(ctx, *s.Next)
}
