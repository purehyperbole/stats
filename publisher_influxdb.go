package stats

import (
	"crypto/tls"
	"fmt"

	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/influxdata/influxdb-client-go/api"
)

// InfluxDBPublisher ships metrics to influxdb
type InfluxDBPublisher struct {
	client   influxdb2.Client
	writeAPI api.WriteApi
}

// NewInfluxDBPublisher creates a new influx db publisher
func NewInfluxDBPublisher(host, token, org, bucket string) *InfluxDBPublisher {
	opts := influxdb2.DefaultOptions()
	opts = opts.SetUseGZip(true)
	opts = opts.SetTlsConfig(&tls.Config{
		InsecureSkipVerify: true,
	})

	client := influxdb2.NewClientWithOptions(host, token, opts)
	writeAPI := client.WriteApi(org, bucket)

	return &InfluxDBPublisher{
		client:   client,
		writeAPI: writeAPI,
	}
}

// Publish publishes a given metric to influxdb
func (p *InfluxDBPublisher) Publish(name string, metric int64) {
	line := fmt.Sprintf("%s avg=%d", name, metric)
	p.writeAPI.WriteRecord(line)
}
