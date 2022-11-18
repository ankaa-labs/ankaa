// Copyright (c) 2022 The ankaa-labs Authors
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package engine

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"github.com/serverlessworkflow/sdk-go/v2/model"
)

type WorkflowInstanceStarting struct {
	pos       int64
	sourcePos int64

	id         string
	version    string
	instanceId string
	createdAt  time.Time
}

// WorkflowInstanceStarted ...
type WorkflowInstanceStarted struct {
	pos       int64
	sourcePos int64

	id string
}

type WorkflowInstanceCompleting struct {
	pos       int64
	sourcePos int64

	id          string
	completedAt time.Time
}

type WorkflowInstanceCompleted struct {
	pos       int64
	sourcePos int64

	id string
}

type WorkflowActivityStarting struct {
	pos       int64
	sourcePos int64

	id        string
	state     string
	createdAt time.Time
}

type WorkflowActivityStarted struct {
	pos       int64
	sourcePos int64

	id    string
	state string
}

func HandleEvent(obj interface{}) {
	switch obj.(type) {
	case *WorkflowInstanceStarting:
		HandleWorkflowInstanceStarting(obj.(*WorkflowInstanceStarting))
	case *WorkflowInstanceStarted:
		HandleWorkflowInstanceStarted(obj.(*WorkflowInstanceStarted))
	case *WorkflowInstanceCompleting:
		HandleWorkflowInstanceCompleting(obj.(*WorkflowInstanceCompleting))
	case *WorkflowInstanceCompleted:
		HandleWorkflowInstanceCompleted(obj.(*WorkflowInstanceCompleted))
	case *WorkflowActivityStarting:
		HandleWorkflowActivityStarting(obj.(*WorkflowActivityStarting))
	case *WorkflowActivityStarted:
		HandleWorkflowActivityStarted(obj.(*WorkflowActivityStarted))
	}
}

func HandleWorkflowInstanceCompleting(obj *WorkflowInstanceCompleting) {
	fmt.Println("start HandleWorkflowInstanceCompleting")
	defer fmt.Println("done HandleWorkflowInstanceCompleting")

	workflowInstance := storage.GetWorkflowInstance(obj.id)
	if workflowInstance == nil {
		fmt.Printf("cannot find workflow instance: %v\n", obj.id)
		return
	}

	workflowInstance.status = "done"

	// Seconds, after the information have been persistent,
	// we emit an event to indicate it have been processed.
	fmt.Println("publish WorkflowInstanceCompleted")
	stream.Publish(&WorkflowInstanceCompleted{
		pos:       atomic.AddInt64(&pos, 1),
		sourcePos: obj.pos,

		id: obj.id,
	})
}

func HandleWorkflowInstanceCompleted(obj *WorkflowInstanceCompleted) {
	fmt.Println("start HandleWorkflowInstanceCompleted")
	defer fmt.Println("done HandleWorkflowInstanceCompleted")

	// do nothing
}

func HandleWorkflowInstanceStarting(obj *WorkflowInstanceStarting) {
	fmt.Println("start HandleWorkflowInstanceStarting")
	defer fmt.Println("done HandleWorkflowInstanceStarting")

	// First, we save it to storage
	storage.AddWorkflowInstance(&WorkflowInstance{
		id:         obj.id,
		version:    obj.version,
		instanceId: obj.instanceId,
		startAt:    obj.createdAt,
		status:     "Pending",
	})

	// Seconds, after the information have been persistent,
	// we emit an event to indicate it have been processed.
	fmt.Println("publish WorkflowInstanceStarted")
	stream.Publish(&WorkflowInstanceStarted{
		pos:       atomic.AddInt64(&pos, 1),
		sourcePos: obj.pos,

		id: obj.instanceId,
	})
}

func HandleWorkflowInstanceStarted(obj *WorkflowInstanceStarted) {
	fmt.Println("start HandleWorkflowInstanceStarted")
	defer fmt.Println("done HandleWorkflowInstanceStarted")

	workflowInstance := storage.GetWorkflowInstance(obj.id)
	if workflowInstance == nil {
		fmt.Printf("cannot find workflow instance: %v\n", obj.id)
		return
	}

	workflowDefinition := storage.GetWorkflowDefinition(workflowInstance.id, workflowInstance.version)
	if workflowDefinition == nil {
		fmt.Printf("cannot find workflow definition: %v %v\n", workflowInstance.id, workflowInstance.version)
		return
	}

	if workflowDefinition.Start != nil {
		fmt.Println("publish WorkflowActivityStarting event with start")
		stream.Publish(
			&WorkflowActivityStarting{
				pos:       atomic.AddInt64(&pos, 1),
				sourcePos: obj.sourcePos,

				id:        obj.id,
				state:     workflowDefinition.Start.StateName,
				createdAt: time.Now(),
			},
		)

		return
	}

	fmt.Println("publish WorkflowActivityStarting event with first state")
	stream.Publish(
		&WorkflowActivityStarting{
			pos:       atomic.AddInt64(&pos, 1),
			sourcePos: obj.sourcePos,

			id:        obj.id,
			state:     workflowDefinition.States[0].GetName(),
			createdAt: time.Now(),
		},
	)
}

