// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"

	"github.com/dashotv/tvdb/openapi/models/shared"
)

// GetAllContentRatings200ApplicationJSON - response
type GetAllContentRatings200ApplicationJSON struct {
	Data   []shared.ContentRating `json:"data,omitempty"`
	Status *string                `json:"status,omitempty"`
}

func (o *GetAllContentRatings200ApplicationJSON) GetData() []shared.ContentRating {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetAllContentRatings200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetAllContentRatingsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	GetAllContentRatings200ApplicationJSONObject *GetAllContentRatings200ApplicationJSON
}

func (o *GetAllContentRatingsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAllContentRatingsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAllContentRatingsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAllContentRatingsResponse) GetGetAllContentRatings200ApplicationJSONObject() *GetAllContentRatings200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetAllContentRatings200ApplicationJSONObject
}