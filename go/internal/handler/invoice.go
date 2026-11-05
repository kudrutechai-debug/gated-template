// Package handler is the only layer that knows about transport.
// It never imports a repository.
package handler

import (
	"encoding/json"
	"net/http"

	"github.com/kudrutechai-debug/gated-template/go/internal/adapter"
	"github.com/kudrutechai-debug/gated-template/go/internal/domain"
)

// Invoice writes one hard-coded invoice as JSON. It stands in for a real route.
func Invoice(w http.ResponseWriter, _ *http.Request) {
	inv := domain.Invoice{Reference: "INV-001", NetCents: 10000, VATBasis: 810}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(adapter.InvoiceToPayload(inv)); err != nil {
		http.Error(w, "could not encode invoice", http.StatusInternalServerError)
	}
}

// Health is liveness. It touches nothing downstream, on purpose.
func Health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
