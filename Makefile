SHELL=bash

default: run

clean:
	@ find . -type f -name "input.txt" -exec rm -v {} \;

.PHONY: check-cookie
check-cookie:
	@ test -n "${COOKIE}" || (echo "env COOKIE not set"; exit 1)

.PHONY: check-day
check-day:
	@ test -n "${DAY}" || (echo "env DAY not set"; exit 1)

.PHONY: get-input
get-input: check-cookie check-day
	@ hack/get-input.sh

.PHONY: run
run: get-input
	@ hack/run.sh

.PHONY: bootstrap
bootstrap: check-day
	@ hack/bootstrap.sh
