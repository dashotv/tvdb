// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"

	"github.com/dashotv/tvdb/openapi/models/shared"
)

type GetSearchResultsByRemoteIDRequest struct {
	// Search for a specific remote id.  Allows searching for an IMDB or EIDR id, for example.
	RemoteID string `pathParam:"style=simple,explode=false,name=remoteId"`
}

func (o *GetSearchResultsByRemoteIDRequest) GetRemoteID() string {
	if o == nil {
		return ""
	}
	return o.RemoteID
}

// GetSearchResultsByRemoteIDResponseBody - response
type GetSearchResultsByRemoteIDResponseBody struct {
	Data   []shared.SearchByRemoteIDResult `json:"data,omitempty"`
	Status *string                         `json:"status,omitempty"`
}

func (o *GetSearchResultsByRemoteIDResponseBody) GetData() []shared.SearchByRemoteIDResult {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetSearchResultsByRemoteIDResponseBody) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetSearchResultsByRemoteIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	Object *GetSearchResultsByRemoteIDResponseBody
}

func (o *GetSearchResultsByRemoteIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSearchResultsByRemoteIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSearchResultsByRemoteIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSearchResultsByRemoteIDResponse) GetObject() *GetSearchResultsByRemoteIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
