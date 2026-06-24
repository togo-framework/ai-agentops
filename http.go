package agentops

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/togo-framework/togo"
)

// Handler exposes agentops over REST. Mount under /api/ai/agentops:
//
//	GET /runs?limit=50  -> [Run]
func Handler(k *togo.Kernel) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /runs", func(w http.ResponseWriter, r *http.Request) {
		svc, ok := FromKernel(k)
		if !ok {
			http.Error(w, "ai-agentops not configured", http.StatusInternalServerError)
			return
		}
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(svc.List(limit))
	})
	return mux
}
