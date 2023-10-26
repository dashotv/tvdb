// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"

	"github.com/dashotv/tvdb/openapi/models/shared"
)

// GetAllLanguages200ApplicationJSON - response
type GetAllLanguages200ApplicationJSON struct {
	Data   []shared.Language `json:"data,omitempty"`
	Status *string           `json:"status,omitempty"`
}

func (o *GetAllLanguages200ApplicationJSON) GetData() []shared.Language {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetAllLanguages200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetAllLanguagesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	GetAllLanguages200ApplicationJSONObject *GetAllLanguages200ApplicationJSON
}

func (o *GetAllLanguagesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAllLanguagesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAllLanguagesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAllLanguagesResponse) GetGetAllLanguages200ApplicationJSONObject() *GetAllLanguages200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetAllLanguages200ApplicationJSONObject
}
