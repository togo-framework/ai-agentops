<!-- togo-header -->
<div align="center">
  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />
  <h1>togo-framework/ai-agentops</h1>
  <p>
    <a href="https://to-go.dev/marketplace"><img src="https://img.shields.io/badge/marketplace-to--go.dev-1FC7DC" alt="marketplace" /></a>
    <a href="https://pkg.go.dev/github.com/togo-framework/ai-agentops"><img src="https://pkg.go.dev/badge/github.com/togo-framework/ai-agentops.svg" alt="pkg.go.dev" /></a>
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT" />
  </p>
  <p><strong>Part of the <a href="https://to-go.dev">togo</a> framework.</strong></p>
</div>

## Install

```bash
togo install togo-framework/ai-agentops
```

<!-- /togo-header -->

# ai-agentops — agent operations & observability for togo

Records agent/LLM runs over the `ai` providers — **token usage, estimated cost,
latency, status** — and exposes run history. Token usage is designed to flow into
the `billing` plugin (set `agentops.UsageReporter`).

```bash
togo install togo-framework/ai-agentops
```

```go
ops, _ := agentops.FromKernel(k)
ops.Track(ctx, "support-bot", "gpt-4o-mini", func() (ai.Usage, error) {
    res, err := aiSvc.Chat(ctx, req)
    return res.Usage, err
})
```

`GET /api/ai/agentops/runs?limit=50` returns recent runs. Cost is estimated from a
built-in model price table (`EstimateCost`). MIT

<!-- togo-sponsors -->
---

<div align="center">
  <h3>Premium sponsors</h3>
  <p>
    <a href="https://id8media.com"><strong>ID8 Media</strong></a> &nbsp;·&nbsp;
    <a href="https://one-studio.co"><strong>One Studio</strong></a>
  </p>
  <p><sub>Support togo — <a href="https://github.com/sponsors/fadymondy">become a sponsor</a>.</sub></p>
</div>
<!-- /togo-sponsors -->
