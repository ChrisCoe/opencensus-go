package main

import (
	"go.opencensus.io/exporter/azure_monitor"
)

func main() {
	exporter := azure_monitor.NewExporter()
	exporter = exporter
}