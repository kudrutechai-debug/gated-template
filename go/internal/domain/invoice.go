// Package domain holds business types. By contract it imports nothing from
// adapter or handler, and no third-party package. go-arch-lint enforces that.
package domain

// Invoice is expressed in the language of the business, not of the wire.
// Amounts are in cents to keep the arithmetic exact.
type Invoice struct {
	Reference string
	NetCents  int64
	VATBasis  int64 // VAT rate in basis points, so 810 is 8.1 percent.
}

// VATCents is the tax due on the net amount.
func (i Invoice) VATCents() int64 {
	return i.NetCents * i.VATBasis / 10000
}

// GrossCents is what the customer actually pays.
func (i Invoice) GrossCents() int64 {
	return i.NetCents + i.VATCents()
}
