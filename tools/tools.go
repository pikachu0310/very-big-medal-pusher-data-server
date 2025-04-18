//go:build tools
// +build tools

package main

import (
	_ "github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen"
	_ "github.com/sqlc-dev/sqlc/cmd/sqlc"
)

//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=server.cfg.yaml ../openapi/openapi.yaml
//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=models.cfg.yaml ../openapi/openapi.yaml

//go:generate go run github.com/sqlc-dev/sqlc/cmd/sqlc generate
