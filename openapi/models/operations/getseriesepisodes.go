// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"

	"github.com/dashotv/tvdb/openapi/models/shared"
	"github.com/dashotv/tvdb/openapi/utils"
)

type GetSeriesEpisodesRequest struct {
	// airDate of the episode, format is yyyy-mm-dd
	AirDate       *string `queryParam:"style=form,explode=true,name=airDate"`
	EpisodeNumber *int64  `default:"0" queryParam:"style=form,explode=true,name=episodeNumber"`
	// id
	ID     float64 `pathParam:"style=simple,explode=false,name=id"`
	Page   int64   `default:"0" queryParam:"style=form,explode=true,name=page"`
	Season *int64  `default:"0" queryParam:"style=form,explode=true,name=season"`
	// season-type
	SeasonType string `pathParam:"style=simple,explode=false,name=season-type"`
}

func (g GetSeriesEpisodesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetSeriesEpisodesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetSeriesEpisodesRequest) GetAirDate() *string {
	if o == nil {
		return nil
	}
	return o.AirDate
}

func (o *GetSeriesEpisodesRequest) GetEpisodeNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.EpisodeNumber
}

func (o *GetSeriesEpisodesRequest) GetID() float64 {
	if o == nil {
		return 0.0
	}
	return o.ID
}

func (o *GetSeriesEpisodesRequest) GetPage() int64 {
	if o == nil {
		return 0
	}
	return o.Page
}

func (o *GetSeriesEpisodesRequest) GetSeason() *int64 {
	if o == nil {
		return nil
	}
	return o.Season
}

func (o *GetSeriesEpisodesRequest) GetSeasonType() string {
	if o == nil {
		return ""
	}
	return o.SeasonType
}

type GetSeriesEpisodes200ApplicationJSONData struct {
	Episodes []shared.EpisodeBaseRecord `json:"episodes,omitempty"`
	// The base record for a series. All series airs time like firstAired, lastAired, nextAired, etc. are in US EST for US series, and for all non-US series, the time of the showâ€™s country capital or most populous city. For streaming services, is the official release time. See https://support.thetvdb.com/kb/faq.php?id=29.
	Series *shared.SeriesBaseRecord `json:"series,omitempty"`
}

func (o *GetSeriesEpisodes200ApplicationJSONData) GetEpisodes() []shared.EpisodeBaseRecord {
	if o == nil {
		return nil
	}
	return o.Episodes
}

func (o *GetSeriesEpisodes200ApplicationJSONData) GetSeries() *shared.SeriesBaseRecord {
	if o == nil {
		return nil
	}
	return o.Series
}

// GetSeriesEpisodes200ApplicationJSON - response
type GetSeriesEpisodes200ApplicationJSON struct {
	Data   *GetSeriesEpisodes200ApplicationJSONData `json:"data,omitempty"`
	Status *string                                  `json:"status,omitempty"`
}

func (o *GetSeriesEpisodes200ApplicationJSON) GetData() *GetSeriesEpisodes200ApplicationJSONData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetSeriesEpisodes200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetSeriesEpisodesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	GetSeriesEpisodes200ApplicationJSONObject *GetSeriesEpisodes200ApplicationJSON
}

func (o *GetSeriesEpisodesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSeriesEpisodesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSeriesEpisodesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSeriesEpisodesResponse) GetGetSeriesEpisodes200ApplicationJSONObject() *GetSeriesEpisodes200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetSeriesEpisodes200ApplicationJSONObject
}