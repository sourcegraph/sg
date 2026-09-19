# AGENTS.md

## What this repo is

This repository holds **binary releases only** for `sg`, the Sourcegraph
developer tool. It contains no application source code. The actual `sg` source
lives in the Sourcegraph monorepo at
[`dev/sg`](https://github.com/sourcegraph/sourcegraph/tree/main/dev/sg#readme).

The only tracked file is `README.md`, which records the latest release
timestamp. Release binaries are distributed via GitHub releases rather than
committed to the working tree.

## Working in this repo

- There is **no build, test, or lint step** here — there is no code to compile.
- Do **not** add source code, package manifests, or CI build pipelines; changes
  to `sg` itself belong in the upstream Sourcegraph monorepo.
- Releases are produced by upstream automation. The `Latest release:` line in
  `README.md` is updated as part of that release process.

## Conventions

- Keep changes minimal and scoped to release/documentation metadata.
- For anything related to `sg`'s behavior, features, or bugs, refer to the
  upstream project at `github.com/sourcegraph/sourcegraph` (`dev/sg`).
