// Package adapter maps domain types to and from outside shapes.
// One file per target; this one is the JSON payload.
package adapter

import (
	"fmt"

	"github.com/kudrutechai-debug/gated-template/go/internal/domain"
)

// InvoicePayload owns the JSON names. The domain type carries no tags.
type InvoicePayload struct {
	Reference string `json:"reference"`
	Net       string `json:"net"`
	VAT       string `json:"vat"`
	Gross     string `json:"gross"`
}

// InvoiceToPayload renders an Invoice for the wire.
func InvoiceToPayload(i domain.Invoice) InvoicePayload {
	return InvoicePayload{
		Reference: i.Reference,
		Net:       money(i.NetCents),
		VAT:       money(i.VATCents()),
		Gross:     money(i.GrossCents()),
	}
}

func money(cents int64) string {
	return fmt.Sprintf("%d.%02d", cents/100, cents%100)
}
