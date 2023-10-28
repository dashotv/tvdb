// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dashotv/tvdb/openapi/models/shared"
)

// GetEpisodeExtendedMeta - meta
type GetEpisodeExtendedMeta string

const (
	GetEpisodeExtendedMetaTranslations GetEpisodeExtendedMeta = "translations"
)

func (e GetEpisodeExtendedMeta) ToPointer() *GetEpisodeExtendedMeta {
	return &e
}

func (e *GetEpisodeExtendedMeta) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "translations":
		*e = GetEpisodeExtendedMeta(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetEpisodeExtendedMeta: %v", v)
	}
}

type GetEpisodeExtendedRequest struct {
	// id
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
	// meta
	Meta *GetEpisodeExtendedMeta `queryParam:"style=form,explode=true,name=meta"`
}

func (o *GetEpisodeExtendedRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *GetEpisodeExtendedRequest) GetMeta() *GetEpisodeExtendedMeta {
	if o == nil {
		return nil
	}
	return o.Meta
}

// GetEpisodeExtended200ApplicationJSON - response
type GetEpisodeExtended200ApplicationJSON struct {
	// extended episode record
	Data   *shared.EpisodeExtendedRecord `json:"data,omitempty"`
	Status *string                       `json:"status,omitempty"`
}

func (o *GetEpisodeExtended200ApplicationJSON) GetData() *shared.EpisodeExtendedRecord {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetEpisodeExtended200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetEpisodeExtendedResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	GetEpisodeExtended200ApplicationJSONObject *GetEpisodeExtended200ApplicationJSON
}

func (o *GetEpisodeExtendedResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetEpisodeExtendedResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetEpisodeExtendedResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetEpisodeExtendedResponse) GetGetEpisodeExtended200ApplicationJSONObject() *GetEpisodeExtended200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetEpisodeExtended200ApplicationJSONObject
}
