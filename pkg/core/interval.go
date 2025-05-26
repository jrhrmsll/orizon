package core

import (
	"github.com/jrhrmsll/orizon"
	"github.com/jrhrmsll/orizon/pkg/core/generator"
)

// Ensure service implements interface.
var _ orizon.IntervalService = (*IntervalService)(nil)

type IntervalService struct{}

func NewIntervalService() *IntervalService {
	return &IntervalService{}
}

func (srv *IntervalService) Find(spec *orizon.IntervalSpec) []orizon.Interval {
	return generator.Factory(spec).Intervals()
}
