// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"

	"github.com/dashotv/tvdb/openapi/models/shared"
)

type GetSeriesArtworksRequest struct {
	// id
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
	// lang
	Lang *string `queryParam:"style=form,explode=true,name=lang"`
	// type
	Type *int64 `queryParam:"style=form,explode=true,name=type"`
}

func (o *GetSeriesArtworksRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *GetSeriesArtworksRequest) GetLang() *string {
	if o == nil {
		return nil
	}
	return o.Lang
}

func (o *GetSeriesArtworksRequest) GetType() *int64 {
	if o == nil {
		return nil
	}
	return o.Type
}

// GetSeriesArtworks200ApplicationJSON - response
type GetSeriesArtworks200ApplicationJSON struct {
	// The extended record for a series. All series airs time like firstAired, lastAired, nextAired, etc. are in US EST for US series, and for all non-US series, the time of the showâ€™s country capital or most populous city. For streaming services, is the official release time. See https://support.thetvdb.com/kb/faq.php?id=29.
	Data   *shared.SeriesExtendedRecord `json:"data,omitempty"`
	Status *string                      `json:"status,omitempty"`
}

func (o *GetSeriesArtworks200ApplicationJSON) GetData() *shared.SeriesExtendedRecord {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetSeriesArtworks200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetSeriesArtworksResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	GetSeriesArtworks200ApplicationJSONObject *GetSeriesArtworks200ApplicationJSON
}

func (o *GetSeriesArtworksResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSeriesArtworksResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSeriesArtworksResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSeriesArtworksResponse) GetGetSeriesArtworks200ApplicationJSONObject() *GetSeriesArtworks200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetSeriesArtworks200ApplicationJSONObject
}
