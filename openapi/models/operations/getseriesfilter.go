// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dashotv/tvdb/openapi/models/shared"
)

// GetSeriesFilterSort - sort by results
type GetSeriesFilterSort string

const (
	GetSeriesFilterSortScore      GetSeriesFilterSort = "score"
	GetSeriesFilterSortFirstAired GetSeriesFilterSort = "firstAired"
	GetSeriesFilterSortLastAired  GetSeriesFilterSort = "lastAired"
	GetSeriesFilterSortName       GetSeriesFilterSort = "name"
)

func (e GetSeriesFilterSort) ToPointer() *GetSeriesFilterSort {
	return &e
}

func (e *GetSeriesFilterSort) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "score":
		fallthrough
	case "firstAired":
		fallthrough
	case "lastAired":
		fallthrough
	case "name":
		*e = GetSeriesFilterSort(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetSeriesFilterSort: %v", v)
	}
}

// GetSeriesFilterSortType - sort type ascending or descending
type GetSeriesFilterSortType string

const (
	GetSeriesFilterSortTypeAsc  GetSeriesFilterSortType = "asc"
	GetSeriesFilterSortTypeDesc GetSeriesFilterSortType = "desc"
)

func (e GetSeriesFilterSortType) ToPointer() *GetSeriesFilterSortType {
	return &e
}

func (e *GetSeriesFilterSortType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetSeriesFilterSortType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetSeriesFilterSortType: %v", v)
	}
}

type GetSeriesFilterRequest struct {
	// production company
	Company *float64 `queryParam:"style=form,explode=true,name=company"`
	// content rating id base on a country
	ContentRating *float64 `queryParam:"style=form,explode=true,name=contentRating"`
	// country of origin
	Country string `queryParam:"style=form,explode=true,name=country"`
	// Genre id. This id can be found using **/genres** endpoint.
	Genre *float64 `queryParam:"style=form,explode=true,name=genre"`
	// original language
	Lang string `queryParam:"style=form,explode=true,name=lang"`
	// sort by results
	Sort *GetSeriesFilterSort `queryParam:"style=form,explode=true,name=sort"`
	// sort type ascending or descending
	SortType *GetSeriesFilterSortType `queryParam:"style=form,explode=true,name=sortType"`
	// status
	Status *float64 `queryParam:"style=form,explode=true,name=status"`
	// release year
	Year *float64 `queryParam:"style=form,explode=true,name=year"`
}

func (o *GetSeriesFilterRequest) GetCompany() *float64 {
	if o == nil {
		return nil
	}
	return o.Company
}

func (o *GetSeriesFilterRequest) GetContentRating() *float64 {
	if o == nil {
		return nil
	}
	return o.ContentRating
}

func (o *GetSeriesFilterRequest) GetCountry() string {
	if o == nil {
		return ""
	}
	return o.Country
}

func (o *GetSeriesFilterRequest) GetGenre() *float64 {
	if o == nil {
		return nil
	}
	return o.Genre
}

func (o *GetSeriesFilterRequest) GetLang() string {
	if o == nil {
		return ""
	}
	return o.Lang
}

func (o *GetSeriesFilterRequest) GetSort() *GetSeriesFilterSort {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *GetSeriesFilterRequest) GetSortType() *GetSeriesFilterSortType {
	if o == nil {
		return nil
	}
	return o.SortType
}

func (o *GetSeriesFilterRequest) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *GetSeriesFilterRequest) GetYear() *float64 {
	if o == nil {
		return nil
	}
	return o.Year
}

// GetSeriesFilter200ApplicationJSON - response
type GetSeriesFilter200ApplicationJSON struct {
	Data []shared.SeriesBaseRecord `json:"data,omitempty"`
}

func (o *GetSeriesFilter200ApplicationJSON) GetData() []shared.SeriesBaseRecord {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetSeriesFilterResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	GetSeriesFilter200ApplicationJSONObject *GetSeriesFilter200ApplicationJSON
}

func (o *GetSeriesFilterResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSeriesFilterResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSeriesFilterResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSeriesFilterResponse) GetGetSeriesFilter200ApplicationJSONObject() *GetSeriesFilter200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetSeriesFilter200ApplicationJSONObject
}