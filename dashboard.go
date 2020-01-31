package grabana

import (
	"github.com/K-Phoen/grabana/row"
	"github.com/grafana-tools/sdk"
)

type Dashboard struct {
	ID  uint   `json:"id"`
	UID string `json:"uid"`
	URL string `json:"url"`
}

// TagAnnotation describes an annotation represented as a Tag.
// See https://grafana.com/docs/grafana/latest/reference/annotations/#query-by-tag
type TagAnnotation struct {
	Name       string
	Datasource string
	IconColor  string
	Tags       []string
}

type DashboardBuilderOption func(dashboard *DashboardBuilder)

type DashboardBuilder struct {
	board *sdk.Board
}

func NewDashboardBuilder(title string, options ...DashboardBuilderOption) DashboardBuilder {
	board := sdk.NewBoard(title)
	board.ID = 0
	board.Timezone = ""

	builder := &DashboardBuilder{board: board}

	for _, opt := range append(dashboardDefaults(), options...) {
		opt(builder)
	}

	return *builder
}

func dashboardDefaults() []DashboardBuilderOption {
	return []DashboardBuilderOption{
		defaultTimePicker(),
		defaultTime(),
		WithSharedCrossHair(),
	}
}

func defaultTime() DashboardBuilderOption {
	return func(builder *DashboardBuilder) {
		builder.board.Time = sdk.Time{
			From: "now-3h",
			To:   "now",
		}
	}
}

func defaultTimePicker() DashboardBuilderOption {
	return func(builder *DashboardBuilder) {
		builder.board.Timepicker = sdk.Timepicker{
			RefreshIntervals: []string{"5s", "10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d"},
			TimeOptions:      []string{"5m", "15m", "1h", "6h", "12h", "24h", "2d", "7d", "30d"},
		}
	}
}

// WithRow adds a row to the dashboard.
func WithRow(title string, options ...row.Option) DashboardBuilderOption {
	return func(builder *DashboardBuilder) {
		row.New(builder.board, title, options...)
	}
}

// WithTagsAnnotation adds a new source of annotation for the dashboard.
func WithTagsAnnotation(annotation TagAnnotation) DashboardBuilderOption {
	return func(builder *DashboardBuilder) {
		builder.board.Annotations.List = append(builder.board.Annotations.List, sdk.Annotation{
			Name:       annotation.Name,
			Datasource: &annotation.Datasource,
			IconColor:  annotation.IconColor,
			Enable:     true,
			Tags:       annotation.Tags,
			Type:       "tags",
		})
	}
}

// Editable marks the graph as editable.
func Editable() DashboardBuilderOption {
	return func(builder *DashboardBuilder) {
		builder.board.Editable = true
	}
}

// ReadOnly marks the graph as non-editable.
func ReadOnly() DashboardBuilderOption {
	return func(builder *DashboardBuilder) {
		builder.board.Editable = false
	}
}

// WithSharedCrossHair configures the graph tooltip to be shared across panels.
func WithSharedCrossHair() DashboardBuilderOption {
	return func(builder *DashboardBuilder) {
		builder.board.SharedCrosshair = true
	}
}

// WithoutSharedCrossHair configures the graph tooltip NOT to be shared across panels.
func WithoutSharedCrossHair() DashboardBuilderOption {
	return func(builder *DashboardBuilder) {
		builder.board.SharedCrosshair = false
	}
}

// WithTags adds the given set of tags to the dashboard.
func WithTags(tags []string) DashboardBuilderOption {
	return func(builder *DashboardBuilder) {
		builder.board.Tags = tags
	}
}
