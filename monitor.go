package stats

import (
	"time"
)

// Monitor monitors given values at different timeframes
type Monitor struct {
	interval time.Duration
	metrics  []Metric
    publishers []Publisher
	stop     chan bool
}

// NewMonitor creates a new monitor
func NewMonitor(interval time.Duration) *Monitor {
	return &Monitor{
		interval: interval,
		stop:     make(chan bool),
	}
}

// Tracks tracks a metric
func (m *Monitor) Track(metric Metric) {
	switch metric.Type {
	case MovingAverage:
		if metric.Samples < 1 {
			metric.Samples = 1
		}

		metric.history = NewRingBuffer(metric.Samples)
	}

	m.metrics = append(m.metrics, metric)
}

// AddPublisher adds a target to publish metrics to
func (m *Monitor) AddPubisher(publisher Publisher) {
    m.publishers = append(m.publishers, publisher)
}

// Start starts the monitor
func (m *Monitor) Start() {
	for {
		select {
		case <-time.After(m.interval):
			for i := range m.metrics {
				v := m.metrics[i].compute()
				m.dispatch(m.metrics[i].Name, v)
			}
        case <-m.stop:
            return
		}
	}
}

// Stop stops the monitor
func (m *Monitor) Stop() {
	m.stop <- true
}

func (m *Monitor) dispatch(name string, metric int64) {
	for _, p := range m.publishers {
        p.Publish(name, metric)
    }
}
