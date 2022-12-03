#!/usr/bin/env bash

set -euo pipefail

echo "Running solution for day=$DAY"

if [[ ${#DAY} -eq 1 ]]; then
    DAY_DIR="day0$DAY"
else
    DAY_DIR="day$DAY"
fi

echo "------------------"
CGO_ENABLED=0 go run "$DAY_DIR/main.go"
