// Package httpservice provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/flypay/go-kit/v4 version (devel) DO NOT EDIT.
package httpservice

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get users with birthdays today
	// (GET /birthday)
	GetUsersWithBirthdayToday(ctx echo.Context, params GetUsersWithBirthdayTodayParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetUsersWithBirthdayToday converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsersWithBirthdayToday(ctx echo.Context) error {
	var err error

	ctx.Set(ApiKeyAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetUsersWithBirthdayTodayParams
	// ------------- Required query parameter "date" -------------

	err = runtime.BindQueryParameter("form", true, true, "date", ctx.QueryParams(), &params.Date)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter date: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUsersWithBirthdayToday(ctx, params)
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

	router.GET(baseURL+"/birthday", wrapper.GetUsersWithBirthdayToday)
}

type GetUsersWithBirthdayTodayRequestObject struct {
	Params GetUsersWithBirthdayTodayParams
}

type GetUsersWithBirthdayTodayResponseObject interface {
	VisitGetUsersWithBirthdayTodayResponse(w http.ResponseWriter) error
}

type GetUsersWithBirthdayToday200JSONResponse []BirthdayRecord

func (response GetUsersWithBirthdayToday200JSONResponse) VisitGetUsersWithBirthdayTodayResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Get users with birthdays today
	// (GET /birthday)
	GetUsersWithBirthdayToday(ctx context.Context, request GetUsersWithBirthdayTodayRequestObject) (GetUsersWithBirthdayTodayResponseObject, error)
}

type (
	StrictHandlerFunc    = strictecho.StrictEchoHandlerFunc
	StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc
)

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetUsersWithBirthdayToday operation middleware
func (sh *strictHandler) GetUsersWithBirthdayToday(ctx echo.Context, params GetUsersWithBirthdayTodayParams) error {
	var request GetUsersWithBirthdayTodayRequestObject

	request.Params = params

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetUsersWithBirthdayToday(ctx.Request().Context(), request.(GetUsersWithBirthdayTodayRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUsersWithBirthdayToday")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetUsersWithBirthdayTodayResponseObject); ok {
		return validResponse.VisitGetUsersWithBirthdayTodayResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{
	"H4sIAAAAAAAC/6yTT2/UMBDFv4o1cEw2C9xyK0Ig1Atqi0Cq9uB1JhuX+E/H41Rhle+O7GT/sBQJBKfd",
	"JOP3PPN+swfljHcWLQeo9xBUh0bmv281cdfI8QaVoya98eQ8EmvM33V+x6NHqCEwabuDaSqA8DFqwgbq",
	"+1SzKQ41bvuAinNNQBVJ83ib7Ga5K6+vcbyK3GVxCzV0KBskKMBKkwS+lu/7kcsrr8trHOEoLPNRmJKy",
	"tq1LAg0GRdqzdklp6xwraXz53WHZ9u4Jqdwu/ZXWsW51NmLNPf5F/YAUZof1ar16BVMBzqOVXkMNb1br",
	"1RoK8JK73GJ1kEgPO+Rf73mDTBoHDCIGpCBackZwh6KRLLcyoHjqXEBxEBI6CHZJMRuTTDofG6jhA/Ln",
	"JPFFc3dI8m6p9JKkQUYKUN9fXqGRjIKdUB2qb8JIVp22u6NjgGIO5zEijads0jE4D58pYrHw9Bwom1Qc",
	"vLNhzv/1ep1+lLOMNo9Get9rlVuqHkK63P5MTzOafPAlYQs1vKhOJFcLxtUFw9MJGSK5EPNz+7dRKQyh",
	"jb04DnRGNhojaZxHu8TzpLk7TWZJYjoHPA/4HO37Teo8IA3Pj/8TuSaq9CDQDpqcNWkcBUTq00ow+1BX",
	"1R8BumrTuvhecuvIrJQzMBWXhrcsdyng/+QWZrnfmL3DAXvnk8m/G5b7xF7wUuE0mzc4ZOMCBklabvsZ",
	"rWPZvHCtjD1DDSnDTO8lA3cdiuMZ4dp5A8+unuLTCgsRQ5R9P4rRRRIHvVVCYDP9CAAA///b0sh1WwUA",
	"AA==",
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
