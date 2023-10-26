// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"

	"github.com/dashotv/tvdb/openapi/models/shared"
)

type GetSeriesTranslationRequest struct {
	// id
	ID float64 `pathParam:"style=simple,explode=false,name=id"`
	// language
	Language string `pathParam:"style=simple,explode=false,name=language"`
}

func (o *GetSeriesTranslationRequest) GetID() float64 {
	if o == nil {
		return 0.0
	}
	return o.ID
}

func (o *GetSeriesTranslationRequest) GetLanguage() string {
	if o == nil {
		return ""
	}
	return o.Language
}

// GetSeriesTranslation200ApplicationJSON - response
type GetSeriesTranslation200ApplicationJSON struct {
	// translation record
	Data   *shared.Translation `json:"data,omitempty"`
	Status *string             `json:"status,omitempty"`
}

func (o *GetSeriesTranslation200ApplicationJSON) GetData() *shared.Translation {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetSeriesTranslation200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetSeriesTranslationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	GetSeriesTranslation200ApplicationJSONObject *GetSeriesTranslation200ApplicationJSON
}

func (o *GetSeriesTranslationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSeriesTranslationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSeriesTranslationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSeriesTranslationResponse) GetGetSeriesTranslation200ApplicationJSONObject() *GetSeriesTranslation200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetSeriesTranslation200ApplicationJSONObject
}