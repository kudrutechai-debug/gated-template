.PHONY: install check audit lint arch test

PY ?= python3

install:
	$(PY) -m pip install -e '.[dev]'

## The same four checks the GitHub gate runs, in the same order.
check: audit lint arch test

audit:
	$(PY) -m pip_audit --skip-editable --progress-spinner off

lint:
	$(PY) -m ruff check .
	$(PY) -m ruff format --check .

arch:
	lint-imports

test:
	$(PY) -m pytest -q
