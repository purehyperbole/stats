# stats
Track and public custom statistics and metrics inside of your go application

Supports tracking:
- int64 counters
- int64 counters (moving average)

Supports publishing to:
- [x] influxdb

# Usage

```go
import (
    "sync/atomic"
    "time"

    "github.com/purehyperbole/stats"
)

func main() {
    var counter, movingAverage int64

    // create a monitor that will aggregate and publish metrics every second
    monitor := stats.NewMonitor(time.Second)

    // add a new destination to publish metrics to
    monitor.AddPublisher(
        stats.NewInfluxDBPublisher("host", "token", "org", "bucket")
    )

    // track a basic counter
    monitor.Track(stats.Metric{
        Type: stats.Counter,
        Name: "simple-counter,environment=production value=",
        Value: &counter,
    })

    // track a moving average. Any polled value will be averaged over a minute
    // (interval of 1 second, with 60 samples stored)
    monitor.Track(stats.Metric{
        Type:    stats.MovingAverage,
        Name:    "moving-average,environment=production avg-reqs=",
        Samples: 60,
        Value:   &movingAverages,
    })

    go monitor.Start()

    // use the values in your app
    for i := 0; i < 1000; i++ {
        atomic.AddInt64(&counter, 1)
        atomic.AddInt64(&movingAverage, 1)
    }
}
```
