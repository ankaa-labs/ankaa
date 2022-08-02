package engine

import (
	"time"
)

type Choice struct {
	// required
	Next string

	ChoiceRule
}

type ChoiceRule struct {
	// >= 1
	And []ChoiceRule
	// >= 1
	Or  []ChoiceRule
	Not *ChoiceRule

	StringEquals     *string
	StringEqualsPath *string

	StringLessThan     *string
	StringLessThanPath *string

	StringGreaterThan     *string
	StringGreaterThanPath *string

	StringLessThanEquals     *string
	StringLessThanEqualsPath *string

	StringGreaterThanEquals     *string
	StringGreaterThanEqualsPath *string

	StringMatches *string

	NumericEquals     *float64
	NumericEqualsPath *string

	NumericLessThan     *float64
	NumericLessThanPath *string

	NumericGreaterThan     *float64
	NumericGreaterThanPath *string

	NumericLessThanEquals     *float64
	NumericLessThanEqualsPath *string

	NumericGreaterThanEquals     *float64
	NumericGreaterThanEqualsPath *string

	BooleanEquals     *bool
	BooleanEqualsPath *string

	TimestampEquals     *time.Time
	TimestampEqualsPath *string

	TimestampLessThan     *time.Time
	TimestampLessThanPath *string

	TimestampGreaterThan     *time.Time
	TimestampGreaterThanPath *string

	TimestampLessThanEquals     *time.Time
	TimestampLessThanEqualsPath *string

	TimestampGreaterThanEquals     *time.Time
	TimestampGreaterThanEqualsPath *string

	IsNull *bool

	IsPresent *bool

	IsNumeric *bool

	IsString *bool

	IsBoolean *bool

	IsTimestamp *bool

	Variable *string
}

type ChoiceState struct {
	Type    string // Choice
	Comment *string

	InputPath  *string
	OutputPath *string

	// required
	// >= 1
	Choices []Choice

	Default *string
}
