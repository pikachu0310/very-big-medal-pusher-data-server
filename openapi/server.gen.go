// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	. "github.com/pikachu0310/very-big-medal-pusher-data-server/openapi/models"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// ゲームデータを送信
	// (GET /data)
	GetData(ctx echo.Context, params GetDataParams) error
	// ヘルスチェック
	// (GET /ping)
	GetPing(ctx echo.Context) error
	// ランキングを取得
	// (GET /rankings)
	GetRankings(ctx echo.Context, params GetRankingsParams) error
	// ユーザーごとのゲームデータを取得
	// (GET /users/{user_id}/data)
	GetUsersUserIdData(ctx echo.Context, userId string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetData converts echo context to params.
func (w *ServerInterfaceWrapper) GetData(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetDataParams
	// ------------- Required query parameter "version" -------------

	err = runtime.BindQueryParameter("form", true, true, "version", ctx.QueryParams(), &params.Version)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter version: %s", err))
	}

	// ------------- Required query parameter "user_id" -------------

	err = runtime.BindQueryParameter("form", true, true, "user_id", ctx.QueryParams(), &params.UserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter user_id: %s", err))
	}

	// ------------- Required query parameter "have_medal" -------------

	err = runtime.BindQueryParameter("form", true, true, "have_medal", ctx.QueryParams(), &params.HaveMedal)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter have_medal: %s", err))
	}

	// ------------- Required query parameter "in_medal" -------------

	err = runtime.BindQueryParameter("form", true, true, "in_medal", ctx.QueryParams(), &params.InMedal)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter in_medal: %s", err))
	}

	// ------------- Required query parameter "out_medal" -------------

	err = runtime.BindQueryParameter("form", true, true, "out_medal", ctx.QueryParams(), &params.OutMedal)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter out_medal: %s", err))
	}

	// ------------- Required query parameter "slot_hit" -------------

	err = runtime.BindQueryParameter("form", true, true, "slot_hit", ctx.QueryParams(), &params.SlotHit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter slot_hit: %s", err))
	}

	// ------------- Required query parameter "get_shirbe" -------------

	err = runtime.BindQueryParameter("form", true, true, "get_shirbe", ctx.QueryParams(), &params.GetShirbe)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter get_shirbe: %s", err))
	}

	// ------------- Required query parameter "start_slot" -------------

	err = runtime.BindQueryParameter("form", true, true, "start_slot", ctx.QueryParams(), &params.StartSlot)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter start_slot: %s", err))
	}

	// ------------- Required query parameter "shirbe_buy300" -------------

	err = runtime.BindQueryParameter("form", true, true, "shirbe_buy300", ctx.QueryParams(), &params.ShirbeBuy300)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter shirbe_buy300: %s", err))
	}

	// ------------- Required query parameter "medal_1" -------------

	err = runtime.BindQueryParameter("form", true, true, "medal_1", ctx.QueryParams(), &params.Medal1)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter medal_1: %s", err))
	}

	// ------------- Required query parameter "medal_2" -------------

	err = runtime.BindQueryParameter("form", true, true, "medal_2", ctx.QueryParams(), &params.Medal2)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter medal_2: %s", err))
	}

	// ------------- Required query parameter "medal_3" -------------

	err = runtime.BindQueryParameter("form", true, true, "medal_3", ctx.QueryParams(), &params.Medal3)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter medal_3: %s", err))
	}

	// ------------- Required query parameter "medal_4" -------------

	err = runtime.BindQueryParameter("form", true, true, "medal_4", ctx.QueryParams(), &params.Medal4)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter medal_4: %s", err))
	}

	// ------------- Required query parameter "medal_5" -------------

	err = runtime.BindQueryParameter("form", true, true, "medal_5", ctx.QueryParams(), &params.Medal5)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter medal_5: %s", err))
	}

	// ------------- Required query parameter "R_medal" -------------

	err = runtime.BindQueryParameter("form", true, true, "R_medal", ctx.QueryParams(), &params.RMedal)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter R_medal: %s", err))
	}

	// ------------- Required query parameter "second" -------------

	err = runtime.BindQueryParameter("form", true, true, "second", ctx.QueryParams(), &params.Second)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter second: %s", err))
	}

	// ------------- Required query parameter "minute" -------------

	err = runtime.BindQueryParameter("form", true, true, "minute", ctx.QueryParams(), &params.Minute)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter minute: %s", err))
	}

	// ------------- Required query parameter "hour" -------------

	err = runtime.BindQueryParameter("form", true, true, "hour", ctx.QueryParams(), &params.Hour)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter hour: %s", err))
	}

	// ------------- Required query parameter "fever" -------------

	err = runtime.BindQueryParameter("form", true, true, "fever", ctx.QueryParams(), &params.Fever)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter fever: %s", err))
	}

	// ------------- Required query parameter "sig" -------------

	err = runtime.BindQueryParameter("form", true, true, "sig", ctx.QueryParams(), &params.Sig)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter sig: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetData(ctx, params)
	return err
}

