// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dashotv/tvdb/openapi/models/shared"
)

// QueryParamMeta - meta
type QueryParamMeta string

const (
	QueryParamMetaTranslations QueryParamMeta = "translations"
)

func (e QueryParamMeta) ToPointer() *QueryParamMeta {
	return &e
}

func (e *QueryParamMeta) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "translations":
		*e = QueryParamMeta(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueryParamMeta: %v", v)
	}
}

type GetMovieExtendedRequest struct {
	// id
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
	// meta
	Meta *QueryParamMeta `queryParam:"style=form,explode=true,name=meta"`
	// reduce the payload and returns the short version of this record without characters, artworks and trailers.
	Short *bool `queryParam:"style=form,explode=true,name=short"`
}

func (o *GetMovieExtendedRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *GetMovieExtendedRequest) GetMeta() *QueryParamMeta {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *GetMovieExtendedRequest) GetShort() *bool {
	if o == nil {
		return nil
	}
	return o.Short
}

// GetMovieExtendedResponseBody - response
type GetMovieExtendedResponseBody struct {
	// extended movie record
	Data   *shared.MovieExtendedRecord `json:"data,omitempty"`
	Status *string                     `json:"status,omitempty"`
}

func (o *GetMovieExtendedResponseBody) GetData() *shared.MovieExtendedRecord {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetMovieExtendedResponseBody) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetMovieExtendedResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	Object *GetMovieExtendedResponseBody
}

func (o *GetMovieExtendedResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMovieExtendedResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMovieExtendedResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetMovieExtendedResponse) GetObject() *GetMovieExtendedResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
