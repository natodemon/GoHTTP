# GoHTTP

## Purpose

This is an http server written in Golang as a component of another project, please see [natodemon/GoBench/README.md](https://github.com/natodemon/GoBench) for more details. The program simply serves an html index file placed in the *web* directory.

## Usage

Please clone this repository and run one of the following commands:

    go run httpserver.go [-p=port]
    ./httpserver [-p=port]

If no port is provided, the default is 8080