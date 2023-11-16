// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"

	"github.com/dashotv/tvdb/openapi/models/shared"
)

// GetAllSeriesStatusesResponseBody - response
type GetAllSeriesStatusesResponseBody struct {
	Data   []shared.Status `json:"data,omitempty"`
	Status *string         `json:"status,omitempty"`
}

func (o *GetAllSeriesStatusesResponseBody) GetData() []shared.Status {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetAllSeriesStatusesResponseBody) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetAllSeriesStatusesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	Object *GetAllSeriesStatusesResponseBody
}

func (o *GetAllSeriesStatusesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAllSeriesStatusesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAllSeriesStatusesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAllSeriesStatusesResponse) GetObject() *GetAllSeriesStatusesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
