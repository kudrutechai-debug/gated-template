# gated-template

A repository whose pull requests **cannot be merged until the gate is green**.

The gate is modelled on the one DIGIMIDI runs on `auarstadtt`, which checks, in order:
**vulnerabilities, lint, architecture, tests**. Two copies of the idea live here so the
mechanism is visible independently of language:

| | Go (`go/`) | Python (`python/`) |
|---|---|---|
| Vulnerabilities | `govulncheck` | `pip-audit` |
| Lint and format | `golangci-lint` | `ruff` |
| **Architecture** | `go-arch-lint` | `import-linter` |
| Tests | `go test` | `pytest` |

## The part worth copying

The architecture check is the unusual one. **Layering is enforced as a build failure, not
as a review comment.** Both halves declare the same rule:

```
handler  ->  adapter  ->  domain
```

`domain` may import nothing. Try adding `import "net/http"` to `go/internal/domain/invoice.go`
and the `go / arch` job fails, with no human needing to notice it in review.

## Run the gate locally

Always do this before pushing. It runs the same checks the gate runs, so a lint failure
costs you no round trip through CI.

```sh
make check
```

First time only, to fetch the Go tools:

```sh
make -C go tools
```

Python needs its dev dependencies once:

```sh
make -C python install
```

## Why the branch is protected

`main` carries a ruleset requiring all eight checks plus a pull request. Pushing straight
to `main` is refused, by you as well as by anyone else.
