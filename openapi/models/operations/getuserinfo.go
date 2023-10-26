// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"

	"github.com/dashotv/tvdb/openapi/models/shared"
)

// GetUserInfo200ApplicationJSON - response
type GetUserInfo200ApplicationJSON struct {
	Data   []shared.UserInfo `json:"data,omitempty"`
	Status *string           `json:"status,omitempty"`
}

func (o *GetUserInfo200ApplicationJSON) GetData() []shared.UserInfo {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetUserInfo200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetUserInfoResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	GetUserInfo200ApplicationJSONObject *GetUserInfo200ApplicationJSON
}

func (o *GetUserInfoResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetUserInfoResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetUserInfoResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetUserInfoResponse) GetGetUserInfo200ApplicationJSONObject() *GetUserInfo200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetUserInfo200ApplicationJSONObject
}