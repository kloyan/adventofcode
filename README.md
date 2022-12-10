# Advent of Code Solutions

Go solutions to the Advent of Code puzzles. Check out https://adventofcode.com.

My goal is to learn Go and improve my problem solving skills.

## Usage

The repository includes a helper script which could be used to download the puzzle input. For this to work it expects the `COOKIE` env variable to be set. Its value is the session cookie generated after successful authentication in the Advent of Code website.

```shell
# Create the boilerplate code for a given day
make bootstrap DAY=1

# Download the input and solve the problem for a given day
make run COOKIE=abcd DAY=1
# > Getting input for day=1
# > Running solution for day=1
# > ------------------
# > Answer for Part 1: 72718
# > Answer for Part 2: 213089
```
