// Code generated by go-swagger; DO NOT EDIT.

package node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// GetNodeListURL generates an URL for the get node list operation
type GetNodeListURL struct {
	Limit  *float64
	Offset *float64
	Order  *string
	Sortby *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetNodeListURL) WithBasePath(bp string) *GetNodeListURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetNodeListURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetNodeListURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/node"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/v1"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var limit string
	if o.Limit != nil {
		limit = swag.FormatFloat64(*o.Limit)
	}
	if limit != "" {
		qs.Set("limit", limit)
	}

	var offset string
	if o.Offset != nil {
		offset = swag.FormatFloat64(*o.Offset)
	}
	if offset != "" {
		qs.Set("offset", offset)
	}

	var order string
	if o.Order != nil {
		order = *o.Order
	}
	if order != "" {
		qs.Set("order", order)
	}

	var sortby string
	if o.Sortby != nil {
		sortby = *o.Sortby
	}
	if sortby != "" {
		qs.Set("sortby", sortby)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetNodeListURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetNodeListURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetNodeListURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetNodeListURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetNodeListURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetNodeListURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
