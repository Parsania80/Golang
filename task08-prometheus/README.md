Task 8: Prometheus Exporter (Observability)
Folder Name: task08-prometheus

Scenario:

You have a service running in production. How do you know if it’s slow? How do you know how many requests it has handled? You don’t read log files for this; you use metrics. Prometheus is the industry standard for scraping metrics from Go applications.

Goal:

Build a simple HTTP server (reusing concepts from Task 4) that exposes a /metrics endpoint. You will instrument the code to count how many requests it receives and track how long each request takes.

Technical Concepts to Learn:

Prometheus Go Client: github.com/prometheus/client_golang/prometheus.
Metric Types:
Counter: A number that only goes up (e.g., total_requests).
Histogram: A complex metric that tracks distribution (e.g., request_duration_seconds buckets: how many requests took <0.1s, <0.5s, etc.).
Middleware Pattern: Wrapping an HTTP handler to automatically record metrics without modifying the core logic.
Requirements:

Dependencies:
go get github.com/prometheus/client_golang/prometheus
go get github.com/prometheus/client_golang/prometheus/promhttp
Define Metrics:
Create a global Counter named myapp_processed_ops_total with a label for status (e.g., “200”, “500”).
Create a global Histogram named myapp_request_duration_seconds.
Register these metrics with prometheus.MustRegister() in your init() or main().
The Application Logic:
Create a simple random handler:
It sleeps for a random time between 0ms and 2000ms (to simulate work).
It randomly returns a 200 OK or a 500 Internal Server Error (to simulate failure).
Instrumentation (The “Middleware”):
Write a wrapper function or modify the handler to:
Start a timer (time.Now()).

Execute the request logic.

Calculate duration (time.Since(start)).

Update the Histogram (Observe(duration.Seconds())).

Increment the Counter (Inc() or WithLabelValues("200").Inc()).

Expose Metrics:

Create an endpoint /metrics.
Use the standard library handler: http.Handle("/metrics", promhttp.Handler()).
Success Criteria (for when you code it):

Run the app.
Hit the main endpoint (e.g., localhost:8080/work) multiple times in your browser or with curl.
Go to localhost:8080/metrics.
You should see plain text output looking like this:
text
    # HELP myapp_processed_ops_total The total number of processed events
    # TYPE myapp_processed_ops_total counter
    myapp_processed_ops_total{status="200"} 4
    myapp_processed_ops_total{status="500"} 1
Why this matters for DevOps:

In Kubernetes, every serious application has a /metrics endpoint. Prometheus (running in the cluster) periodically “scrapes” (visits) this URL to collect data. This data is then visualized on Grafana dashboards. You are building the data source for those dashboards.