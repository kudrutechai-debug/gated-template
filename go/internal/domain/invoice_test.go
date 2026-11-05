package domain

import "testing"

func TestVATAndGrossAreDerivedFromNet(t *testing.T) {
	inv := Invoice{Reference: "INV-001", NetCents: 10000, VATBasis: 810}

	if got, want := inv.VATCents(), int64(810); got != want {
		t.Errorf("VATCents() = %d, want %d", got, want)
	}
	if got, want := inv.GrossCents(), int64(10810); got != want {
		t.Errorf("GrossCents() = %d, want %d", got, want)
	}
}

func TestZeroRateLeavesNetUnchanged(t *testing.T) {
	inv := Invoice{Reference: "INV-002", NetCents: 5000, VATBasis: 0}

	if got, want := inv.GrossCents(), int64(5000); got != want {
		t.Errorf("GrossCents() = %d, want %d", got, want)
	}
}
