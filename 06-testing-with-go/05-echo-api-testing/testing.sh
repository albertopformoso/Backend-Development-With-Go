#!/bin/sh

for d in $(go list ./...); do
    echo -e "\n\nTesting package\033[1;36m $d \033[0;37m"
    go test -v $d
done
