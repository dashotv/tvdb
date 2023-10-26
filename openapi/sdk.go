// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package openapi

import (
	"context"
	"fmt"
	"github.com/dashotv/tvdb/openapi/models/shared"
	"github.com/dashotv/tvdb/openapi/utils"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	"https://api4.thetvdb.com/v4",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient     HTTPClient
	SecurityClient    HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *utils.RetryConfig
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// SDK - TVDB API V4: Documentation of [TheTVDB](https://thetvdb.com/) API V4. All related information is linked from our [Github repo](https://github.com/thetvdb/v4-api). You might also want to use our [Postman collection] (https://www.getpostman.com/collections/7a9397ce69ff246f74d0)
// ## Authentication
// 1. Use the /login endpoint and provide your API key as "apikey". If you have a user-supported key, also provide your subscriber PIN as "pin". Otherwise completely remove "pin" from your call.
// 2. Executing this call will provide you with a bearer token, which is valid for 1 month.
// 3. Provide your bearer token for subsequent API calls by clicking Authorize below or including in the header of all direct API calls: `Authorization: Bearer [your-token]`
//
// ## Notes
// 1. "score" is a field across almost all entities.  We generate scores for different types of entities in various ways, so no assumptions should be made about the meaning of this value.  It is simply used to hint at relative popularity for sorting purposes.
type SDK struct {
	Artwork          *artwork
	ArtworkStatuses  *artworkStatuses
	ArtworkTypes     *artworkTypes
	AwardCategories  *awardCategories
	Awards           *awards
	Characters       *characters
	Companies        *companies
	ContentRatings   *contentRatings
	Countries        *countries
	EntityTypes      *entityTypes
	Episodes         *episodes
	Favorites        *favorites
	Genders          *genders
	Genres           *genres
	InspirationTypes *inspirationTypes
	Languages        *languages
	Lists            *lists
	Login            *login
	MovieStatuses    *movieStatuses
	Movies           *movies
	People           *people
	PeopleTypes      *peopleTypes
	Search           *search
	Seasons          *seasons
	Series           *series
	SeriesStatuses   *seriesStatuses
	SourceTypes      *sourceTypes
	Updates          *updates
	UserInfo         *userInfo

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*SDK)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *SDK) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *SDK) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

func withSecurity(security interface{}) func(context.Context) (interface{}, error) {
	return func(context.Context) (interface{}, error) {
		return &security, nil
	}
}

// WithSecurity configures the SDK to use the provided security details

func WithSecurity(bearerAuth string) SDKOption {
	return func(sdk *SDK) {
		security := shared.Security{BearerAuth: bearerAuth}
		sdk.sdkConfiguration.Security = withSecurity(&security)
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *SDK {
	sdk := &SDK{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "4.7.8",
			SDKVersion:        "v0.1.0-6-gc93f04c-dirty",
			GenVersion:        "2.172.0",
			UserAgent:         "speakeasy-sdk/go v0.1.0-6-gc93f04c-dirty 2.172.0 4.7.8 github.com/dashotv/tvdb/openapi",
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		if sdk.sdkConfiguration.Security != nil {
			sdk.sdkConfiguration.SecurityClient = utils.ConfigureSecurityClient(sdk.sdkConfiguration.DefaultClient, sdk.sdkConfiguration.Security)
		} else {
			sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
		}
	}

	sdk.Artwork = newArtwork(sdk.sdkConfiguration)

	sdk.ArtworkStatuses = newArtworkStatuses(sdk.sdkConfiguration)

	sdk.ArtworkTypes = newArtworkTypes(sdk.sdkConfiguration)

	sdk.AwardCategories = newAwardCategories(sdk.sdkConfiguration)

	sdk.Awards = newAwards(sdk.sdkConfiguration)

	sdk.Characters = newCharacters(sdk.sdkConfiguration)

	sdk.Companies = newCompanies(sdk.sdkConfiguration)

	sdk.ContentRatings = newContentRatings(sdk.sdkConfiguration)

	sdk.Countries = newCountries(sdk.sdkConfiguration)

	sdk.EntityTypes = newEntityTypes(sdk.sdkConfiguration)

	sdk.Episodes = newEpisodes(sdk.sdkConfiguration)

	sdk.Favorites = newFavorites(sdk.sdkConfiguration)

	sdk.Genders = newGenders(sdk.sdkConfiguration)

	sdk.Genres = newGenres(sdk.sdkConfiguration)

	sdk.InspirationTypes = newInspirationTypes(sdk.sdkConfiguration)

	sdk.Languages = newLanguages(sdk.sdkConfiguration)

	sdk.Lists = newLists(sdk.sdkConfiguration)

	sdk.Login = newLogin(sdk.sdkConfiguration)

	sdk.MovieStatuses = newMovieStatuses(sdk.sdkConfiguration)

	sdk.Movies = newMovies(sdk.sdkConfiguration)

	sdk.People = newPeople(sdk.sdkConfiguration)

	sdk.PeopleTypes = newPeopleTypes(sdk.sdkConfiguration)

	sdk.Search = newSearch(sdk.sdkConfiguration)

	sdk.Seasons = newSeasons(sdk.sdkConfiguration)

	sdk.Series = newSeries(sdk.sdkConfiguration)

	sdk.SeriesStatuses = newSeriesStatuses(sdk.sdkConfiguration)

	sdk.SourceTypes = newSourceTypes(sdk.sdkConfiguration)

	sdk.Updates = newUpdates(sdk.sdkConfiguration)

	sdk.UserInfo = newUserInfo(sdk.sdkConfiguration)

	return sdk
}
