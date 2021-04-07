# GoHTTP

## Purpose

This is a very simple http server written in Golang as a component of another project, please see [natodemon/GoBench/README.md](https://github.com/natodemon/GoBench). It simply serves an html index file placed in the *web* directory.

## Usage

    go run httpserver.go [-p=_port_]
    ./httpserver [-p=_port_]

If no port is provided, the default is 8080