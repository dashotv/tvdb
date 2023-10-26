// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tvdb/openapi/models/shared"
	"github.com/dashotv/tvdb/openapi/utils"
	"net/http"
)

type GetSeriesSeasonEpisodesTranslatedRequest struct {
	// id
	ID   float64 `pathParam:"style=simple,explode=false,name=id"`
	Lang string  `pathParam:"style=simple,explode=false,name=lang"`
	Page int64   `default:"0" queryParam:"style=form,explode=true,name=page"`
	// season-type
	SeasonType string `pathParam:"style=simple,explode=false,name=season-type"`
}

func (g GetSeriesSeasonEpisodesTranslatedRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetSeriesSeasonEpisodesTranslatedRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetSeriesSeasonEpisodesTranslatedRequest) GetID() float64 {
	if o == nil {
		return 0.0
	}
	return o.ID
}

func (o *GetSeriesSeasonEpisodesTranslatedRequest) GetLang() string {
	if o == nil {
		return ""
	}
	return o.Lang
}

func (o *GetSeriesSeasonEpisodesTranslatedRequest) GetPage() int64 {
	if o == nil {
		return 0
	}
	return o.Page
}

func (o *GetSeriesSeasonEpisodesTranslatedRequest) GetSeasonType() string {
	if o == nil {
		return ""
	}
	return o.SeasonType
}

type GetSeriesSeasonEpisodesTranslated200ApplicationJSONData struct {
	// The base record for a series. All series airs time like firstAired, lastAired, nextAired, etc. are in US EST for US series, and for all non-US series, the time of the showâ€™s country capital or most populous city. For streaming services, is the official release time. See https://support.thetvdb.com/kb/faq.php?id=29.
	Series *shared.SeriesBaseRecord `json:"series,omitempty"`
}

func (o *GetSeriesSeasonEpisodesTranslated200ApplicationJSONData) GetSeries() *shared.SeriesBaseRecord {
	if o == nil {
		return nil
	}
	return o.Series
}

// GetSeriesSeasonEpisodesTranslated200ApplicationJSON - response
type GetSeriesSeasonEpisodesTranslated200ApplicationJSON struct {
	Data   *GetSeriesSeasonEpisodesTranslated200ApplicationJSONData `json:"data,omitempty"`
	Status *string                                                  `json:"status,omitempty"`
}

func (o *GetSeriesSeasonEpisodesTranslated200ApplicationJSON) GetData() *GetSeriesSeasonEpisodesTranslated200ApplicationJSONData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetSeriesSeasonEpisodesTranslated200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetSeriesSeasonEpisodesTranslatedResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	GetSeriesSeasonEpisodesTranslated200ApplicationJSONObject *GetSeriesSeasonEpisodesTranslated200ApplicationJSON
}

func (o *GetSeriesSeasonEpisodesTranslatedResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSeriesSeasonEpisodesTranslatedResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSeriesSeasonEpisodesTranslatedResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSeriesSeasonEpisodesTranslatedResponse) GetGetSeriesSeasonEpisodesTranslated200ApplicationJSONObject() *GetSeriesSeasonEpisodesTranslated200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetSeriesSeasonEpisodesTranslated200ApplicationJSONObject
}
