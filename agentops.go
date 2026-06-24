// Package agentops is togo's agent operations/observability plugin: it records
// agent/LLM runs (tokens, estimated cost, latency, status) over the ai providers
// and exposes run history. Designed so usage flows into the billing plugin.
// Blank-import it; wrap calls with Service.Track or push Runs with Record.
package agentops

import (
	"context"
	"sync"
	"time"

	"github.com/togo-framework/ai"
	"github.com/togo-framework/togo"
)

// Run is one recorded agent/LLM invocation.
type Run struct {
	ID        string    `json:"id"`
	Agent     string    `json:"agent"`
	Model     string    `json:"model"`
	Input     int       `json:"input_tokens"`
	Output    int       `json:"output_tokens"`
	Total     int       `json:"total_tokens"`
	CostUSD   float64   `json:"cost_usd"`
	LatencyMs int64     `json:"latency_ms"`
	Status    string    `json:"status"`
	Error     string    `json:"error,omitempty"`
	StartedAt time.Time `json:"started_at"`
}

// Service stores recent runs (ring-buffered) and estimates cost.
type Service struct {
	mu   sync.Mutex
	runs []Run
	max  int
	seq  int
}

// UsageReporter, if set, receives each run (e.g. the billing plugin).
var UsageReporter func(Run)

func init() {
	togo.RegisterProviderFunc("ai-agentops", togo.PriorityService, func(k *togo.Kernel) error {
		k.Set("ai-agentops", &Service{max: 2000})
		return nil
	})
}

// FromKernel returns the agentops service bound to the kernel.
func FromKernel(k *togo.Kernel) (*Service, bool) {
	v, ok := k.Get("ai-agentops")
	if !ok {
		return nil, false
	}
	s, ok := v.(*Service)
	return s, ok
}

// Record stores a run (cost is filled from the model price table when zero).
func (s *Service) Record(r Run) {
	if r.CostUSD == 0 {
		r.CostUSD = EstimateCost(r.Model, r.Input, r.Output)
	}
	if r.Total == 0 {
		r.Total = r.Input + r.Output
	}
	s.mu.Lock()
	s.seq++
	if r.ID == "" {
		r.ID = time.Now().Format("20060102T150405") + "-" + itoa(s.seq)
	}
	s.runs = append(s.runs, r)
	if len(s.runs) > s.max {
		s.runs = s.runs[len(s.runs)-s.max:]
	}
	s.mu.Unlock()
	if UsageReporter != nil {
		UsageReporter(r)
	}
}

// List returns up to limit most-recent runs (newest first).
func (s *Service) List(limit int) []Run {
	s.mu.Lock()
	defer s.mu.Unlock()
	n := len(s.runs)
	if limit <= 0 || limit > n {
		limit = n
	}
	out := make([]Run, 0, limit)
	for i := n - 1; i >= 0 && len(out) < limit; i-- {
		out = append(out, s.runs[i])
	}
	return out
}

// Track times fn (an LLM/agent call returning its token usage) and records a Run.
func (s *Service) Track(_ context.Context, agent, model string, fn func() (ai.Usage, error)) error {
	start := time.Now()
	u, err := fn()
	r := Run{
		Agent: agent, Model: model,
		Input: u.PromptTokens, Output: u.CompletionTokens, Total: u.TotalTokens,
		LatencyMs: time.Since(start).Milliseconds(),
		StartedAt: start, Status: "ok",
	}
	if err != nil {
		r.Status = "error"
		r.Error = err.Error()
	}
	s.Record(r)
	return err
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var b [20]byte
	i := len(b)
	for n > 0 {
		i--
		b[i] = byte('0' + n%10)
		n /= 10
	}
	return string(b[i:])
}
