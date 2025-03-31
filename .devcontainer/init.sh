#!/bin/bash

set -e

go install github.com/rakyll/hey@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
