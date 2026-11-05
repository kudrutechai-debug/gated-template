from decimal import Decimal

from gated_template.domain import Invoice
from gated_template.handler import render_invoice


def test_vat_and_gross_are_derived_from_net() -> None:
    invoice = Invoice(reference="INV-001", net=Decimal("100.00"), vat_rate=Decimal("0.081"))
    assert invoice.vat == Decimal("8.100")
    assert invoice.gross == Decimal("108.100")


def test_handler_returns_the_wire_shape() -> None:
    invoice = Invoice(reference="INV-002", net=Decimal("50.00"), vat_rate=Decimal("0.00"))
    assert render_invoice(invoice) == {
        "reference": "INV-002",
        "net": "50.00",
        "vat": "0.0000",
        "gross": "50.0000",
    }
