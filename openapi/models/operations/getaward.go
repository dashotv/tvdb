// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"

	"github.com/dashotv/tvdb/openapi/models/shared"
)

type GetAwardRequest struct {
	// id
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetAwardRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

// GetAward200ApplicationJSON - response
type GetAward200ApplicationJSON struct {
	// base award record
	Data   *shared.AwardBaseRecord `json:"data,omitempty"`
	Status *string                 `json:"status,omitempty"`
}

func (o *GetAward200ApplicationJSON) GetData() *shared.AwardBaseRecord {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetAward200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetAwardResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	GetAward200ApplicationJSONObject *GetAward200ApplicationJSON
}

func (o *GetAwardResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAwardResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAwardResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAwardResponse) GetGetAward200ApplicationJSONObject() *GetAward200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetAward200ApplicationJSONObject
}
