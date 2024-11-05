package main

import "go.opentelemetry.io/otel/metric"

type measures struct {
	requests metric.Int64Counter
}

// createMeasures instantiate all measures required by this service.
func createMeasures(meter metric.Meter) *measures {
	r, err := meter.Int64Counter("server.requests",
		metric.WithDescription("Number of received requests"),
		metric.WithUnit("1"))
	if err != nil {
		panic(err)
	}

	return &measures{
		requests: r,
	}
}
