"""The only layer that knows about transport. Never imports a repository."""

from gated_template.handler.invoice import render_invoice

__all__ = ["render_invoice"]
