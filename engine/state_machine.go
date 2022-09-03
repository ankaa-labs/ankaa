package engine

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

var (
	index = int32(0)
)

// temporal EventType: https://docs.temporal.io/references/events/
// 1. WorkflowExecutionStarted
// 2. WorkflowExecutionSignaled
// 3. WorkflowTaskScheduled
// 4. WorkflowTaskStarted
// 5. WorkflowTaskCompleted
// 6. MarkerRecorded
// 7. TimerStarted
// 8. ActivityTaskScheduled
// 9. ActivityTaskStarted
// 10. ActivityTaskCompleted
// 11. WorkflowTaskScheduled
// 12. WorkflowTaskStarted
// 13. WorkflowTaskCompleted
// 14.

// zeebe https://docs.camunda.io/docs/0.26/components/zeebe/technical-concepts/workflow-lifecycles/

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

	States States `json:"States"`
}

// 1. 写 WorkflowStarting 事件 -> log 持久化
// 2. 得到 Starting 事件，修改本地状态(生成实例信息)，发送 Started 事件 -> log 持久化
// 3. 得到 Started 事件，发送 StateActivting 事件 -> log 持久化
// 4. 得到 Activting 事件，修改本地状态，发送 Activted 事件 -> log 持久化
// 如上的问题是：
// 1. 2/3/4 步骤中，如果修改本地状态成功，发送事件失败怎么处理？(所以应该是在发送事件之前，不做任何的本地修改？)
// 2. 如何处理重复收到某事件的问题？(记录下已经收到的事件的 index 即可, 要求 index 有序)

// 假设：
// 1. 用户发送启动流程实例的命令
// 2. broker 收到命令，appendLog1
// 3. 所有 broker 应用 log1
// 4. leader 应用 log1 成功，生成实例信息（此时是否保存到 db?），appendLog2（follower 不做任何处理，只记录 log）
//    4.1 如果保存，那么后续 leader 再收到这条 log 时只需要记录一下已经处理即可
//    4.2 如果不保存，那么后续 leader 再收到这条 log 时需要保存记录
//    4.3 问题在于，如何做恢复？假设生成实例信息，appendLog 成功之前 leader crash，新的 leader 如何判断是否还要生成实例信息？以及如何生成一模一样的实例信息？（因为 leader 可能已经返回信息给客户端）
// 5. appendLog1 成功则返回给客户端（leader）
// 5. 所有 broker 应用 log2
// 6. leader

// 在 workflow 中，和 herddb 之类的不一样，每个命令的操作结果如果在 leader || follower 中单独执行，会生成不一样的记录，导致状态机不一致
// 以启动流程实例为例：‘
// 1. leader 收到 cmd，发送 log
// 2. leader 收到 log，生成实例（包含 id、时间之类的），保存
// 3. leader 返回数据
// 4. follower 收到 log，生成实例（包含 id、时间之类的），保存
// 此处难以协调 2/4 步骤中生成的实例信息一致，有如下两种做法：
// 1. 在写 log 之前生成所有的信息，保证 log 是决定性的（herddb 做法）
//   1.1 这种需要考虑：首先生成信息，然后 log，最后应用的情况下，如果做到根据最新的数据来生成信息
//   1.2 例如：实例1的 job1 和 job2 都完成了，此时要生成 job3/job4，如何确保 job3/4 的 id 不重复
//   1.3 同时，也要处理 log 失败的情况
// 2. follower 不做任何生成相关的操作，只持久化已经发生的信息（zeebe 当前做法）
//   2.1 这种需要 leader 生成信息，问题在于：收到 log，持久化 log 信息，生成新的事件，发送到 raft 等需要保证事务特性，不然假设前3步成功（这个容易通过事务保证，第4步失败，则难以确保后续的操作能够继续
//   2.2 虽然我们能够通过可靠事件消息等方式来弥补上述缺点，但是无法解决如下问题
//     2.2.1 前3步都成功了，然后 crash（已经返回给客户端）
//     2.2.2 follower 提升为 leader
//     2.2.3 follower 应用了所有日志，发现 log1 没有后续操作，继续应用 log1，生成实例
//     2.2.4 但是此时生成的实例信息可能和第1步中是不同的（例如实例id），对于客户端来说，就会看到同一个请求，出现了两个实例id

func (s *StateMachine) Execute() error {
	ctx := context.Background()
	cancel := context.CancelFunc(func() {})
	if s.TimeoutSeconds != nil {
		ctx, cancel = context.WithTimeout(ctx, time.Duration(*s.TimeoutSeconds)*time.Second)
	}
	defer cancel()

	return s.executeState(ctx, s.StartAt)
}

func (s *StateMachine) executeState(ctx context.Context, n string) error {
	state, ok := s.States[n]
	if !ok {
		return fmt.Errorf("unknown state %v", n)
	}

	return state.Execute(ctx, s.executeState)
}

type InnerStateMachine struct {
	StartAt string `json:"StartAt"`
	States  States `json:"States"`
}

func (s *InnerStateMachine) Execute(ctx context.Context) error {
	return s.executeState(ctx, s.StartAt)
}

func (s *InnerStateMachine) executeState(ctx context.Context, n string) error {
	state, ok := s.States[n]
	if !ok {
		return fmt.Errorf("unknown state %v", n)
	}

	return state.Execute(ctx, s.executeState)
}

type State interface {
	Execute(context.Context, func(context.Context, string) error) error
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
