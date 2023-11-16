// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dashotv/tvdb/openapi/models/shared"
)

type Action string

const (
	ActionDelete Action = "delete"
	ActionUpdate Action = "update"
)

func (e Action) ToPointer() *Action {
	return &e
}

func (e *Action) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "delete":
		fallthrough
	case "update":
		*e = Action(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Action: %v", v)
	}
}

type Type string

const (
	TypeArtwork              Type = "artwork"
	TypeAwardNominees        Type = "award_nominees"
	TypeCompanies            Type = "companies"
	TypeEpisodes             Type = "episodes"
	TypeLists                Type = "lists"
	TypePeople               Type = "people"
	TypeSeasons              Type = "seasons"
	TypeSeries               Type = "series"
	TypeSeriespeople         Type = "seriespeople"
	TypeArtworktypes         Type = "artworktypes"
	TypeAwardCategories      Type = "award_categories"
	TypeAwards               Type = "awards"
	TypeCompanyTypes         Type = "company_types"
	TypeContentRatings       Type = "content_ratings"
	TypeCountries            Type = "countries"
	TypeEntityTypes          Type = "entity_types"
	TypeGenres               Type = "genres"
	TypeLanguages            Type = "languages"
	TypeMovies               Type = "movies"
	TypeMovieGenres          Type = "movie_genres"
	TypeMovieStatus          Type = "movie_status"
	TypePeopletypes          Type = "peopletypes"
	TypeSeasontypes          Type = "seasontypes"
	TypeSourcetypes          Type = "sourcetypes"
	TypeTagOptions           Type = "tag_options"
	TypeTags                 Type = "tags"
	TypeTranslatedcharacters Type = "translatedcharacters"
	TypeTranslatedcompanies  Type = "translatedcompanies"
	TypeTranslatedepisodes   Type = "translatedepisodes"
	TypeTranslatedlists      Type = "translatedlists"
	TypeTranslatedmovies     Type = "translatedmovies"
	TypeTranslatedpeople     Type = "translatedpeople"
	TypeTranslatedseasons    Type = "translatedseasons"
	TypeTranslatedserierk    Type = "translatedserierk"
)

func (e Type) ToPointer() *Type {
	return &e
}

func (e *Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "artwork":
		fallthrough
	case "award_nominees":
		fallthrough
	case "companies":
		fallthrough
	case "episodes":
		fallthrough
	case "lists":
		fallthrough
	case "people":
		fallthrough
	case "seasons":
		fallthrough
	case "series":
		fallthrough
	case "seriespeople":
		fallthrough
	case "artworktypes":
		fallthrough
	case "award_categories":
		fallthrough
	case "awards":
		fallthrough
	case "company_types":
		fallthrough
	case "content_ratings":
		fallthrough
	case "countries":
		fallthrough
	case "entity_types":
		fallthrough
	case "genres":
		fallthrough
	case "languages":
		fallthrough
	case "movies":
		fallthrough
	case "movie_genres":
		fallthrough
	case "movie_status":
		fallthrough
	case "peopletypes":
		fallthrough
	case "seasontypes":
		fallthrough
	case "sourcetypes":
		fallthrough
	case "tag_options":
		fallthrough
	case "tags":
		fallthrough
	case "translatedcharacters":
		fallthrough
	case "translatedcompanies":
		fallthrough
	case "translatedepisodes":
		fallthrough
	case "translatedlists":
		fallthrough
	case "translatedmovies":
		fallthrough
	case "translatedpeople":
		fallthrough
	case "translatedseasons":
		fallthrough
	case "translatedserierk":
		*e = Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Type: %v", v)
	}
}

type UpdatesRequest struct {
	Since  int64   `queryParam:"style=form,explode=true,name=since"`
	Action *Action `queryParam:"style=form,explode=true,name=action"`
	// name
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
	Type *Type  `queryParam:"style=form,explode=true,name=type"`
}

func (o *UpdatesRequest) GetSince() int64 {
	if o == nil {
		return 0
	}
	return o.Since
}

func (o *UpdatesRequest) GetAction() *Action {
	if o == nil {
		return nil
	}
	return o.Action
}

func (o *UpdatesRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *UpdatesRequest) GetType() *Type {
	if o == nil {
		return nil
	}
	return o.Type
}

// UpdatesResponseBody - response
type UpdatesResponseBody struct {
	Data []shared.EntityUpdate `json:"data,omitempty"`
	// Links for next, previous and current record
	Links  *shared.Links `json:"links,omitempty"`
	Status *string       `json:"status,omitempty"`
}

func (o *UpdatesResponseBody) GetData() []shared.EntityUpdate {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *UpdatesResponseBody) GetLinks() *shared.Links {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *UpdatesResponseBody) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type UpdatesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// response
	Object *UpdatesResponseBody
}

func (o *UpdatesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdatesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdatesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdatesResponse) GetObject() *UpdatesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
