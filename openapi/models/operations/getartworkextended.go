// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"

	"github.com/dashotv/tvdb/openapi/models/shared"
)

type GetArtworkExtendedRequest struct {
	// id
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetArtworkExtendedRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

// GetArtworkExtended200ApplicationJSON - response
type GetArtworkExtended200ApplicationJSON struct {
	// extended artwork record
	Data   *shared.ArtworkExtendedRecord `json:"data,omitempty"`
	Status *string                       `json:"status,omitempty"`
}

func (o *GetArtworkExtended200ApplicationJSON) GetData() *shared.ArtworkExtendedRecord {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetArtworkExtended200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetArtworkExtendedResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	GetArtworkExtended200ApplicationJSONObject *GetArtworkExtended200ApplicationJSON
}

func (o *GetArtworkExtendedResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetArtworkExtendedResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetArtworkExtendedResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetArtworkExtendedResponse) GetGetArtworkExtended200ApplicationJSONObject() *GetArtworkExtended200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetArtworkExtended200ApplicationJSONObject
}
