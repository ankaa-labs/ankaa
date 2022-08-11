package engine

import (
	"time"
)

type Choice struct {
	// required
	Next string `json:"Next"`

	ChoiceRule
}

type ChoiceRule struct {
	// >= 1
	And []ChoiceRule `json:"And"`
	// >= 1
	Or  []ChoiceRule `json:"Or"`
	Not *ChoiceRule  `json:"Not"`

	StringEquals     *string `json:"StringEquals"`
	StringEqualsPath *string `json:"StringEqualsPath"`

	StringLessThan     *string `json:"StringLessThan"`
	StringLessThanPath *string `json:"StringLessThanPath"`

	StringGreaterThan     *string `json:"StringGreaterThan"`
	StringGreaterThanPath *string `json:"StringGreaterThanPath"`

	StringLessThanEquals     *string `json:"StringLessThanEquals"`
	StringLessThanEqualsPath *string `json:"StringLessThanEqualsPath"`

	StringGreaterThanEquals     *string `json:"StringGreaterThanEquals"`
	StringGreaterThanEqualsPath *string `json:"StringGreaterThanEqualsPath"`

	StringMatches *string `json:"StringMatches"`

	NumericEquals     *float64 `json:"NumericEquals"`
	NumericEqualsPath *string  `json:"NumericEqualsPath"`

	NumericLessThan     *float64 `json:"NumericLessThan"`
	NumericLessThanPath *string  `json:"NumericLessThanPath"`

	NumericGreaterThan     *float64 `json:"NumericGreaterThan"`
	NumericGreaterThanPath *string  `json:"NumericGreaterThanPath"`

	NumericLessThanEquals     *float64 `json:"NumericLessThanEquals"`
	NumericLessThanEqualsPath *string  `json:"NumericLessThanEqualsPath"`

	NumericGreaterThanEquals     *float64 `json:"NumericGreaterThanEquals"`
	NumericGreaterThanEqualsPath *string  `json:"NumericGreaterThanEqualsPath"`

	BooleanEquals     *bool   `json:"BooleanEquals"`
	BooleanEqualsPath *string `json:"BooleanEqualsPath"`

	TimestampEquals     *time.Time `json:"TimestampEquals"`
	TimestampEqualsPath *string    `json:"TimestampEqualsPath"`

	TimestampLessThan     *time.Time `json:"TimestampLessThan"`
	TimestampLessThanPath *string    `json:"TimestampLessThanPath"`

	TimestampGreaterThan     *time.Time `json:"TimestampGreaterThan"`
	TimestampGreaterThanPath *string    `json:"TimestampGreaterThanPath"`

	TimestampLessThanEquals     *time.Time `json:"TimestampLessThanEquals"`
	TimestampLessThanEqualsPath *string    `json:"TimestampLessThanEqualsPath"`

	TimestampGreaterThanEquals     *time.Time `json:"TimestampGreaterThanEquals"`
	TimestampGreaterThanEqualsPath *string    `json:"TimestampGreaterThanEqualsPath"`

	IsNull *bool `json:"IsNull"`

	IsPresent *bool `json:"IsPresent"`

	IsNumeric *bool `json:"IsNumeric"`

	IsString *bool `json:"IsString"`

	IsBoolean *bool `json:"IsBoolean"`

	IsTimestamp *bool `json:"IsTimestamp"`

	Variable *string `json:"Variable"`
}

type ChoiceState struct {
	Type    string  `json:"Type"` // Choice
	Comment *string `json:"Comment"`

	InputPath  *string `json:"InputPath"`
	OutputPath *string `json:"OutputPath"`

	// required
	// >= 1
	Choices []Choice `json:"Choices"`

	Default *string `json:"Default"`
}
