package main

import (
	"strconv"

	"go.opentelemetry.io/otel/attribute"
)

func tpls2SpanAttrs(tpls []string) []attribute.KeyValue {
	attrs := []attribute.KeyValue{}
	for i, t := range tpls {
		name := "file" + strconv.Itoa(i)
		attrs = append(attrs, attribute.String(name, t))
	}

	return attrs
}
