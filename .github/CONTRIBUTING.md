# Contributing Guide

**Working on your first Pull Request?** You can learn how from this *free*
series [How to Contribute to an Open Source Project on GitHub][First PR]

## Major changes

Prior to starting work on a major change please open an [issue][]. This allows
you to connect with a maintainer to verify that your change will get merged
upstream. Doing so might save you a lot of wasted effort.

## How to contribute

* Fork the project from the `master` branch and submit a Pull Request (PR)
  * Explain what the PR fixes or improves
* Use sensible commit messages
  * If your PR fixes a separate issue number, include it in the commit message
* Use a sensible number of commits as well
  * e.g. Your PR should not have 100s of commits

## Things to keep in mind

* Smaller Pull Requests are likely to be merged more quickly than bigger changes
* This project is using a [KISS Workflow][]
  * Pull Requests and bugfixes are directly merged into `master` after sanity testing
  * `master` is considered the main developer branch
  * the version tags are considered stable and frozen
* This project is using [Semantic Versioning 2.0.0](http://semver.org/)
  * If a bugfix or PR is *not* trivial it will likely end up in the next **MINOR** version
  * If a bugfix or PR *is* trivial *or* critical it will likely end up in the next **PATCH** version
* Some tests are OS specific

## Commit messages

* Squashing to 1 commit is **not** required at this time
* Use sensible commit messages (when in doubt: `git log`)
* Use a sensible number of commit messages
* If your PR fixes a specific issue number, include it in the commit message: `"Fixes XYZ error (fixes #123)"`

## Code standards

* Try to follow [Effective Go Style Guide][]
* For all new code try to include an effective test
* Lint and test code
  * `make lint`
  * `make test`

_This file was based off work from [Ryan L McIntyre][contributing_ref]_

<!-- link references -->

[issue]: https://github.com/clburlison/bakeit/issues
[KISS Workflow]: https://en.wikipedia.org/wiki/KISS_principle
[Effective Go Style Guide]: https://en.wikipedia.org/wiki/KISS_principle
[First PR]: https://egghead.io/series/how-to-contribute-to-an-open-source-project-on-github
[contributing_ref]: https://github.com/ryanoasis/nerd-fonts/blob/1.2.0/contributing.md
