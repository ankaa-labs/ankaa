package engine

import (
	"encoding/json"
	"fmt"
)

type BaseStruct struct {
	Type string `json:"Type"`
}

func UnmarshalJSON(b []byte) (State, error) {
	bs := &BaseStruct{}
	err := json.Unmarshal(b, bs)
	if err != nil {
		return nil, err
	}

	var ss State
	switch bs.Type {
	case "Wait":
		ss = &WaitState{}
	case "Task":
		ss = &TaskState{}
	case "Succeed":
		ss = &SucceedState{}
	case "Pass":
		ss = &PassState{}
	case "Map":
		ss = &MapState{}
	case "Fail":
		ss = &FailState{}
	case "Choice":
		ss = &ChoiceState{}
	case "Parallel":
		ss = &ParallelState{}
	default:
		return nil, fmt.Errorf("unsupport type: %s", bs.Type)
	}

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
		o, err := UnmarshalJSON([]byte(v))
		if err != nil {
			return err
		}

		(*s)[k] = o
	}

	return nil
}
