# TimeXtender Orchestration CLI — incomplete review snapshot

> **NOT READY FOR USE. Do not supply credentials, run triggers, or connect this MCP server to a tenant.**
> This repository is a public code-review snapshot, not a release or a completed Printing Press workflow.

Go CLI and MCP source generated with the real [CLI Printing Press](https://github.com/mvanhorn/cli-printing-press) 4.33.8 binary for TimeXtender Data Platform 26.2+ Public Orchestration API v0.

## Source and scope

- Official reference: https://support.timextender.com/general-137/timextender-orchestration-api-3946
- `spec/timextender.openapi.json` is **documentation-derived**, not an official tenant Swagger export. No tenant was available.
- The ten documented operations are under `/api/v0/orchestration`. Deprecated Jobs, Deliver and private APIs are excluded.
- `docs/endpoint-matrix.json` links each operation to emitted CLI source and an MCP tool. It is a source inventory, **not tested endpoint coverage**. All rows are explicitly unverified.

## Blocking issues

1. Generator validation failed on **GO-2026-6809**, a reachable Go standard-library vulnerability fixed in Go 1.26.10. The generator emits `go 1.26.9` and pins its scanner to that patch version; selecting Go 1.26.10 in the parent environment did not fix generation validation.
2. Generated triggers have **not** been hardened to preview by default with zero network calls and require explicit consent plus an explicit environment. **Do not invoke them.**
3. Generic mutation output treats all 2xx as success. This must be corrected so **202 means accepted, never execution completed**.
4. CLI/MCP endpoint, parameter and mutation-safety acceptance tests have not run. Do not infer safety from generated unit tests.
5. Approved composite commands `triage`, `failures`, `diagnose` are only scaffolds, not delivered features.
6. Mandatory tenant reachability and live acceptance remain open. No tenant requests or real starts were performed.

See [review status](docs/REVIEW-STATUS.md) for the exact validation outcome and phase status.

## Local review checks

These compile/test source only; do not run endpoint commands or supply a real API key.

```bash
GOTOOLCHAIN=go1.26.10 go test -count=1 ./...
GOTOOLCHAIN=go1.26.10 go vet ./...
GOTOOLCHAIN=go1.26.10 go build ./...
```

Using the patched toolchain for these checks does **not** retroactively pass the failed generator security gate. There are no release binaries, install instructions or deployment claims.

## Contract ambiguities awaiting tenant Swagger

- The common documentation says all lists support `limit`/`offset`, including lists whose individual endpoint sections omit pagination. The derived spec exposes them under that blanket rule; verify against tenant Swagger.
- The log endpoint has no documented `uat` parameter; none was invented.
- Execution timestamps are documented as local time without a specified timezone.
- Task identity includes both `taskType` and `taskId`; IDs are not globally unique.

## What is intentionally excluded

Credentials, tenant configuration, private responses, local runstate and receipt ledger, caches, binaries, preservation snapshots, raw documentation copies and auto-release workflows are not published here. This is not an official TimeXtender product and is not a public Printing Press library submission.

## License and provenance

Generated source retains its copyright headers, Apache-2.0 `LICENSE` and `NOTICE`. See `NOTICE` for Printing Press attribution. This review README and review-status documentation were added after generation to disclose incomplete and unsafe behavior.
