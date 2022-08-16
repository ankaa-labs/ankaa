package engine

import (
	"errors"
	"strings"
)

// Catcher represent transitions when State's reports an error
// and there is no Retrier or retries have failed to resolve the error.
// When the error appears in the value of ErrroEquals field, transition
// the machine to the state named in the value of the 'Next' field.
type Catcher struct {
	// ErrorEquals is a non-empty array of Strings, which match Error Names.
	ErrorEquals []string `json:"ErrorEquals"`

	// Next is a string exactly matching a State Name.
	Next string `json:"Next"`

	// ResultPath must be a Reference Path, which specifies the raw input's combination with or replacement by the state's result.
	// Defaults to $
	ResultPath *string `json:"ResultPath"`
}

// Validate will validate the Catcher configuration
func (c *Catcher) Validate() error {
	if len(c.ErrorEquals) == 0 {
		return errors.New("ErrorEquals cannot be empty")
	}

	for _, errorEqual := range c.ErrorEquals {
		if len(errorEqual) == 0 {
			return errors.New("ErrorEqual cannot be empty")
		}

		if strings.HasPrefix(errorEqual, ErrorCodeStatesPrefix) {
			if ok := AllErrorCodes[errorEqual]; !ok {
				return errors.New("ErrorEqual is not pre defined")
			}

			if errorEqual == ErrorCodeStatesAll && len(c.ErrorEquals) != 1 {
				return errors.New("States.ALL must be be only element in ErrorEquals")
			}
		}
	}

	if len(c.Next) == 0 {
		return errors.New("Next cannot be empty")
	}

	return nil
}

// IsAnyErrorWildcard represent where is Catcher will matches any Error Names
func (c *Catcher) IsAnyErrorWildcard() bool {
	if len(c.ErrorEquals) != 1 {
		return false
	}

	return c.ErrorEquals[0] == ErrorCodeStatesAll
}

// Catchers must be an array of Catcher.
type Catchers []Catcher

// Validate will validate the Catchers configuration
func (c Catchers) Validate() error {
	for i, catcher := range c {
		if err := catcher.Validate(); err != nil {
			return err
		}

		if i != len(c)-1 && catcher.IsAnyErrorWildcard() {
			return errors.New("States.ALL must be the last catcher")
		}
	}

	return nil
}
