# Changelog

## Unreleased

## [v0.3.2] 2023-03-16 Thu

- Add: refactor code to be more standard.

## [v0.3.1]

- Add: upgrade modules.

## [v0.3.0]

- Add [#7]: Sort files first by date in the file name (if present), then
            by modification date.
- Add [#6]: Optionally, do not delete any files, instead move them to an
            ``archive/delete-me`` directory.

## [v0.2.0]

- Add [#5]: support multiple files regex patterns in one directory. It allows
            to organize separate categories of files independently from each
            other.

## [v0.1.0]

- Add [#4]: move/delete files according to their groups
- Add [#3]: categorize files according to how old they are
- Add [#2]: collect files from backup directories
- Add [#1]: read config file

## Footnotes

This document follows [changelog guidelines]

[v0.3.0]: https://github.com/dimus/backme/compare/v0.2.0...v0.3.0
[v0.2.0]: https://github.com/dimus/backme/compare/v0.1.0...v0.2.0
[v0.1.0]: https://github.com/dimus/backme/tree/v0.1.0

[#6]: https://github.com/dimus/backme/issues/6
[#5]: https://github.com/dimus/backme/issues/5
[#4]: https://github.com/dimus/backme/issues/4
[#3]: https://github.com/dimus/backme/issues/3
[#2]: https://github.com/dimus/backme/issues/2
[#1]: https://github.com/dimus/backme/issues/1

[changelog guidelines]: https://github.com/olivierlacan/keep-a-changelog
