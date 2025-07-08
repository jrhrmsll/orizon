package iterator

import (
	"context"
	"time"

	"github.com/jrhrmsll/orizon"
	"github.com/jrhrmsll/orizon/pkg/core/iterator/internal"
)

const weekDays = 7

// Ensure service implements interface.
var _ internal.Iterator = (*CalendarWeek)(nil)

type CalendarWeek struct {
	state  *internal.IteratorState
	offset int
}

func NewCalendarWeek(spec *orizon.IntervalSpec, weeks int) *CalendarWeek {
	return &CalendarWeek{
		state:  internal.NewIteratorState(spec),
		offset: weekDays * (weeks - 1),
	}
}

func (iterator *CalendarWeek) Intervals(ctx context.Context) ([]orizon.Interval, error) {
	intervals := make([]orizon.Interval, 0)

	ref := iterator.state.Ref

	for i := 0; i < iterator.state.Limit; i += 1 {
		if err := context.Cause(ctx); err != nil {
			return nil, err
		}

		weekday := int(ref.Weekday())
		// To make Monday the start of the week, treat Sunday (0) as 7.
		if weekday == 0 {
			weekday = 7
		}

		weekStart := time.Date(ref.Year(), ref.Month(), ref.Day(), 0, 0, 0, 0, ref.Location()).AddDate(0, 0, -weekday+1)
		weekEnd := weekStart.Add(7 * 24 * time.Hour)

		start, end := time.Time{}, time.Time{}
		switch iterator.state.Direction {
		case orizon.IntervalSpecDirectionBackward:
			start = weekStart.AddDate(0, 0, -iterator.offset)
			end = weekEnd

			ref = start.AddDate(0, 0, -1)
		case orizon.IntervalSpecDirectionForward:
			start = weekStart
			end = weekEnd.AddDate(0, 0, iterator.offset)

			ref = end.AddDate(0, 0, 1)
		}

		end = end.Add(-1 * time.Microsecond)

		intervals = append(intervals, orizon.NewInterval(start, end))
	}

	return intervals, nil
}
