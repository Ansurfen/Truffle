@echo off

go env -w GOOS=linux

go build main.go

go env -w GOOS=windows

go build main.go