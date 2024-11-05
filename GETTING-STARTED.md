Now that you have cloned the repo, what's next?

1. Make sure everything is working. Run `make start` and when everything is ready visit [localhost](http://localhost). You should see a site, "Cinema Website".
2. Click around the UI to generate some observability signals.
3. Inspect the various local observability tools available, you should see traces, metrics and logs! Follow [docs/observability](docs/observability.md).
4. Click around and inspect what is collected and how it is displayed. When you feel ready move on to the tasks below.

## Tasklist

- Not all services are instrumented! Goal: instrument a service by collecting traces from the HTTP server.
- Server traces alone are not enough, add specific traces to inspect communication with the database.
- Why does the `website: /users/view/{id}` trace for `website` service reports the `HTTP GET` span twice? What's going on?
- We do not collect any metrics from our DB (MongoDB). Make your inner DBA happy, add the collection of metrics and logs.
- Automatic instrumentation with eBPF is available, why not try it? https://github.com/open-telemetry/opentelemetry-go-instrumentation
