"""One file per target. This one is the JSON payload shape."""

from __future__ import annotations

from gated_template.domain import Invoice


def invoice_to_payload(invoice: Invoice) -> dict[str, str]:
    """The wire shape is owned here, never by the domain type."""
    return {
        "reference": invoice.reference,
        "net": str(invoice.net),
        "vat": str(invoice.vat),
        "gross": str(invoice.gross),
    }
