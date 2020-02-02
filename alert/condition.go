package alert

import (
	"github.com/grafana-tools/sdk"
)

type Operator string
type ConditionOption func(condition *condition)

const And Operator = "and"
const Or Operator = "or"

type evaluator struct {
	Params []float64 `json:"params,omitempty"`
	Type   string    `json:"type,omitempty"`
}

type conditionOperator struct {
	Type string `json:"type,omitempty"`
}

type query struct {
	Params []string `json:"params,omitempty"`
}

type reducer struct {
	Params []string `json:"params,omitempty"`
	Type   string   `json:"type,omitempty"`
}

type condition struct {
	builder *sdk.AlertCondition
}

func newCondition(options ...ConditionOption) *condition {
	cond := &condition{builder: &sdk.AlertCondition{
		Type: "query",
	}}

	for _, opt := range options {
		opt(cond)
	}

	return cond
}

// Avg defines the query to execute.
// See https://grafana.com/docs/grafana/latest/alerting/rules/#query-condition-example
func Avg(refID string, from string, to string) ConditionOption {
	return func(cond *condition) {
		cond.builder.Query = query{Params: []string{refID, from, to}}
		cond.builder.Reducer = reducer{Type: "avg", Params: []string{}}
	}
}

// Avg defines the query to execute.
// See https://grafana.com/docs/grafana/latest/alerting/rules/#query-condition-example
func Sum(refID string, from string, to string) ConditionOption {
	return func(cond *condition) {
		cond.builder.Query = query{Params: []string{refID, from, to}}
		cond.builder.Reducer = reducer{Type: "sum", Params: []string{}}
	}
}

// Avg defines the query to execute.
// See https://grafana.com/docs/grafana/latest/alerting/rules/#query-condition-example
func Count(refID string, from string, to string) ConditionOption {
	return func(cond *condition) {
		cond.builder.Query = query{Params: []string{refID, from, to}}
		cond.builder.Reducer = reducer{Type: "count", Params: []string{}}
	}
}

// Avg defines the query to execute.
// See https://grafana.com/docs/grafana/latest/alerting/rules/#query-condition-example
func Last(refID string, from string, to string) ConditionOption {
	return func(cond *condition) {
		cond.builder.Query = query{Params: []string{refID, from, to}}
		cond.builder.Reducer = reducer{Type: "last", Params: []string{}}
	}
}

// Min defines the query to execute.
// See https://grafana.com/docs/grafana/latest/alerting/rules/#query-condition-example
func Min(refID string, from string, to string) ConditionOption {
	return func(cond *condition) {
		cond.builder.Query = query{Params: []string{refID, from, to}}
		cond.builder.Reducer = reducer{Type: "min", Params: []string{}}
	}
}

// Max defines the query to execute.
// See https://grafana.com/docs/grafana/latest/alerting/rules/#query-condition-example
func Max(refID string, from string, to string) ConditionOption {
	return func(cond *condition) {
		cond.builder.Query = query{Params: []string{refID, from, to}}
		cond.builder.Reducer = reducer{Type: "max", Params: []string{}}
	}
}

// Median defines the query to execute.
// See https://grafana.com/docs/grafana/latest/alerting/rules/#query-condition-example
func Median(refID string, from string, to string) ConditionOption {
	return func(cond *condition) {
		cond.builder.Query = query{Params: []string{refID, from, to}}
		cond.builder.Reducer = reducer{Type: "median", Params: []string{}}
	}
}

// Diff defines the query to execute.
// See https://grafana.com/docs/grafana/latest/alerting/rules/#query-condition-example
func Diff(refID string, from string, to string) ConditionOption {
	return func(cond *condition) {
		cond.builder.Query = query{Params: []string{refID, from, to}}
		cond.builder.Reducer = reducer{Type: "diff", Params: []string{}}
	}
}

// PercentDiff defines the query to execute.
// See https://grafana.com/docs/grafana/latest/alerting/rules/#query-condition-example
func PercentDiff(refID string, from string, to string) ConditionOption {
	return func(cond *condition) {
		cond.builder.Query = query{Params: []string{refID, from, to}}
		cond.builder.Reducer = reducer{Type: "percent_diff", Params: []string{}}
	}
}

func HasNoValue() ConditionOption {
	return func(cond *condition) {
		cond.builder.Evaluator = evaluator{Type: "no_value", Params: []float64{}}
	}
}

func IsAbove(value float64) ConditionOption {
	return func(cond *condition) {
		cond.builder.Evaluator = evaluator{Type: "gt", Params: []float64{value}}
	}
}

func IsBelow(value float64) ConditionOption {
	return func(cond *condition) {
		cond.builder.Evaluator = evaluator{Type: "lt", Params: []float64{value}}
	}
}

func IsOutsideRange(min float64, max float64) ConditionOption {
	return func(cond *condition) {
		cond.builder.Evaluator = evaluator{Type: "outside_range", Params: []float64{min, max}}
	}
}

func IsWithinRange(min float64, max float64) ConditionOption {
	return func(cond *condition) {
		cond.builder.Evaluator = evaluator{Type: "within_range", Params: []float64{min, max}}
	}
}
