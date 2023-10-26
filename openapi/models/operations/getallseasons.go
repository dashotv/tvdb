// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"

	"github.com/dashotv/tvdb/openapi/models/shared"
)

type GetAllSeasonsRequest struct {
	// page number
	Page *float64 `queryParam:"style=form,explode=true,name=page"`
}

func (o *GetAllSeasonsRequest) GetPage() *float64 {
	if o == nil {
		return nil
	}
	return o.Page
}

// GetAllSeasons200ApplicationJSON - response
type GetAllSeasons200ApplicationJSON struct {
	Data   []shared.SeasonBaseRecord `json:"data,omitempty"`
	Status *string                   `json:"status,omitempty"`
}

func (o *GetAllSeasons200ApplicationJSON) GetData() []shared.SeasonBaseRecord {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetAllSeasons200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetAllSeasonsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	GetAllSeasons200ApplicationJSONObject *GetAllSeasons200ApplicationJSON
}

func (o *GetAllSeasonsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAllSeasonsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAllSeasonsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAllSeasonsResponse) GetGetAllSeasons200ApplicationJSONObject() *GetAllSeasons200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetAllSeasons200ApplicationJSONObject
}
