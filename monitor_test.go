package stats

import (
    "fmt"
	"sync/atomic"
	"testing"
	"time"
)

type testPublisher struct {
    values []string
}

func (p *testPublisher) Publish(name string, metric int64) {
    p.values = append(p.values, fmt.Sprintf("%s=%d", name, metric))
}

func TestMonitorCounter(t *testing.T) {
	m := NewMonitor(time.Second)

    p := testPublisher{}

    m.AddPubisher(&p)

	var counter int64

	m.Track(Metric{
		Type:  Counter,
		Name:  "basic counter",
		Value: &counter,
	})

	go m.Start()

	for i := 0; i < 4; i++ {
        go func() {
            for x := 0; x < 10; x++ {
                for y := 0; y < 1000; y++ {
                    atomic.AddInt64(&counter, 1)
                    time.Sleep(time.Millisecond)
                }
            }
        }()
	}

    for len(p.values) < 3 {
        fmt.Println(p.values)
        time.Sleep(time.Second)
    }

    m.Stop()

    fmt.Println(len(p.values))
}

func TestMonitorMovingAverage(t *testing.T) {
	m := NewMonitor(time.Second)

    p := testPublisher{}

    m.AddPubisher(p)

	var counter int64

	m.Track(Metric{
		Type:  MovingAverage,
		Name:  "moving average counter",
		Value: &counter,
	})

	go m.Start()

	for i := 0; i < 4; i++ {
        go func() {
            for x := 0; x < 10; x++ {
                for y := 0; y < 1000; y++ {
                    atomic.AddInt64(&counter, 1)
                    time.Sleep(time.Millisecond)
                }
            }
        }()
	}

    for len(p.values) < 3 {
        fmt.Println(p.values)
        time.Sleep(time.Second)
    }

    m.Stop()

    fmt.Println(len(p.values))
}
