package agentops

// price is USD per 1K tokens {input, output}.
type price struct{ in, out float64 }

var prices = map[string]price{
	"gpt-4o":          {0.0025, 0.01},
	"gpt-4o-mini":     {0.00015, 0.0006},
	"claude-3-5-sonnet": {0.003, 0.015},
	"claude-3-5-haiku":  {0.0008, 0.004},
	"gemini-1.5-pro":  {0.00125, 0.005},
	"gemini-1.5-flash": {0.000075, 0.0003},
	"deepseek-chat":   {0.00027, 0.0011},
}

// EstimateCost returns the estimated USD cost for a run; 0 when the model is unknown.
func EstimateCost(model string, in, out int) float64 {
	p, ok := prices[model]
	if !ok {
		return 0
	}
	return float64(in)/1000*p.in + float64(out)/1000*p.out
}
