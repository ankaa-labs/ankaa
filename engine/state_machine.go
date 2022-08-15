package engine

import (
	"encoding/json"
	"fmt"
)

// BaseStruct is the struct for unmarshal State's Type field
type BaseStruct struct {
	Type string `json:"Type"`
}

// StateFactory to construct a type reference State object
type StateFactory func() State

// StateType present the type of State object
type StateType = string

const (
	// StateTypeWait is the type of Wait State
	StateTypeWait StateType = "Wait"

	// StateTypeTask is the type of Task State
	StateTypeTask StateType = "Task"

	// StateTypeSucceed is the type of Succeed State
	StateTypeSucceed StateType = "Succeed"

	// StateTypePass is the type of Pass State
	StateTypePass StateType = "Pass"

	// StateTypeMap is the type of Map State
	StateTypeMap StateType = "Map"

	// StateTypeFail is the type of Fail State
	StateTypeFail StateType = "Fail"

	// StateTypeChoice is the type of Choice State
	StateTypeChoice StateType = "Choice"

	// StateTypeParallel is the type of Parallel State
	StateTypeParallel StateType = "Parallel"
)

var (
	stateFactories = map[StateType]StateFactory{
		StateTypeWait: func() State {
			return &WaitState{}
		},
		StateTypeTask: func() State {
			return &TaskState{}
		},
		StateTypeSucceed: func() State {
			return &SucceedState{}
		},
		StateTypePass: func() State {
			return &PassState{}
		},
		StateTypeMap: func() State {
			return &MapState{}
		},
		StateTypeFail: func() State {
			return &FailState{}
		},
		StateTypeChoice: func() State {
			return &ChoiceState{}
		},
		StateTypeParallel: func() State {
			return &ParallelState{}
		},
	}
)

// UnmarshalStateFromJSON will unmarshal the json's byte slice to State object
func UnmarshalStateFromJSON(b []byte) (State, error) {
	bs := &BaseStruct{}
	err := json.Unmarshal(b, bs)
	if err != nil {
		return nil, err
	}

	objectFactory, ok := stateFactories[bs.Type]
	if !ok {
		return nil, fmt.Errorf("unsupport state type: %s", bs.Type)
	}

	ss := objectFactory()
	err = json.Unmarshal(b, ss)
	if err != nil {
		return nil, err
	}
	return ss, nil
}

type StateMachine struct {
	Comment        *string `json:"Comment"`
	StartAt        string  `json:"StartAt"`
	Version        string  `json:"Version"`
	TimeoutSeconds *int    `json:"TimeoutSeconds"`

	States map[string]State `json:"States"`
}

type InnerStateMachine struct {
	StartAt string           `json:"StartAt"`
	States  map[string]State `json:"States"`
}

type State interface {
}

type States map[string]State

func (s *States) UnmarshalJSON(b []byte) error {
	rawStatesMap := map[string]json.RawMessage{}
	err := json.Unmarshal(b, &rawStatesMap)
	if err != nil {
		return err
	}

	*s = make(map[string]State, len(rawStatesMap))
	for k, v := range rawStatesMap {
		o, err := UnmarshalStateFromJSON([]byte(v))
		if err != nil {
			return err
		}

		(*s)[k] = o
	}

	return nil
}
