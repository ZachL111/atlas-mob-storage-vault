# atlas-mob-storage-vault

`atlas-mob-storage-vault` is a compact Go repository for mobile workflows, centered on this goal: Create a Go reference implementation for storage workflows, centered on constraint solving, bounded scenario files, and conflict explanations.

## Why It Exists

This is intentionally local and self-contained so it can be inspected without credentials, services, or seeded history.

## Atlas Mob Storage Vault Review Notes

The first comparison I would make is `local state` against `conflict cost` because it shows where the rule is most opinionated.

## Features

- `fixtures/domain_review.csv` adds cases for form pressure and sync drift.
- `metadata/domain-review.json` records the same cases in structured form.
- `config/review-profile.json` captures the read order and the two review questions.
- `examples/atlas-mob-storage-walkthrough.md` walks through the case spread.
- The Go code includes a review path for `local state` and `conflict cost`.
- `docs/field-notes.md` explains the strongest and weakest cases.

## Architecture Notes

The implementation keeps the scoring rule plain: reward signal and confidence, preserve slack, penalize drag, then classify the result into a review lane.

The Go implementation avoids hidden state so fixture changes are easy to reason about.

## Usage

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/verify.ps1
```

## Tests

The check exercises the source code and the review fixture. `edge` is the high score at 204; `recovery` is the low score at 150.

## Limitations And Roadmap

This remains a local project with deterministic fixtures. It does not depend on credentials, hosted services, or live data. Future work should add richer malformed inputs before widening the public API.