func HandleWorkflowActivityStarting(obj *WorkflowActivityStarting) {
	fmt.Printf("start HandleWorkflowActivityStarting\n")
	defer fmt.Printf("done HandleWorkflowActivityStarting\n")

	workflowInstance := storage.GetWorkflowInstance(obj.id)
	if workflowInstance == nil {
		fmt.Printf("cannot find workflow instance: %v\n", obj.id)
		return
	}

	workflowDefinition := storage.GetWorkflowDefinition(workflowInstance.id, workflowInstance.version)
	if workflowDefinition == nil {
		fmt.Printf("cannot find workflow definition: %v %v\n", workflowInstance.id, workflowInstance.version)
		return
	}

	var state model.State
	for _, s := range workflowDefinition.States {
		if s.GetName() == obj.state {
			state = s
		}
	}

	if state == nil {
		fmt.Printf("cannot find state %v in workflow definition: %v %v\n", obj.state, workflowInstance.id, workflowInstance.version)
		return
	}

	// First, we save it to storage
	storage.AddWorkflowActivity(&WorkflowActivity{
		id:        obj.id,
		stateName: obj.state,
		startAt:   obj.createdAt,
	})

	// Seconds, we emit an event to indicate it have been processed.
	fmt.Println("publish WorkflowActivityStarted")
	stream.Publish(&WorkflowActivityStarted{
		pos:       atomic.AddInt64(&pos, 1),
		sourcePos: obj.pos,

		id:    obj.id,
		state: obj.state,
	})
}

func HandleWorkflowActivityStarted(obj *WorkflowActivityStarted) {
	fmt.Println("start HandleWorkflowActivityStarted")
	defer fmt.Println("done HandleWorkflowActivityStarted")

	workflowInstance := storage.GetWorkflowInstance(obj.id)
	if workflowInstance == nil {
		fmt.Printf("cannot find workflow instance: %v\n", obj.id)
		return
	}

	workflowDefinition := storage.GetWorkflowDefinition(workflowInstance.id, workflowInstance.version)
	if workflowDefinition == nil {
		fmt.Printf("cannot find workflow definition: %v %v\n", workflowInstance.id, workflowInstance.version)
		return
	}

	var state model.State
	for _, s := range workflowDefinition.States {
		if s.GetName() == obj.state {
			state = s
		}
	}

	if state == nil {
		fmt.Printf("cannot find state %v in workflow definition: %v %v\n", obj.state, workflowInstance.id, workflowInstance.version)
		return
	}

	// This shoule be in completing event (we haven't implement it)
	if state.GetEnd() != nil {
		fmt.Println("publish WorkflowInstanceCompleting event")
		stream.Publish(
			&WorkflowInstanceCompleting{
				pos:       atomic.AddInt64(&pos, 1),
				sourcePos: obj.pos,

				id:          obj.id,
				completedAt: time.Now(),
			},
		)
		return
	}

	fmt.Println("publish WorkflowActivityStarting event")

	fmt.Printf("change state '%v' to state '%v'\n", obj.state, state.GetTransition().NextState)
	stream.Publish(
		&WorkflowActivityStarting{
			pos:       atomic.AddInt64(&pos, 1),
			sourcePos: obj.pos,

			id:        obj.id,
			state:     state.GetTransition().NextState,
			createdAt: time.Now(),
		},
	)
}

var (
	pos int64 = -1
)

var (
	triggers = make(map[string]chan interface{})
)

// Execute ...
func Execute(m *model.Workflow) {
	storage.AddWorkflowDefinition(m)

	fmt.Println("publish WorkflowInstanceStarting event")

	v := &WorkflowInstanceStarting{
		sourcePos: -1,
		pos:       atomic.AddInt64(&pos, 1),

		instanceId: uuid.New().String(),
		id:         m.ID,
		version:    m.Version,
		createdAt:  time.Now(),
	}
	triggers[v.instanceId] = make(chan interface{})

	stream.On(func(obj interface{}) {
		switch obj.(type) {
		case *WorkflowInstanceCompleted:
			trigger, ok := triggers[obj.(*WorkflowInstanceCompleted).id]
			if !ok {
				return
			}

			trigger <- time.Now()
		}
	})

	stream.Publish(v)

	vv := <-triggers[v.instanceId]
	fmt.Printf("trigger: %v\n", vv)
}