// GetPing converts echo context to params.
func (w *ServerInterfaceWrapper) GetPing(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetPing(ctx)
	return err
}

// GetRankings converts echo context to params.
func (w *ServerInterfaceWrapper) GetRankings(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetRankingsParams
	// ------------- Optional query parameter "sort" -------------

	err = runtime.BindQueryParameter("form", true, false, "sort", ctx.QueryParams(), &params.Sort)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter sort: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetRankings(ctx, params)
	return err
}

// GetUsersUserIdData converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsersUserIdData(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "user_id" -------------
	var userId string

	err = runtime.BindStyledParameterWithOptions("simple", "user_id", ctx.Param("user_id"), &userId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter user_id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUsersUserIdData(ctx, userId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/data", wrapper.GetData)
	router.GET(baseURL+"/ping", wrapper.GetPing)
	router.GET(baseURL+"/rankings", wrapper.GetRankings)
	router.GET(baseURL+"/users/:user_id/data", wrapper.GetUsersUserIdData)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/6yX3U4bxxfAX8Wa//9ysQnGvfAdTSvCRSRE29ykyBrswZ7U+5HZWRQLWWJ2lYQUUAil",
	"hTaqQkQU3FKgSZWWJg48zGH5uMorVDNre9d4aQnem9Vqd87vnDlnzsfMoqKpW6ZBDG6j/CyyixWiY/U6",
	"inXyGeZYvlvMtAjjlKg/EwWdlHBVvvKaRVAeUYOTMmFIQ/cGTGzRgaJZImViDJB7nOEBjstKsDSF8oi1",
	"pOt1DU2TGcJiOHUNlQkv2BXKpsgV9UQAUlUFz5C+7I4AFM90LrCcGn3p6YhLLeqtcO2KqLZ0SBrqizQU",
	"IWX7ImUjpOG+SMMRUq4vUi4gUcPhJD6ypsP7Cm0oLzXZpGgapXhNwcEtTDm17ODgFbV1M5TGqskLFcqv",
	"CmyLKxbHjBfkp6vSQoDkOTZhBRp1h80ZNcqXYrWFJWiGMJuaRi9I/mx9MafukKLUq9J12lQgYhcZtbiS",
	"RbcmrlcwP1ltgPccvDnwtsFbA88D9y/wNsFrgrenntvgPYK5RXD3wH0PYgvEJognIB6eX+++Vus3YG4J",
	"xK76+xu4L8B7qP4egrtyNieODp+D+BHchZHxMZhzvzaQhjjlVWn1LcJqqU9pOXVTnqHUuGNXCEvJOp2S",
	"hTo1Mj6GIvtH19KD6Wvq3FrEwBZFeZRND6azSEMW5hXlvkypVeLLhPd64cbNkesn71/7y0tH79ZBSMNP",
	"3rw6bcwfe/f9jVfgrox+/iWIrajhSOljWCLGSiiPRglXfURqZVgnnDAb5W/PIipV3HUIqyENGViXe2xb",
	"ryFG7jqUkRLKc+YQrdWaYsMaj2ofigRQkdp/CVonjy/CdUp8ErCwqCRB6+R4ErBIC07EtrBiJILrqo9J",
	"ENv9NjnWUIKsbIKs4QRZuURYE0kmQdCZE9liME4kgVITZxKgYOr+WFJvZxj44sbIUO6ToEF8aM6fbTzw",
	"3z72n771d38C712r/p4eNEEcgtg6WX12PL/8ofkIafFep+WPqtWTcrFtmYYdXEuGgkmp287jnU1/fx/E",
	"9tHhz/7OOojvwV0EcQBiDcQz2RyH48SO9peOdzZB/BpsLmWyFHhPwPtFjQOyXUvRXJwouG9Uo18Gr+k/",
	"uH/mNcBtKMlmMDg5uo5ZTa1sjwQ9M4BambHkRsPe3NNWx4P56L+9ELXppNH0vaWj/Z3z1njrcppx/wZP",
	"gLulZpe9wA6GjW+oEcxbF9ky0V5zqTZvm0wW8TC4JTKNnSo/32uJ4egof7v7Y/v8Rtpo2AQntcs29SrV",
	"6QVG5Aa13hy44LwVTYMTQzkFW1aVFpVbMnfsYAYN6ZQTXQn+n5FplEf/y4R370zr4p3p3LrDYRUzhmvB",
	"rHouqvJQ/QHujnr+3jlFPYHtXuau+I9/8A/WgtjKLLUzs61krZ+fB3vi/JVcLx9jpX+Z6uR82d8k1q+3",
	"L+fkOKe+VF78Uz7Fbm+OBjVjOCbJuiQXT18ugHgBYgHcb1XBeQrudz2hiYqsgmjEqoyGTF0c2Uzb2w6r",
	"ypzh3LLzmYzl2JU0Z9hKl7FObFSfrP8TAAD//7IpiEXiEQAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
