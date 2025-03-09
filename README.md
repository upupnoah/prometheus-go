# Prometheus Go Examples

This repository contains a collection of examples demonstrating how to use the Prometheus client library for Go. These examples showcase various monitoring and instrumentation patterns for Go applications.

## Prerequisites

- Go (1.16 or newer recommended)
- Basic understanding of Prometheus concepts

## Installation

Clone this repository and navigate to any example directory to run it:

```bash
git clone https://github.com/yourusername/prometheus-go.git
cd prometheus-go/examples/simple
go run main.go
```

## Examples

### Simple (`examples/simple`)

A minimal example showing how to set up Prometheus instrumentation in a Go application, including:

- Creating a custom registry
- Registering Go runtime metrics and process collectors
- Exposing metrics via an HTTP endpoint

### Go Collector (`examples/gocollector`)

Demonstrates the use of the Go collector to automatically collect runtime metrics from your Go application.

### Version Collector (`examples/versioncollector`)

Shows how to set up a version collector to expose version information about your application.

### Runtime Metrics (`examples/runtimemetrics`)

Provides an example of collecting detailed runtime metrics from a Go application.

### Created Timestamps (`examples/createdtimestamps`)

Demonstrates how to add creation timestamps to metrics, which can be useful for tracking when metrics were first observed.

### Custom Labels (`examples/customlables`)

Shows how to add custom labels to all metrics exported by a collector, useful for adding context to your metrics.

### Exemplars (`examples/exemplars`)

Illustrates how to record exemplars with metrics, which are useful for correlating metrics with specific events or traces.

### Middleware (`examples/middleware`)

Contains examples of HTTP middleware for instrumenting HTTP handlers:

- Measuring request duration
- Counting requests
- Tracking in-flight requests

### Random (`examples/random`)

Demonstrates generating random metrics for testing or demonstration purposes.

## Usage

Each example can be run independently. Navigate to the example directory and run:

```bash
go run main.go
```

Then, access the metrics endpoint at http://localhost:8080/metrics (or the port specified in the example).

## Documentation

For more detailed information on the Prometheus client library for Go, visit:

- [Prometheus Go Client](https://github.com/prometheus/client_golang)
- [GoDoc Documentation](https://pkg.go.dev/github.com/prometheus/client_golang/prometheus)

## License

This project is licensed under the Apache License 2.0 - see the license headers in each file for details.
