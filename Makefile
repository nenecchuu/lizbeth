#!/usr/bin/env bash

swag:
	# @export PATH="$HOME/go/bin:$PATH"
	@echo "> Generate Swagger Docs"
	# @if ! command -v swag &> /dev/null; then go install github.com/swaggo/swag/cmd/swag ; fi
	@swag init --parseVendor --parseDependency
