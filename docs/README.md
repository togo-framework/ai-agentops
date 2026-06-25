# ai-agentops — documentation

  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />

## Overview

Package agentops is togo's agent operations/observability plugin: it records
agent/LLM runs (tokens, estimated cost, latency, status) over the ai providers
and exposes run history. Designed so usage flows into the billing plugin.
Blank-import it; wrap calls with Service.Track or push Runs with Record.

## Install

```bash
togo install togo-framework/ai-agentops
```

A capability plugin — it self-registers on boot; no driver selector needed.

## Configuration

Environment variables read by this plugin (extracted from the source):

_No environment variables read directly (uses the kernel/base config or the app DB)._

## Usage

See the package API in the source.

## Links

- Marketplace: https://to-go.dev/marketplace
- Source: https://github.com/togo-framework/ai-agentops
- README: ../README.md
