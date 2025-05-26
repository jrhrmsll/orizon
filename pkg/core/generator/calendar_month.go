package generator

import (
	"time"

	"github.com/jrhrmsll/orizon"
	"github.com/jrhrmsll/orizon/pkg/core/generator/internal"
)

// Ensure service implements interface.
var _ orizon.IntervalGenerator = (*CalendarMonth)(nil)

type CalendarMonth struct {
	state internal.GeneratorState
}

func NewCalendarMonth(spec *orizon.IntervalSpec) *CalendarMonth {
	return &CalendarMonth{
		state: internal.NewIteratorState(spec),
	}
}

func (iterator *CalendarMonth) Intervals() []orizon.Interval {
	intervals := make([]orizon.Interval, 0)

	ref := iterator.state.Ref

	for i := 0; i < iterator.state.Limit; i += 1 {
		start := time.Date(ref.Year(), ref.Month(), 1, 0, 0, 0, 0, ref.Location())
		end := start.AddDate(0, 1, 0).Add(-1 * time.Microsecond)

		switch iterator.state.Direction {
		case orizon.IntervalSpecDirectionBackward:
			ref = start.AddDate(0, 0, -1)
		case orizon.IntervalSpecDirectionForward:
			ref = end.AddDate(0, 0, 1)
		}

		intervals = append(intervals, orizon.NewInterval(start, end))
	}

	return intervals
}
