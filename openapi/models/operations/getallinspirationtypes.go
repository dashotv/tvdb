// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"

	"github.com/dashotv/tvdb/openapi/models/shared"
)

// GetAllInspirationTypes200ApplicationJSON - response
type GetAllInspirationTypes200ApplicationJSON struct {
	Data   []shared.InspirationType `json:"data,omitempty"`
	Status *string                  `json:"status,omitempty"`
}

func (o *GetAllInspirationTypes200ApplicationJSON) GetData() []shared.InspirationType {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetAllInspirationTypes200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetAllInspirationTypesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	GetAllInspirationTypes200ApplicationJSONObject *GetAllInspirationTypes200ApplicationJSON
}

func (o *GetAllInspirationTypesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAllInspirationTypesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAllInspirationTypesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAllInspirationTypesResponse) GetGetAllInspirationTypes200ApplicationJSONObject() *GetAllInspirationTypes200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetAllInspirationTypes200ApplicationJSONObject
}