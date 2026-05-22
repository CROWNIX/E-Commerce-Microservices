# Observability Stack (OpenTelemetry + Jaeger)

## Start stack (Podman)

From repository root:

```bash
make observability-up
```

Or:

```bash
cd deploy/observability
podman compose up -d
```

Jaeger UI: http://localhost:16686

OTLP endpoint for Go services (host): `localhost:4317`

## Stop stack

```bash
make observability-down
```

## Service configuration

Add these keys to each service `env.json` (auth, product, cart, order):

```json
{
  "OTEL_ENABLED": true,
  "OTEL_EXPORTER_OTLP_ENDPOINT": "localhost:4317",
  "OTEL_EXPORTER_OTLP_USERNAME": "",
  "OTEL_EXPORTER_OTLP_PASSWORD": ""
}
```

Set `OTEL_ENABLED` to `false` to run without the observability stack.

## Verify distributed trace (cart → product gRPC)

1. Start observability stack (`make observability-up`).
2. Start **product-service** and **cart-service** (`make run-api` in each).
3. Send a request that hits cart and calls product via gRPC, e.g. `POST /carts` through api-gateway (with valid JWT) or call cart REST API directly.
4. Open Jaeger UI → Service: `cart-service` (or your `APP_NAME`) → Find Traces.
5. Open a trace and confirm spans include:
   - HTTP handler on cart-service
   - gRPC client span (cart → product)
   - gRPC server span on product-service

## Architecture

```
Go services → OTLP :4317 → otel-collector
                              ├─ traces  → Jaeger (OTLP) → Jaeger UI :16686
                              └─ metrics → debug exporter (collector logs)
```

- **Traces** are forwarded to Jaeger for distributed tracing in the UI.
- **Metrics** are handled by a separate collector pipeline (Jaeger does not accept OTLP metrics). In development they are exported via the `debug` exporter and appear in collector container logs (`podman logs ecommerce-otel-collector`).

api-gateway is not instrumented; traces start at the first microservice that handles the request.
