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

### Commit Message Format

This project uses [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) for release automation. The commit message should be structured as follows:

```
<type>[optional scope]: <description>

[optional body]

[optional footer]
```

The `<type>` is mandatory and must be one of the following:

*   **feat**: A new feature
*   **fix**: A bug fix
*   **docs**: Documentation only changes
*   **style**: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
*   **refactor**: A code change that neither fixes a bug nor adds a feature
*   **perf**: A code change that improves performance
*   **test**: Adding missing tests or correcting existing tests
*   **build**: Changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm)
*   **ci**: Changes to our CI configuration files and scripts (example scopes: Travis, Circle, BrowserStack, SauceLabs)
*   **chore**: Other changes that don't modify src or test files
*   **revert**: Reverts a previous commit

A `!` appended to the type/scope indicates a breaking change, which will result in a major version bump.

For example:

```
feat: add user authentication endpoint

This commit introduces a new endpoint for user authentication.
It includes the necessary logic for handling user login and registration.
```

```
fix(api)!: correct handling of user sessions

BREAKING CHANGE: The session management has been overhauled.
Refer to the documentation for the new session handling mechanism.
```

This ensures that versioning and releases can be automated correctly.

### Versioning

The package `version` should follow [semver](https://semver.org/).

Currently, the package is under active development and not ready for general use. During this period it will remain below 1.0.0

### Pre-commit

This repo supports the [pre-commit](https://pre-commit.com) framework. Functionality may be incomplete, and you are welcome to file an issue or open a PR

### Taskfile

A `Taskfile` is provided along with a set of `.taskfiles` for use with the [go-task](https://taskfile.dev/) utility. Full capabilities may be incomplete, but the intention is to help with common tasks like regenerating ent schemas or protobufs.

### Entgo

This project uses [entgo.io](https://entgo.io) as an ORM and primarily focuses on postgresql as a backing datastore. To add a new schema, or make changes to the existing schema, you will focus on the `ent/schema` directory. This is the only place under the `ent` directory where you will make any manual changes. The rest of the code is generated.

After making any changes in the `ent/schemas` directory, you must regenerate the files to keep them in sync. To do this, you can use the existing taskfile task `task schemas:ent`
This will also maintain a checksum file under the `.task` directory. Ensure you commit this along with the rest of the generated files when committing schema changes.

### Env variables

For local development, I keep a `.env.dev` file that looks like this for minimal config testing

```bash
export botToken=botTokenViaDiscord
export appId=appIdViaDiscord
export GOOGLE_APPLICATION_CREDENTIALS=somefile.json
```
