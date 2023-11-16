// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"

	"github.com/dashotv/tvdb/openapi/models/shared"
)

type GetEpisodeBaseRequest struct {
	// id
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetEpisodeBaseRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

// GetEpisodeBaseResponseBody - response
type GetEpisodeBaseResponseBody struct {
	// base episode record
	Data   *shared.EpisodeBaseRecord `json:"data,omitempty"`
	Status *string                   `json:"status,omitempty"`
}

func (o *GetEpisodeBaseResponseBody) GetData() *shared.EpisodeBaseRecord {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetEpisodeBaseResponseBody) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetEpisodeBaseResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	Object *GetEpisodeBaseResponseBody
}

func (o *GetEpisodeBaseResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetEpisodeBaseResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetEpisodeBaseResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetEpisodeBaseResponse) GetObject() *GetEpisodeBaseResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
