// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"

	"github.com/dashotv/tvdb/openapi/models/shared"
)

// GetAllPeopleTypesResponseBody - response
type GetAllPeopleTypesResponseBody struct {
	Data   []shared.PeopleType `json:"data,omitempty"`
	Status *string             `json:"status,omitempty"`
}

func (o *GetAllPeopleTypesResponseBody) GetData() []shared.PeopleType {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetAllPeopleTypesResponseBody) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetAllPeopleTypesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	Object *GetAllPeopleTypesResponseBody
}

func (o *GetAllPeopleTypesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAllPeopleTypesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAllPeopleTypesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAllPeopleTypesResponse) GetObject() *GetAllPeopleTypesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
