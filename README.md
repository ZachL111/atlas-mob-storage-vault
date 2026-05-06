# atlas-mob-storage-vault

`atlas-mob-storage-vault` packages a practical mobile workflows exercise in Go. The emphasis is on deterministic behavior, a small public API, and examples that explain the tradeoffs.

## How I Read Atlas Mob Storage Vault

The useful thing to inspect here is how the same score rule is represented in code, metadata, and examples. If those three pieces disagree, the audit script should make the drift visible.

## Problem Shape

This is not a wrapper around a service. It is a self-contained project that shows how the model behaves when demand, capacity, latency, risk, and weight move in different directions.

## Main Behaviors

- Models local state with deterministic scoring and explicit review decisions.
- Uses fixture data to keep sync pressure changes visible in code review.
- Includes extended examples for form constraints, including `surge` and `degraded`.
- Documents offline paths tradeoffs in `docs/operations.md`.
- Runs locally with a single verification command and no external credentials.

## Internal Model

The core is a scoring model over demand, capacity, latency, risk, and weight. That keeps local state, sync pressure, and form constraints in one explicit decision path. The threshold is 166, with risk penalty 4, latency penalty 3, and weight bonus 3. The Go layout uses small packages and table-oriented tests so the behavior stays easy to follow.

## Repository Map

- `policy`: Go package with the core model
- `cmd`: small command entry point
- `fixtures`: compact golden scenarios
- `examples`: expanded scenario set
- `metadata`: project constants and verification metadata
- `docs`: operations and extension notes
- `scripts`: local verification and audit commands
- `go.mod`: Go module metadata

## Run It Locally

Clone the repository, enter the directory, and run the verifier. No database server, cloud account, or token is required.

## How To Run It

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/verify.ps1
```

This runs the language-level build or test path against the compact fixture set.

## Validation

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/audit.ps1
```

The audit command checks repository structure and README constraints before it delegates to the verifier.

## Scenario Walkthrough

The extended cases are not random smoke tests. `degraded` keeps pressure on the review path, while `surge` shows the model when capacity and weight are strong enough to clear the threshold.

## Known Edges

The fixture set is deliberately small. That keeps the review surface clear, but it also means the model should not be treated as a complete domain simulator.

## Follow-Up Work

- Split the scoring constants into a typed configuration object and validate it before use.
- Add a comparison mode that shows how decisions change when one signal is adjusted.
- Add a loader for `examples/extended_cases.csv` and promote selected cases into the language test suite.
- Add one more mobile workflows fixture that focuses on a malformed or borderline input.
