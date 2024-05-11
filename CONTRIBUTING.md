# Contributing Guidelines

Contributions are welcome via GitHub pull requests. This document outlines the process to help get your contribution accepted.

## Sign off Your Work

The Developer Certificate of Origin (DCO) is a lightweight way for contributors to certify that they wrote or otherwise have the right to submit the code they are contributing to the project. Here is the full text of the [DCO](http://developercertificate.org/). Contributors must sign-off that they adhere to these requirements by adding a `Signed-off-by` line to commit messages.

```text
This is my commit message

Signed-off-by: Random J Developer <random@developer.example.org>
```

See `git help commit`:

```text
-s, --signoff
    Add Signed-off-by line by the committer at the end of the commit log
    message. The meaning of a signoff depends on the project, but it typically
    certifies that committer has the rights to submit this work under the same
    license and agrees to a Developer Certificate of Origin (see
    http://developercertificate.org/ for more information).
```

## How to Contribute

1. Fork this repository, develop, and test your changes
1. Remember to sign off your commits as described above
1. Submit a pull request

***NOTE***: In order to make testing and merging of PRs easier, please submit changes to multiple components in separate PRs.

### PR Expectations

- Only solve one problem at a time
- Write clean code (See [Effective Go](https://go.dev/doc/effective_go))
- Write testcases for your code changes
- Include clear documentation especially via ([GoDoc comments](https://tip.golang.org/doc/comment))
- Ensure the PR is linked to an issue. If your solution has a design element, please review that in the issue discussion before working on the PR

### Versioning

The package `version` should follow [semver](https://semver.org/).

Currently, the package is under active development and not ready for general use. During this period it will remain below 1.0.0

### Pre-commit

This repo supports the [pre-commit](https://pre-commit.com) framework. Functionality may be incomplete, and you are welcome to file an issue or open a PR

### Taskfile

A `Taskfile` is provided along with a set of `.taskfiles` for use with the [go-task](https://taskfile.dev/) utility. Full capabilities may be incomplete, but the intention is to help with common tasks like regenerating ent schemas or protobufs
