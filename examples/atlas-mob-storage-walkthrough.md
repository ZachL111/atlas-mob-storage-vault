# Atlas Mob Storage Vault Walkthrough

This note is the quickest way to read the extra review model in `atlas-mob-storage-vault`.

| Case | Focus | Score | Lane |
| --- | --- | ---: | --- |
| baseline | form pressure | 177 | ship |
| stress | sync drift | 159 | ship |
| edge | local state | 204 | ship |
| recovery | conflict cost | 150 | ship |
| stale | form pressure | 195 | ship |

Start with `edge` and `recovery`. They create the widest contrast in this repository's fixture set, which makes them better review anchors than the middle cases.

The next useful expansion would be a malformed fixture around sync drift and conflict cost.
