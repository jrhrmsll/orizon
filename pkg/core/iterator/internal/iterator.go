package internal

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/jrhrmsll/orizon"
)

type Iterator interface {
	Intervals(ctx context.Context) ([]orizon.Interval, error)
}

type IteratorState struct {
	Ref       time.Time
	Direction string
	Limit     int
}

func NewIteratorState(spec *orizon.IntervalSpec) *IteratorState {
	if strings.ToLower(spec.Location) == "utc" {
		spec.Location = "UTC"
	}

	location, err := time.LoadLocation(spec.Location)
	if err != nil {
		log.Println(err)
		location, _ = time.LoadLocation("UTC")
	}

	start := time.Unix(int64(spec.Start), 0).In(location)
	end := time.Unix(int64(spec.End), 0).In(location)

	ref := time.Now().In(location)
	switch spec.Direction {
	case orizon.IntervalSpecDirectionBackward:
		ref = end
	case orizon.IntervalSpecDirectionForward:
		ref = start
	}

	return &IteratorState{
		Ref:       ref,
		Direction: spec.Direction,
		Limit:     spec.Limit,
	}
}
