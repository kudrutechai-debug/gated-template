"""Pure business types. Imports nothing from adapter or handler, by contract."""

from gated_template.domain.models import Invoice

__all__ = ["Invoice"]
