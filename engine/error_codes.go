package engine

const (
	ErrorCodeStatesAll                    string = "States.ALL"
	ErrorCodeStatesHeartbeatTimeout       string = "States.HeartbeatTimeout"
	ErrorCodeStatesTimeout                string = "States.Timeout"
	ErrorCodeStatesTaskFailed             string = "States.TaskFailed"
	ErrorCodeStatesPermissions            string = "States.Permissions"
	ErrorCodeStatesResultPathMatchFailure string = "States.ResultPathMatchFailure"
	ErrorCodeStatesParameterPathFailure   string = "States.ParameterPathFailure"
	ErrorCodeStatesBranchFailed           string = "States.BranchFailed"
	ErrorCodeStatesNoChoiceMatched        string = "States.NoChoiceMatched"
	ErrorCodeStatesIntrinsicFailure       string = "States.IntrinsicFailure"
)

const (
	ErrorCodeStatesPrefix = "States."
)

var (
	AllErrorCodes = map[string]bool{
		ErrorCodeStatesAll:                    true,
		ErrorCodeStatesHeartbeatTimeout:       true,
		ErrorCodeStatesTimeout:                true,
		ErrorCodeStatesTaskFailed:             true,
		ErrorCodeStatesPermissions:            true,
		ErrorCodeStatesResultPathMatchFailure: true,
		ErrorCodeStatesParameterPathFailure:   true,
		ErrorCodeStatesBranchFailed:           true,
		ErrorCodeStatesNoChoiceMatched:        true,
		ErrorCodeStatesIntrinsicFailure:       true,
	}
)
