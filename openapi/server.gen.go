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

	"H4sIAAAAAAAC/6yXT08bRxTAv4o17XGxAeMefKNpRThEQrTNJUXWYA/2pt4/mZ1FsZAlZldJSAGFUFpo",
	"oypERMEtBZpUaWniwId5LH9O+QrVzHq9a7w0xN7LarU77/fevDfvz8yhoqGZhk50ZqH8HLKKFaJh+TqG",
	"NfIFZli8m9QwCWUqkX8mCxop4ap4ZTWToDxSdUbKhCIF3R0wsKkOFI0SKRN9gNxlFA8wXJaCpWmUR7Ql",
	"Xa8raIbMEhrDqSuoTFjBqqh0mvSoJwIQqip4lvRldwQgeYZ9ieVqKfLdYlTVy/Kz3pf6trhQLt8KQz2i",
	"AumQNNwXaThCyvZFykZII32RRiKkXF+knE9SdZuR+IAbNusrtKG80GSRoqGX4jX557kwbdeyg4M9autk",
	"SI1VgxUqKusVGIhLFsOUFcSnXmkhQPBsi9BCXEZdhRUIC9AsoZZq6DGpWVeCL8b0bVIUemW6zhgSRKwi",
	"VU0mZdHNyWsVzE7XGuA+A3ce3B1w18F1wfkH3C1wm+Duy+cOuA9hfgmcfXDeAd8GvgX8MfAHF9c7r+T6",
	"TZhfBr4n//4BznNwH8i/R+Csns/z46NnwH8GZ3F0YhzmnW91pCCmsqqw+iahtdTnajl1Q5yh1IRtVQhN",
	"ifKdEvU7NToxjiL7R0PpwfSQPLcm0bGpojzKpgfTWaQgE7OKdF+m1Kr8ZcK6vXD9xui103evvJXl47cb",
	"wIXhp69fnjUWTtx73uZLcFbHvvwa+HbUcCT1USwQ4yWUR2OEyfYitFKsEUaohfK35pAqVNyxCa0hBelY",
	"E3sMrFcQJXdslZISyjNqE6XVsWLDGo8KDkUCqEhLuAKtnceX4dolPglYWFSSoLVzPAlYpDMnYltYMRLB",
	"ddTHJIhBv02ONZwgK5sgayRBVi4R1mSSSeB35kS26I8TSaDkIJoEyB/GP5bU3RkGvro+Opz7zG8Q75sL",
	"55v3vTePvCdvvL1fwH3bqr9nh03gR8C3T9eeniysvG8+REq819XyR9XqKbHYMg3d8m8rw/6k1Gnnye6W",
	"d3AAfOf46FdvdwP4j+AsAT8Evg78qWiOI3FixwfLJ7tbwH/3N5cyaArcx+D+JscB0a6FaC5OFJzXstGv",
	"gNv07t87dxvgNKRk0x+cbE3DtCZXBiNB1wwgV2ZMsdGwN3e11Ql/PvqwF6I2nTaanrt8fLB70Rp3Q0wz",
	"zr/gcnC25eyy79tBsf6dqvvz1mW2TAZrrtTmLYOKIh4Gt0RmsF1lF3st0W0N5W91fgzOb6SNhk1wSrlq",
	"U6+qmnqJEblBpTsHLjlvRUNnRJdOwaZZVYvSLZnblj+DhnSVEU0KfkrJDMqjTzLhlTzTuo9n2pfxcFjF",
	"lOKaP6teiKo4VH+Bsyuff7ZPUVdgO5c5q96jn7zDdT+2IkutzFwrWesX58GuOH8j1ovHeOl/pjoxX/Y3",
	"ifXr7as5Oc6pL6QX/xZPvtedo37NGIlJsg7JpbMXi8CfA18E53tZcJ6A80NXaKIia8AbsSqjIZMXRzob",
	"eNumVZEzjJlWPpMxbauSZhSb6TLWiIXqU/X/AgAA//9p0ICN+REAAA==",
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
