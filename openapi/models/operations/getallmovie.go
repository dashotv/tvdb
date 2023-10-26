// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"

	"github.com/dashotv/tvdb/openapi/models/shared"
)

type GetAllMovieRequest struct {
	// page number
	Page *float64 `queryParam:"style=form,explode=true,name=page"`
}

func (o *GetAllMovieRequest) GetPage() *float64 {
	if o == nil {
		return nil
	}
	return o.Page
}

// GetAllMovie200ApplicationJSON - response
type GetAllMovie200ApplicationJSON struct {
	Data []shared.MovieBaseRecord `json:"data,omitempty"`
	// Links for next, previous and current record
	Links  *shared.Links `json:"links,omitempty"`
	Status *string       `json:"status,omitempty"`
}

func (o *GetAllMovie200ApplicationJSON) GetData() []shared.MovieBaseRecord {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetAllMovie200ApplicationJSON) GetLinks() *shared.Links {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *GetAllMovie200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetAllMovieResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	GetAllMovie200ApplicationJSONObject *GetAllMovie200ApplicationJSON
}

func (o *GetAllMovieResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAllMovieResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAllMovieResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAllMovieResponse) GetGetAllMovie200ApplicationJSONObject() *GetAllMovie200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetAllMovie200ApplicationJSONObject
}