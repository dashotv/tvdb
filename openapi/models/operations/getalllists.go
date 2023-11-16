// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"

	"github.com/dashotv/tvdb/openapi/models/shared"
)

type GetAllListsRequest struct {
	// page number
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
}

func (o *GetAllListsRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

// GetAllListsResponseBody - response
type GetAllListsResponseBody struct {
	Data []shared.ListBaseRecord `json:"data,omitempty"`
	// Links for next, previous and current record
	Links  *shared.Links `json:"links,omitempty"`
	Status *string       `json:"status,omitempty"`
}

func (o *GetAllListsResponseBody) GetData() []shared.ListBaseRecord {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetAllListsResponseBody) GetLinks() *shared.Links {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *GetAllListsResponseBody) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetAllListsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	Object *GetAllListsResponseBody
}

func (o *GetAllListsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAllListsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAllListsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAllListsResponse) GetObject() *GetAllListsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
