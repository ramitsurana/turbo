# How to Contribute

Turbo project is [Apache 2.0 licensed](LICENSE) and accept contributions via
GitHub pull requests.  This document outlines some of the conventions on
development workflow, commit message formatting, contact points and other
resources to make it easier to get your contribution accepted.

### Email and Chat

Currently the project has:
- Email: [ramitsurana@gmail.com](mailto:ramitsurana@gmail.com)

### Getting Started

- Fork the repository on GitHub
- Play with the project, submit bugs, submit patches!

### Contribution Flow

This is a rough outline of what a contributor's workflow looks like:

- Create a topic branch from where you want to base your work (usually master).
- Make commits of logical units.
- Make sure your commit messages are in the proper format (see below).
- Push your changes to a topic branch in your fork of the repository.
- Make sure the [tests](tests/README.md#manually-running-the-tests) pass, and add any new tests as appropriate.
- Submit a pull request to the original repository.
- Submit a comment with the sole content "@reviewer PTAL" (please take a look) in GitHub
  and replace "@reviewer" with the correct recipient.
- When addressing pull request review comments add new commits to the existing pull request or,
  if the added commits are about the same size as the previous commits,
  squash them into the existing commits.
- Once your PR is labelled as "reviewed/lgtm" squash the addressed commits in one commit.
- If your PR addresses multiple subsystems reorganize your PR and create multiple commits per subsystem.
- Your contribution is ready to be merged.

Thanks for your contributions!

### Format of the Pull Request

The pull request title and the first paragraph of the pull request description
is being used to generate the changelog of the next release.

The convention follows the same rules as for commit messages. The PR title reflects the
what and the first paragraph of the PR description reflects the why.
In most cases one can reuse the commit title as the PR title
and the commit messages as the PR description for the PR.

If your PR includes more commits spanning mulitple subsystems one should change the PR title
and the first paragraph of the PR description to reflect a summary of all changes involved.

A large PR must be split into multiple commits, each with clear commit messages.
Intermediate commits should compile and pass tests. Exceptions to non-compilable must have a valid reason, i.e. dependency bumps.
