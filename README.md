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
