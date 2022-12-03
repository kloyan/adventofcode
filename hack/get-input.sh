#!/usr/bin/env bash

set -euo pipefail

if [[ ${#DAY} -eq 1 ]]; then
    DAY_DIR="day0$DAY"
else
    DAY_DIR="day$DAY"
fi

if [[ -f "$DAY_DIR/input.txt" ]]; then
    exit 0
fi

echo "Getting input for day=$DAY"

curl -sSLf \
    -b "session=$COOKIE" \
    -o "$DAY_DIR/input.txt" \
    "https://adventofcode.com/2022/day/$DAY/input"
