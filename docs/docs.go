// Package classification Redirect service API
//
// Documentation for Redirect service API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//  - application/x-msgpack
//
//	Produces:
//	- application/json
//  - application/x-msgpack
//
// swagger:meta
package docs

import "github.com/praveen4g0/url-shortner/shortner"

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// swagger:response RedirectServiceResponse
type RedirectsResponseWrapper struct {
	// Newley created Redirect
	// in: body
	Body shortner.Redirect
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters fetchCode
type RedirectCodeParamsWrapper struct {
	// The code to redirect the actual URL
	// in: path
	// required: true
	Code string `json:"code" msgpack:"code"`
}
