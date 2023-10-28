// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"

	"github.com/dashotv/tvdb/openapi/models/shared"
)

type GetGenreBaseRequest struct {
	// id
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetGenreBaseRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

// GetGenreBase200ApplicationJSON - response
type GetGenreBase200ApplicationJSON struct {
	// base genre record
	Data   *shared.GenreBaseRecord `json:"data,omitempty"`
	Status *string                 `json:"status,omitempty"`
}

func (o *GetGenreBase200ApplicationJSON) GetData() *shared.GenreBaseRecord {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetGenreBase200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetGenreBaseResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	GetGenreBase200ApplicationJSONObject *GetGenreBase200ApplicationJSON
}

func (o *GetGenreBaseResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetGenreBaseResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetGenreBaseResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetGenreBaseResponse) GetGetGenreBase200ApplicationJSONObject() *GetGenreBase200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetGenreBase200ApplicationJSONObject
}
