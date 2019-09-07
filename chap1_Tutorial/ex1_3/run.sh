#!/bin/bash
args=()
for i in {1..1000}
do
    args+=("$i")
done
go test -bench=. -args "${args[@]}"
