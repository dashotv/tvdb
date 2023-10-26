// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"

	"github.com/dashotv/tvdb/openapi/models/shared"
)

type GetAllEpisodesRequest struct {
	// page number
	Page *float64 `queryParam:"style=form,explode=true,name=page"`
}

func (o *GetAllEpisodesRequest) GetPage() *float64 {
	if o == nil {
		return nil
	}
	return o.Page
}

// GetAllEpisodes200ApplicationJSON - response
type GetAllEpisodes200ApplicationJSON struct {
	Data []shared.EpisodeBaseRecord `json:"data,omitempty"`
	// Links for next, previous and current record
	Links  *shared.Links `json:"links,omitempty"`
	Status *string       `json:"status,omitempty"`
}

func (o *GetAllEpisodes200ApplicationJSON) GetData() []shared.EpisodeBaseRecord {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetAllEpisodes200ApplicationJSON) GetLinks() *shared.Links {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *GetAllEpisodes200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetAllEpisodesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	GetAllEpisodes200ApplicationJSONObject *GetAllEpisodes200ApplicationJSON
}

func (o *GetAllEpisodesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAllEpisodesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAllEpisodesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAllEpisodesResponse) GetGetAllEpisodes200ApplicationJSONObject() *GetAllEpisodes200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetAllEpisodes200ApplicationJSONObject
}
