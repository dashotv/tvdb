// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"

	"github.com/dashotv/tvdb/openapi/models/shared"
)

type GetMovieBaseRequest struct {
	// id
	ID float64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetMovieBaseRequest) GetID() float64 {
	if o == nil {
		return 0.0
	}
	return o.ID
}

// GetMovieBaseResponseBody - response
type GetMovieBaseResponseBody struct {
	// base movie record
	Data   *shared.MovieBaseRecord `json:"data,omitempty"`
	Status *string                 `json:"status,omitempty"`
}

func (o *GetMovieBaseResponseBody) GetData() *shared.MovieBaseRecord {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetMovieBaseResponseBody) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetMovieBaseResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	Object *GetMovieBaseResponseBody
}

func (o *GetMovieBaseResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMovieBaseResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMovieBaseResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetMovieBaseResponse) GetObject() *GetMovieBaseResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
