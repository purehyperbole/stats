package stats

import (
	"sync/atomic"
	"time"
)

// MetricType denotes the type of metric to track
type MetricType int

const (
	Counter MetricType = iota
	MovingAverage
)

// Metric
type Metric struct {
	Type     MetricType
	Name     string
	Interval time.Duration
	Samples  int
	Value    *int64
	history  *RingBuffer
}

func (m Metric) compute() int64 {
	v := atomic.LoadInt64(m.Value)

	switch m.Type {
	case Counter:
		return v
	case MovingAverage:
		m.history.Queue(v)

		var c int64

		m.history.Iterate(func(hv int64) {
			c = c + hv
		})

		// reset the value/counter
		atomic.AddInt64(m.Value, -v)

		return c / int64(m.Samples)
	}

	return v
}
