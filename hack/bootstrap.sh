#!/usr/bin/env bash

set -euo pipefail

echo "Bootstrapping day=$DAY"

if [[ ${#DAY} -eq 1 ]]; then
    DAY_DIR="day0$DAY"
else
    DAY_DIR="day$DAY"
fi

mkdir "$DAY_DIR"
cat <<EOF > "$DAY_DIR/main.go"
package main

import _ "embed"

//go:embed input.txt
var input string

func main() {

}
EOF
