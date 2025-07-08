package core

import (
	"context"

	"github.com/jrhrmsll/orizon"
	"github.com/jrhrmsll/orizon/pkg/core/iterator"
)

// Ensure service implements interface.
var _ orizon.IntervalService = (*IntervalService)(nil)

type IntervalService struct{}

func NewIntervalService() *IntervalService {
	return &IntervalService{}
}

func (srv *IntervalService) Find(ctx context.Context, spec *orizon.IntervalSpec) ([]orizon.Interval, error) {
	return iterator.Factory(spec).Intervals(ctx)
}
