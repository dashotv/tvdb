// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"

	"github.com/dashotv/tvdb/openapi/models/shared"
)

type GetAwardCategoryRequest struct {
	// id
	ID float64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetAwardCategoryRequest) GetID() float64 {
	if o == nil {
		return 0.0
	}
	return o.ID
}

// GetAwardCategoryResponseBody - response
type GetAwardCategoryResponseBody struct {
	// base award category record
	Data   *shared.AwardCategoryBaseRecord `json:"data,omitempty"`
	Status *string                         `json:"status,omitempty"`
}

func (o *GetAwardCategoryResponseBody) GetData() *shared.AwardCategoryBaseRecord {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetAwardCategoryResponseBody) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetAwardCategoryResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	Object *GetAwardCategoryResponseBody
}

func (o *GetAwardCategoryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAwardCategoryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAwardCategoryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAwardCategoryResponse) GetObject() *GetAwardCategoryResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
