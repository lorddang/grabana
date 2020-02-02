package alert

import (
	"github.com/grafana-tools/sdk"
)

type ErrorMode string
type NoDataMode string
type Option func(alert *Alert)

const Alerting ErrorMode = "alerting"
const LastState ErrorMode = "keep_state"

const NoData NoDataMode = "no_data"
const Error NoDataMode = "alerting"
const KeepLastState NoDataMode = "keep_state"
const OK NoDataMode = "ok"

type Channel struct {
	ID   uint   `json:"id"`
	UID  string `json:"uid"`
	Name string `json:"Name"`
	Type string `json:"type"`
}

type Alert struct {
	Builder *sdk.Alert
}

func New(name string, options ...Option) *Alert {
	alert := &Alert{Builder: &sdk.Alert{
		Name:                name,
		Handler:             1, // TODO: what's that?
		ExecutionErrorState: string(LastState),
		NoDataState:         string(KeepLastState),
	}}

	for _, opt := range options {
		opt(alert)
	}

	return alert
}

// Notification adds a notification for this alert.
func Notification(channel *Channel) Option {
	return func(alert *Alert) {
		alert.Builder.Notifications = append(alert.Builder.Notifications, sdk.AlertNotification{
			ID:  int64(channel.ID),
			UID: channel.UID,
		})
	}
}

// Message sets the message associated to the alert.
func Message(content string) Option {
	return func(alert *Alert) {
		alert.Builder.Message = content
	}
}

// For sets the time interval during which a query violating the threshold
// before the alert being actually triggered.
// See https://grafana.com/docs/grafana/latest/alerting/rules/#for
func For(duration string) Option {
	return func(alert *Alert) {
		alert.Builder.For = duration
	}
}

// EvaluateEvery defines the evaluation interval.
func EvaluateEvery(interval string) Option {
	return func(alert *Alert) {
		alert.Builder.Frequency = interval
	}
}

// OnExecutionError defines the behavior on execution error.
// See https://grafana.com/docs/grafana/latest/alerting/rules/#execution-errors-or-timeouts
func OnExecutionError(mode ErrorMode) Option {
	return func(alert *Alert) {
		alert.Builder.ExecutionErrorState = string(mode)
	}
}

// OnNoData defines the behavior when the query returns no data.
// See https://grafana.com/docs/grafana/latest/alerting/rules/#no-data-null-values
func OnNoData(mode NoDataMode) Option {
	return func(alert *Alert) {
		alert.Builder.NoDataState = string(mode)
	}
}

// If adds a condition that could trigger the alert.
// See https://grafana.com/docs/grafana/latest/alerting/rules/#conditions
func If(operator Operator, opts ...ConditionOption) Option {
	return func(alert *Alert) {
		cond := newCondition(opts...)
		cond.builder.Operator = conditionOperator{Type: string(operator)}

		alert.Builder.Conditions = append(alert.Builder.Conditions, *cond.builder)
	}
}
