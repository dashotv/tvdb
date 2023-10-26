// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tvdb/openapi/models/shared"
	"net/http"
)

// GetAllAwards200ApplicationJSON - response
type GetAllAwards200ApplicationJSON struct {
	Data   []shared.AwardBaseRecord `json:"data,omitempty"`
	Status *string                  `json:"status,omitempty"`
}

func (o *GetAllAwards200ApplicationJSON) GetData() []shared.AwardBaseRecord {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetAllAwards200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetAllAwardsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	GetAllAwards200ApplicationJSONObject *GetAllAwards200ApplicationJSON
}

func (o *GetAllAwardsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAllAwardsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAllAwardsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAllAwardsResponse) GetGetAllAwards200ApplicationJSONObject() *GetAllAwards200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetAllAwards200ApplicationJSONObject
}
