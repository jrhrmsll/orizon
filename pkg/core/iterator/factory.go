package iterator

import (
	"github.com/jrhrmsll/orizon"
	"github.com/jrhrmsll/orizon/pkg/core/iterator/internal"
)

func Factory(spec *orizon.IntervalSpec) internal.Iterator {
	switch spec.Kind {
	case orizon.IntervalSpecKindCalendarMonth:
		return NewCalendarMonth(spec)
	case orizon.IntervalSpecKindOneCalendarWeek:
		return NewCalendarWeek(spec, 1)
	case orizon.IntervalSpecKindTwoCalendarWeeks:
		return NewCalendarWeek(spec, 2)
	case orizon.IntervalSpecKindFourCalendarWeeks:
		return NewCalendarWeek(spec, 4)
	default:
		return nil
	}
}
