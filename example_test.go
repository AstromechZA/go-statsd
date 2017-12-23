package statsd_test

import (
	"github.com/AstromechZA/go-statsd"
)

func Example() {
	client, err := statsd.New(statsd.Address("127.0.0.1:8125"))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	client.Increment("thing.counter")

	client.Gauge("value.dial", 42.3)
}
