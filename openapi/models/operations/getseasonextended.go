// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"

	"github.com/dashotv/tvdb/openapi/models/shared"
)

type GetSeasonExtendedRequest struct {
	// id
	ID float64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetSeasonExtendedRequest) GetID() float64 {
	if o == nil {
		return 0.0
	}
	return o.ID
}

// GetSeasonExtendedResponseBody - response
type GetSeasonExtendedResponseBody struct {
	// extended season record
	Data   *shared.SeasonExtendedRecord `json:"data,omitempty"`
	Status *string                      `json:"status,omitempty"`
}

func (o *GetSeasonExtendedResponseBody) GetData() *shared.SeasonExtendedRecord {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetSeasonExtendedResponseBody) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetSeasonExtendedResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	Object *GetSeasonExtendedResponseBody
}

func (o *GetSeasonExtendedResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSeasonExtendedResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSeasonExtendedResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSeasonExtendedResponse) GetObject() *GetSeasonExtendedResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
