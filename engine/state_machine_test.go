package engine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"testing"

	"github.com/davecgh/go-spew/spew"
	. "github.com/onsi/gomega"

	"github.com/ankaa-labs/ankaa/test"
	"github.com/ankaa-labs/ankaa/utils/pointer"
)

func TestUnmarshalStateFromJSON(t *testing.T) {
	type testCase struct {
		desp        string
		b           string
		expectState State
		err         string
	}

	testCases := []testCase{
		{
			desp:        "BaseStruct unmarshal failed",
			b:           `{"Type": 1}`,
			expectState: nil,
			err:         "cannot unmarshal number into Go struct field BaseStruct.Type of type strin",
		},
		{
			desp: "normal Wait state",
			b: `{
				"Type": "Wait",
				"Next": "w"
			}`,
			expectState: &WaitState{
				Type: StateTypeWait,
				Next: pointer.StringPtr("w"),
			},
			err: ``,
		},
		{
			desp: "Wait unmarshal failed",
			b: `{
				"Type": "Wait",
				"Next": 1
			}`,
			expectState: nil,
			err:         `cannot unmarshal number into Go struct field WaitState.Next of type string`,
		},
		{
			desp: "normal Task state",
			b:    `{"Type": "Task", "Resource": "r"}`,
			expectState: &TaskState{
				Type:     StateTypeTask,
				Resource: "r",
			},
			err: ``,
		},
		{
			desp:        "Task unmarshal failed",
			b:           `{"Type": "Task", "Resource": 1}`,
			expectState: nil,
			err:         `cannot unmarshal number into Go struct field TaskState.Resource of type string`,
		},
		{
			desp: "normal Succeed state",
			b:    `{"Type": "Succeed", "InputPath": "r"}`,
			expectState: &SucceedState{
				Type:      StateTypeSucceed,
				InputPath: pointer.StringPtr("r"),
			},
			err: ``,
		},
		{
			desp:        "Succeed unmarshal failed",
			b:           `{"Type": "Succeed", "InputPath": 1}`,
			expectState: nil,
			err:         `cannot unmarshal number into Go struct field SucceedState.InputPath of type string`,
		},
		{
			desp: "normal Pass state",
			b:    `{"Type": "Pass", "InputPath": "r"}`,
			expectState: &PassState{
				Type:      StateTypePass,
				InputPath: pointer.StringPtr("r"),
			},
			err: ``,
		},
		{
			desp:        "Pass unmarshal failed",
			b:           `{"Type": "Pass", "InputPath": 1}`,
			expectState: nil,
			err:         `cannot unmarshal number into Go struct field PassState.InputPath of type string`,
		},
		{
			desp: "normal Map state",
			b:    `{"Type": "Map", "InputPath": "r"}`,
			expectState: &MapState{
				Type:      StateTypeMap,
				InputPath: pointer.StringPtr("r"),
			},
			err: ``,
		},
		{
			desp:        "Map unmarshal failed",
			b:           `{"Type": "Map", "InputPath": 1}`,
			expectState: nil,
			err:         `cannot unmarshal number into Go struct field MapState.InputPath of type string`,
		},
		{
			desp: "normal Fail state",
			b:    `{"Type": "Fail", "Comment": "r"}`,
			expectState: &FailState{
				Type:    StateTypeFail,
				Comment: pointer.StringPtr("r"),
			},
			err: ``,
		},
		{
			desp:        "Fail unmarshal failed",
			b:           `{"Type": "Fail", "Comment": 1}`,
			expectState: nil,
			err:         `cannot unmarshal number into Go struct field FailState.Comment of type string`,
		},
		{
			desp: "normal Choice state",
			b:    `{"Type": "Choice", "InputPath": "r"}`,
			expectState: &ChoiceState{
				Type:      StateTypeChoice,
				InputPath: pointer.StringPtr("r"),
			},
			err: ``,
		},
		{
			desp:        "Choice unmarshal failed",
			b:           `{"Type": "Choice", "InputPath": 1}`,
			expectState: nil,
			err:         `cannot unmarshal number into Go struct field ChoiceState.InputPath of type string`,
		},
		{
			desp: "normal Parallel state",
			b:    `{"Type": "Parallel", "InputPath": "r"}`,
			expectState: &ParallelState{
				Type:      StateTypeParallel,
				InputPath: pointer.StringPtr("r"),
			},
			err: ``,
		},
		{
			desp:        "Parallel unmarshal failed",
			b:           `{"Type": "Parallel", "InputPath": 1}`,
			expectState: nil,
			err:         `cannot unmarshal number into Go struct field ParallelState.InputPath of type string`,
		},
		{
			desp:        "unsupport state type",
			b:           `{"Type": "Parallel1", "InputPath": "r"}`,
			expectState: nil,
			err:         `unsupport state type: Parallel1`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desp, func(t *testing.T) {
			g := NewWithT(t)

			actual, err := UnmarshalStateFromJSON([]byte(tc.b))
			if tc.err != "" {
				g.Expect(err).To(HaveOccurred())
				g.Expect(err.Error()).To(MatchRegexp(tc.err))
				g.Expect(actual).To(BeNil())
				return
			}

			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(actual).To(Equal(tc.expectState))
		})
	}
}

func TestStateMachineUnmarshal(t *testing.T) {
	g := NewWithT(t)

	f := path.Join(test.CurrentProjectPath() + "/hack/asl/map-with-null-itemspath.json")
	data, err := ioutil.ReadFile(f)
	g.Expect(err).ToNot(HaveOccurred())

	s := &StateMachine{}
	err = json.Unmarshal(data, s)
	g.Expect(err).ToNot(HaveOccurred())

	spew.Config.Indent = "\t"
	spew.Dump(s)
	fmt.Printf("%+v\n", spew.Sprint(s))
}
