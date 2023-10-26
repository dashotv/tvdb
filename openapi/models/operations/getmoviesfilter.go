// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dashotv/tvdb/openapi/models/shared"
)

// GetMoviesFilterSort - sort by results
type GetMoviesFilterSort string

const (
	GetMoviesFilterSortScore      GetMoviesFilterSort = "score"
	GetMoviesFilterSortFirstAired GetMoviesFilterSort = "firstAired"
	GetMoviesFilterSortName       GetMoviesFilterSort = "name"
)

func (e GetMoviesFilterSort) ToPointer() *GetMoviesFilterSort {
	return &e
}

func (e *GetMoviesFilterSort) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "score":
		fallthrough
	case "firstAired":
		fallthrough
	case "name":
		*e = GetMoviesFilterSort(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetMoviesFilterSort: %v", v)
	}
}

type GetMoviesFilterRequest struct {
	// production company
	Company *float64 `queryParam:"style=form,explode=true,name=company"`
	// content rating id base on a country
	ContentRating *float64 `queryParam:"style=form,explode=true,name=contentRating"`
	// country of origin
	Country string `queryParam:"style=form,explode=true,name=country"`
	// genre
	Genre *float64 `queryParam:"style=form,explode=true,name=genre"`
	// original language
	Lang string `queryParam:"style=form,explode=true,name=lang"`
	// sort by results
	Sort *GetMoviesFilterSort `queryParam:"style=form,explode=true,name=sort"`
	// status
	Status *float64 `queryParam:"style=form,explode=true,name=status"`
	// release year
	Year *float64 `queryParam:"style=form,explode=true,name=year"`
}

func (o *GetMoviesFilterRequest) GetCompany() *float64 {
	if o == nil {
		return nil
	}
	return o.Company
}

func (o *GetMoviesFilterRequest) GetContentRating() *float64 {
	if o == nil {
		return nil
	}
	return o.ContentRating
}

func (o *GetMoviesFilterRequest) GetCountry() string {
	if o == nil {
		return ""
	}
	return o.Country
}

func (o *GetMoviesFilterRequest) GetGenre() *float64 {
	if o == nil {
		return nil
	}
	return o.Genre
}

func (o *GetMoviesFilterRequest) GetLang() string {
	if o == nil {
		return ""
	}
	return o.Lang
}

func (o *GetMoviesFilterRequest) GetSort() *GetMoviesFilterSort {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *GetMoviesFilterRequest) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *GetMoviesFilterRequest) GetYear() *float64 {
	if o == nil {
		return nil
	}
	return o.Year
}

// GetMoviesFilter200ApplicationJSON - response
type GetMoviesFilter200ApplicationJSON struct {
	Data   []shared.MovieBaseRecord `json:"data,omitempty"`
	Status *string                  `json:"status,omitempty"`
}

func (o *GetMoviesFilter200ApplicationJSON) GetData() []shared.MovieBaseRecord {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetMoviesFilter200ApplicationJSON) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetMoviesFilterResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	GetMoviesFilter200ApplicationJSONObject *GetMoviesFilter200ApplicationJSON
}

func (o *GetMoviesFilterResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMoviesFilterResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMoviesFilterResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetMoviesFilterResponse) GetGetMoviesFilter200ApplicationJSONObject() *GetMoviesFilter200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetMoviesFilter200ApplicationJSONObject
}
