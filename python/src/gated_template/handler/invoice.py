"""Turns a domain object into something a caller can read."""

from __future__ import annotations

from gated_template.adapter import invoice_to_payload
from gated_template.domain import Invoice


def render_invoice(invoice: Invoice) -> dict[str, str]:
    return invoice_to_payload(invoice)
