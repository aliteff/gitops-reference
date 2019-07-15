# Git Usage Reference

In this repository, you'll find an example CI file (`.drone.yml`) which defines a pretty basic git branching strategy.

## Branches
* All commits on all branches are built and tested
* All commits to master (which should be **squashed and merged**) builds a docker image and are deployed to "staging"

## Tags
* Git tags which contain `-release` (ex. `v0.1-release`, `v0.1-release-rc1`) are deployed to production
