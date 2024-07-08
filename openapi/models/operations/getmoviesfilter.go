// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dashotv/tvdb/openapi/models/shared"
)

// Sort - sort by results
type Sort string

const (
	SortScore      Sort = "score"
	SortFirstAired Sort = "firstAired"
	SortName       Sort = "name"
)

func (e Sort) ToPointer() *Sort {
	return &e
}
func (e *Sort) UnmarshalJSON(data []byte) error {
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
		*e = Sort(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Sort: %v", v)
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
	Sort *Sort `queryParam:"style=form,explode=true,name=sort"`
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

func (o *GetMoviesFilterRequest) GetSort() *Sort {
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

// GetMoviesFilterResponseBody - response
type GetMoviesFilterResponseBody struct {
	Data   []shared.MovieBaseRecord `json:"data,omitempty"`
	Status *string                  `json:"status,omitempty"`
}

func (o *GetMoviesFilterResponseBody) GetData() []shared.MovieBaseRecord {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetMoviesFilterResponseBody) GetStatus() *string {
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
	Object *GetMoviesFilterResponseBody
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

func (o *GetMoviesFilterResponse) GetObject() *GetMoviesFilterResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
