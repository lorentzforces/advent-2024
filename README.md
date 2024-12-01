# advent-2024

_**Advent of Code 2024**_

This year I'm planning to work a little slower than I tried in previous years, and to let myself be a little less strict about talking with friends about solutions if I get stuck (or if they finished first and want to talk about it). I might skip a day if it ends up being particularly onerous or requires special knowledge (like Day 8/Part 2 last year). I'm trying to fixate less on perfection and therefore do more puzzles.

I'm using Golang this year - as usual, the general goal is to do as much implementation as possible with the standard library, dipping into external dependencies for some super basic niceties (like test assertions).

## Building the project

### Build requirements:

- a Golang installation (built & tested on Go v1.23)
- an internet connection to download dependencies (only necessary if dependencies have changed or this is the first build)
- a `make` installation. This project is built with GNU make v4; full compatibility with other versions of make (such as that shipped by Apple) is not guaranteed, but it _should_ be broadly compatible.

To build the project, simply run `make` in the project's root directory to build the output executable.

> _Note: running with `make` is not strictly necessary. Reference the provided `Makefile` for typical development commands._
