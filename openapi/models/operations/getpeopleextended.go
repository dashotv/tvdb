// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dashotv/tvdb/openapi/models/shared"
)

// GetPeopleExtendedQueryParamMeta - meta
type GetPeopleExtendedQueryParamMeta string

const (
	GetPeopleExtendedQueryParamMetaTranslations GetPeopleExtendedQueryParamMeta = "translations"
)

func (e GetPeopleExtendedQueryParamMeta) ToPointer() *GetPeopleExtendedQueryParamMeta {
	return &e
}
func (e *GetPeopleExtendedQueryParamMeta) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "translations":
		*e = GetPeopleExtendedQueryParamMeta(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetPeopleExtendedQueryParamMeta: %v", v)
	}
}

type GetPeopleExtendedRequest struct {
	// id
	ID float64 `pathParam:"style=simple,explode=false,name=id"`
	// meta
	Meta *GetPeopleExtendedQueryParamMeta `queryParam:"style=form,explode=true,name=meta"`
}

func (o *GetPeopleExtendedRequest) GetID() float64 {
	if o == nil {
		return 0.0
	}
	return o.ID
}

func (o *GetPeopleExtendedRequest) GetMeta() *GetPeopleExtendedQueryParamMeta {
	if o == nil {
		return nil
	}
	return o.Meta
}

// GetPeopleExtendedResponseBody - response
type GetPeopleExtendedResponseBody struct {
	// extended people record
	Data   *shared.PeopleExtendedRecord `json:"data,omitempty"`
	Status *string                      `json:"status,omitempty"`
}

func (o *GetPeopleExtendedResponseBody) GetData() *shared.PeopleExtendedRecord {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetPeopleExtendedResponseBody) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetPeopleExtendedResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	Object *GetPeopleExtendedResponseBody
}

func (o *GetPeopleExtendedResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPeopleExtendedResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPeopleExtendedResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetPeopleExtendedResponse) GetObject() *GetPeopleExtendedResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
