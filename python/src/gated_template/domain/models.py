"""Domain types carry no serialisation concerns: no JSON names, no HTTP status."""

from __future__ import annotations

from dataclasses import dataclass
from decimal import Decimal


@dataclass(frozen=True)
class Invoice:
    """An invoice, in the language of the business rather than the wire."""

    reference: str
    net: Decimal
    vat_rate: Decimal

    @property
    def vat(self) -> Decimal:
        return self.net * self.vat_rate

    @property
    def gross(self) -> Decimal:
        return self.net + self.vat
