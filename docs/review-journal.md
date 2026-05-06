# Review Journal

The repository goal stays the same: create a Go reference implementation for storage workflows, centered on constraint solving, bounded scenario files, and conflict explanations. This note explains the added review angle.

The local checks classify each case as `ship`, `watch`, or `hold`. That gives the project a small review vocabulary that matches its mobile workflows focus without claiming live deployment or external usage.

## Cases

- `baseline`: `form pressure`, score 177, lane `ship`
- `stress`: `sync drift`, score 159, lane `ship`
- `edge`: `local state`, score 204, lane `ship`
- `recovery`: `conflict cost`, score 150, lane `ship`
- `stale`: `form pressure`, score 195, lane `ship`

## Note

The useful failure mode here is a wrong decision on a named case, not a vague style disagreement.
