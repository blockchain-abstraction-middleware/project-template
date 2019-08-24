#!/bin/bash

gofmt -s -w .
gofmt -r '(a) -> a' -l *.go .
gofmt -r '(a) -> a' -w *.go .
gofmt -r 'α[β:len(α)] -> α[β:]' -w $GOROOT/src .