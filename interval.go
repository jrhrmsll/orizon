package orizon

import (
	"encoding/json"
	"fmt"
	"math"
	"slices"
	"time"
)

const (
	IntervalSpecKindCalendarMonth     = "cm:1"
	IntervalSpecKindOneCalendarWeek   = "cw:1"
	IntervalSpecKindTwoCalendarWeeks  = "cw:2"
	IntervalSpecKindFourCalendarWeeks = "cw:4"

	IntervalSpecDirectionBackward = "backward"
	IntervalSpecDirectionForward  = "forward"
)

var kinds = []string{
	IntervalSpecKindCalendarMonth,
	IntervalSpecKindOneCalendarWeek,
	IntervalSpecKindTwoCalendarWeeks,
	IntervalSpecKindFourCalendarWeeks,
}

var directions = []string{
	IntervalSpecDirectionBackward,
	IntervalSpecDirectionForward,
}

var (
	ErrInvalidIntervalKind      = fmt.Errorf("invalid interval kind")
	ErrInvalidIntervalDirection = fmt.Errorf("invalid interval direction")
)

type IntervalSpec struct {
	Start     float64 `json:"start"`
	End       float64 `json:"end"`
	Location  string  `json:"location"`
	Kind      string  `json:"kind"`
	Direction string  `json:"direction"`
	Limit     int     `json:"limit"`
}

func (spec *IntervalSpec) UnmarshalJSON(data []byte) error {
	type IntervalSpecAlias IntervalSpec

	if err := json.Unmarshal(data, (*IntervalSpecAlias)(spec)); err != nil {
		return err
	}

	if !slices.Contains(kinds, spec.Kind) {
		return ErrInvalidIntervalKind
	}

	if !slices.Contains(directions, spec.Direction) {
		return ErrInvalidIntervalDirection
	}

	return nil
}

type Interval struct {
	Start      time.Time `json:"start"`
	End        time.Time `json:"end"`
	Seconds    float64   `json:"seconds"`
	Days       float64   `json:"days"`
	Periods    float64   `json:"periods"`
	AtModifier int64     `json:"at_modifier"`
	Duration   string    `json:"duration"`
}

func NewInterval(start, end time.Time) Interval {
	duration := end.Sub(start).Round(time.Second)

	return Interval{
		Start:      start,
		End:        end.Truncate(time.Second),
		Seconds:    duration.Seconds(),
		Days:       math.Round(duration.Seconds() / (60 * 60 * 24)),
		Periods:    duration.Seconds() / (60 * 5),
		AtModifier: start.Add(duration).Unix(),
		Duration:   fmt.Sprintf("%0.fs", duration.Seconds()),
	}
}

type IntervalGenerator interface {
	Intervals() []Interval
}

type IntervalService interface {
	Find(spec *IntervalSpec) []Interval
}
