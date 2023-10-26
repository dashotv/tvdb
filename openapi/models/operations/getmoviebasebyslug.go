// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tvdb/openapi/models/shared"
	"net/http"
)

type GetMovieBaseBySlugRequest struct {
	// slug
	Slug string `pathParam:"style=simple,explode=false,name=slug"`
}

func (o *GetMovieBaseBySlugRequest) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

// GetMovieBaseBySlug200ApplicationJSON - response
type GetMovieBaseBySlug200ApplicationJSON struct {
	// base movie record
	Data   *shared.MovieBaseRecord `json:"data,omitempty"`
	Status *string                 `json:"status,omitempty"`
}

func (o *GetMovieBaseBySlug200ApplicationJSON) GetData() *shared.MovieBaseRecord {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetMovieBaseBySlug200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetMovieBaseBySlugResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	GetMovieBaseBySlug200ApplicationJSONObject *GetMovieBaseBySlug200ApplicationJSON
}

func (o *GetMovieBaseBySlugResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMovieBaseBySlugResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMovieBaseBySlugResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetMovieBaseBySlugResponse) GetGetMovieBaseBySlug200ApplicationJSONObject() *GetMovieBaseBySlug200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetMovieBaseBySlug200ApplicationJSONObject
}
