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
	"time"

	"github.com/serverlessworkflow/sdk-go/v2/model"
)

type WorkflowInstance struct {
	id      string
	version string

	instanceId string
	startAt    time.Time
	status     string
}

type WorkflowActivity struct {
	id        string
	stateName string
	startAt   time.Time
}

type Storage struct {
	workflowDefinitions []*model.Workflow
	workflowInstances   []*WorkflowInstance
	workflowActivitys   []*WorkflowActivity
}

var (
	storage = &Storage{
		workflowDefinitions: make([]*model.Workflow, 0),
		workflowInstances:   make([]*WorkflowInstance, 0),
		workflowActivitys:   make([]*WorkflowActivity, 0),
	}
)

func (s *Storage) AddWorkflowDefinition(m *model.Workflow) {
	s.workflowDefinitions = append(s.workflowDefinitions, m)
}

func (s *Storage) GetWorkflowDefinition(id string, version string) *model.Workflow {
	for _, m := range s.workflowDefinitions {
		if m.ID == id && m.Version == version {
			return m
		}
	}

	return nil
}

func (s *Storage) AddWorkflowInstance(v *WorkflowInstance) {
	s.workflowInstances = append(s.workflowInstances, v)
}

func (s *Storage) GetWorkflowInstance(id string) *WorkflowInstance {
	for _, v := range s.workflowInstances {
		if v.instanceId == id {
			return v
		}
	}

	return nil
}

func (s *Storage) AddWorkflowActivity(v *WorkflowActivity) {
	s.workflowActivitys = append(s.workflowActivitys, v)
}

func (s *Storage) GetWorkflowActivity(id string, stateName string) *WorkflowActivity {
	for _, v := range s.workflowActivitys {
		if v.id == id && v.stateName == stateName {
			return v
		}
	}

	return nil
}
