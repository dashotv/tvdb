// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"

	"github.com/dashotv/tvdb/openapi/models/shared"
)

type GetSeasonTranslationRequest struct {
	// id
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
	// language
	Language string `pathParam:"style=simple,explode=false,name=language"`
}

func (o *GetSeasonTranslationRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *GetSeasonTranslationRequest) GetLanguage() string {
	if o == nil {
		return ""
	}
	return o.Language
}

// GetSeasonTranslation200ApplicationJSON - response
type GetSeasonTranslation200ApplicationJSON struct {
	// translation record
	Data   *shared.Translation `json:"data,omitempty"`
	Status *string             `json:"status,omitempty"`
}

func (o *GetSeasonTranslation200ApplicationJSON) GetData() *shared.Translation {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetSeasonTranslation200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetSeasonTranslationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	GetSeasonTranslation200ApplicationJSONObject *GetSeasonTranslation200ApplicationJSON
}

func (o *GetSeasonTranslationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSeasonTranslationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSeasonTranslationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSeasonTranslationResponse) GetGetSeasonTranslation200ApplicationJSONObject() *GetSeasonTranslation200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetSeasonTranslation200ApplicationJSONObject
}
