> This distribution provides a *very* simple and slightly opinionated Go
> package helpful to create [slog.Logger][LOGGER] instances. I found myself
> using the same patterns repeatedly, and decided to lift the code into
> this independent package. Unless you, somehow, find this convenient,
> there is probably no reason why you should actually use this 🤷

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/license/mit)
[![codecov](https://codecov.io/github/carloscarnero/go-logger/graph/badge.svg?token=Bg1xghG6HS)](https://codecov.io/github/carloscarnero/go-logger)

Changes to this project will be documented in this file, in a format based
on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/).

## [1.1.0] - 2025-02-12

### Added

* Accepts the `NONE` level to discard all and any output. It is
  functionally equivalent to using `nil` or `io.Discard` as the writer of
  the logger instance.

### Changed

* The Go version required to build is 1.24.0 since that's the version that
  introduced `slog.DiscardHandler`.

## [1.0.1] - 2025-01-13

### Changed

* Simplified the generated error values on invalid conditions. Removed from
  the public the `ErrLogger` package variable (this type of action *should*
  bring about a new major version but an analysis of the dependents of this
  library provided that it was safe to do so.)

* Included `2025` in the source code copyright notices.

## [1.0.0] - 2024-12-03

This release is considered the first stable release. It is also when this
`CHANGELOG.md` file is introduced; a decision made early in the project's
life cycle so as not to polute too much the human-readable history (more
details, most of them boring, are always available in the actual Git
commit history.)

[1.1.0]: https://github.com/carloscarnero/go-logger/releases/tag/v1.1.0
[1.0.1]: https://github.com/carloscarnero/go-logger/releases/tag/v1.0.1
[1.0.0]: https://github.com/carloscarnero/go-logger/releases/tag/v1.0.0
