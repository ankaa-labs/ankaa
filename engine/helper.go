package engine

import (
	"errors"
	"strings"
)

// IsAnyErrorWildcard represent where is ErrorEquals will matches any Error Names
func IsAnyErrorWildcard(errorEquals []string) bool {
	if len(errorEquals) != 1 {
		return false
	}

	return errorEquals[0] == ErrorCodeStatesAll
}

// ValidateErrorEquals validate ErrorEquals configuration
func ValidateErrorEquals(errorEquals []string) error {
	if len(errorEquals) == 0 {
		return errors.New("ErrorEquals cannot be empty")
	}

	for _, errorEqual := range errorEquals {
		if len(errorEqual) == 0 {
			return errors.New("ErrorEqual cannot be empty")
		}

		if strings.HasPrefix(errorEqual, ErrorCodeStatesPrefix) {
			if ok := AllErrorCodes[errorEqual]; !ok {
				return errors.New("ErrorEqual is not pre defined")
			}

			if errorEqual == ErrorCodeStatesAll && len(errorEquals) != 1 {
				return errors.New("States.ALL must be be only element in ErrorEquals")
			}
		}
	}

	return nil
}

// ValidateTerminalState validate Next & End configuration
func ValidateTerminalState(next *string, end bool) error {
	if next == nil && !end {
		return errors.New("Next is nil and End is false")
	}

	if next != nil {
		if *next == "" {
			return errors.New("Next is configured but is empty")
		}

		if end {
			return errors.New("Next is configured and End is true")
		}
	}

	return nil
}

// ValidatePath validate the path string is a valid JSONPath(https://github.com/json-path/JsonPath)
// TODO
func ValidatePath(path *string) error {
	return nil
}

// ValidateReferencePath validate the reference path string is a valid JSONPath,
// But it can only identify a single node in a JSON structure: The operators "@", ",", ":", and "?" are not supported,
// And Reference Paths MUST be unambiguous references to a single value, array, or object (subtree).
// TODO
func ValidateReferencePath(path *string) error {
	return nil
}
