#!/bin/sh

go test -coverprofile=profile.out ./...
echo -e "\n\n"
go tool cover -func=profile.out
go tool cover -html=profile.out
