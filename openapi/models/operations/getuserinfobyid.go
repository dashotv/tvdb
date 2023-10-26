// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tvdb/openapi/models/shared"
	"net/http"
)

type GetUserInfoByIDRequest struct {
	// id
	ID float64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetUserInfoByIDRequest) GetID() float64 {
	if o == nil {
		return 0.0
	}
	return o.ID
}

// GetUserInfoByID200ApplicationJSON - response
type GetUserInfoByID200ApplicationJSON struct {
	Data   []shared.UserInfo `json:"data,omitempty"`
	Status *string           `json:"status,omitempty"`
}

func (o *GetUserInfoByID200ApplicationJSON) GetData() []shared.UserInfo {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetUserInfoByID200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetUserInfoByIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	GetUserInfoByID200ApplicationJSONObject *GetUserInfoByID200ApplicationJSON
}

func (o *GetUserInfoByIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetUserInfoByIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetUserInfoByIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetUserInfoByIDResponse) GetGetUserInfoByID200ApplicationJSONObject() *GetUserInfoByID200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetUserInfoByID200ApplicationJSONObject
}
