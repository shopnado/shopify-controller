@echo off
if "%1%"=="" (
    set cmd=ci
) else (
    set cmd=%1%
)
if "%cmd%"=="ci" (
    goto :eof
)
if "%cmd%"=="generate" (
    docker run --rm -it -v %cd%:/shopnado -w /shopnado -e GOPATH="" golang:1.13-alpine go generate
    goto :eof
)