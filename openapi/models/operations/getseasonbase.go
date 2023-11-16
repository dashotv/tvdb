// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"

	"github.com/dashotv/tvdb/openapi/models/shared"
)

type GetSeasonBaseRequest struct {
	// id
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetSeasonBaseRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

// GetSeasonBaseResponseBody - response
type GetSeasonBaseResponseBody struct {
	// season genre record
	Data   *shared.SeasonBaseRecord `json:"data,omitempty"`
	Status *string                  `json:"status,omitempty"`
}

func (o *GetSeasonBaseResponseBody) GetData() *shared.SeasonBaseRecord {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetSeasonBaseResponseBody) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetSeasonBaseResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	Object *GetSeasonBaseResponseBody
}

func (o *GetSeasonBaseResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSeasonBaseResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSeasonBaseResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSeasonBaseResponse) GetObject() *GetSeasonBaseResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
