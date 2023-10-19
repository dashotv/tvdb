package tvdb

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type contextKey string

const (
	contextBearerToken contextKey = "bearerToken"
)

// Tvdb wrapper client
type Tvdb struct {
	URL   string
	Key   string
	token string
	c     *ClientWithResponses
	ctx   context.Context
}

// New creates a new Tvdb wrapper client
//
//	url is the base URL of the TVDB API
//
// Example: https://api4.thetvdb.com/v4
//
// See https://thetvdb.github.io/v4-api/#/Introduction for more information
//
// This wraps the generated client with a few helper functions, it also sets
// up authentication and simplifies API access.
func New(url string) *Tvdb {
	c, _ := newClientWithResponses(url)
	return &Tvdb{
		URL: url,
		c:   c,
		ctx: context.Background(),
	}
}

// Login authenticates with the TVDB API and sets the bearer token for future requests
func (c *Tvdb) Login(apiKey string) error {
	c.Key = apiKey

	b := PostLoginJSONRequestBody{}
	b.Apikey = apiKey

	resp, err := c.c.PostLoginWithResponse(c.ctx, b)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return errors.Errorf("unexpected status code: %d", resp.StatusCode())
	}
	if resp == nil || resp.JSON200 == nil || resp.JSON200.Data == nil || resp.JSON200.Data.Token == nil {
		return errors.New("response does not contain valid token")
	}

	return c.SetAuthToken(*resp.JSON200.Data.Token)
}

// SetAuthToken sets the bearer token for future requests
//
// Use this if you have your auth token stored elsewhere. These tokens
// expire after 30 days, so it's best to store in your database and refresh
// when needed.
func (c *Tvdb) SetAuthToken(token string) (err error) {
	if token == "" {
		return errors.Errorf("token cannot be empty")
	}

	c.token = token
	c.ctx = context.WithValue(c.ctx, contextBearerToken, token)

	c.c, err = newClientWithResponses(c.URL, WithRequestEditorFn(requestAuth))
	if err != nil {
		return err
	}

	return nil
}

// requestAuth wires up the bearer token to the request
func requestAuth(ctx context.Context, req *http.Request) error {
	token, ok := ctx.Value(contextBearerToken).(string)
	if !ok {
		return errors.Errorf("no bearer token found in context")
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	return nil
}

// GetAllArtworkStatusesWithResponse is a wrapper for ClientWithResponses.GetAllArtworkStatusesWithResponse
func (c *Tvdb) GetAllArtworkStatuses() (*GetAllArtworkStatusesResponse, error) {
	return c.c.GetAllArtworkStatusesWithResponse(c.ctx)
}

// GetAllArtworkTypesWithResponse is a wrapper for ClientWithResponses.GetAllArtworkTypesWithResponse
func (c *Tvdb) GetAllArtworkTypes() (*GetAllArtworkTypesResponse, error) {
	return c.c.GetAllArtworkTypesWithResponse(c.ctx)
}

// GetArtworkBaseWithResponse is a wrapper for ClientWithResponses.GetArtworkBaseWithResponse
func (c *Tvdb) GetArtworkBase(id float32) (*GetArtworkBaseResponse, error) {
	return c.c.GetArtworkBaseWithResponse(c.ctx, id)
}

// GetArtworkExtendedWithResponse is a wrapper for ClientWithResponses.GetArtworkExtendedWithResponse
func (c *Tvdb) GetArtworkExtended(id float32) (*GetArtworkExtendedResponse, error) {
	return c.c.GetArtworkExtendedWithResponse(c.ctx, id)
}

// GetAllAwardsWithResponse is a wrapper for ClientWithResponses.GetAllAwardsWithResponse
func (c *Tvdb) GetAllAwards() (*GetAllAwardsResponse, error) {
	return c.c.GetAllAwardsWithResponse(c.ctx)
}

// GetAwardCategoryWithResponse is a wrapper for ClientWithResponses.GetAwardCategoryWithResponse
func (c *Tvdb) GetAwardCategory(id float32) (*GetAwardCategoryResponse, error) {
	return c.c.GetAwardCategoryWithResponse(c.ctx, id)
}

// GetAwardCategoryExtendedWithResponse is a wrapper for ClientWithResponses.GetAwardCategoryExtendedWithResponse
func (c *Tvdb) GetAwardCategoryExtended(id float32) (*GetAwardCategoryExtendedResponse, error) {
	return c.c.GetAwardCategoryExtendedWithResponse(c.ctx, id)
}

// GetAwardWithResponse is a wrapper for ClientWithResponses.GetAwardWithResponse
func (c *Tvdb) GetAward(id float32) (*GetAwardResponse, error) {
	return c.c.GetAwardWithResponse(c.ctx, id)
}

// GetAwardExtendedWithResponse is a wrapper for ClientWithResponses.GetAwardExtendedWithResponse
func (c *Tvdb) GetAwardExtended(id float32) (*GetAwardExtendedResponse, error) {
	return c.c.GetAwardExtendedWithResponse(c.ctx, id)
}

// GetCharacterBaseWithResponse is a wrapper for ClientWithResponses.GetCharacterBaseWithResponse
func (c *Tvdb) GetCharacterBase(id float32) (*GetCharacterBaseResponse, error) {
	return c.c.GetCharacterBaseWithResponse(c.ctx, id)
}

// GetAllCompaniesWithResponse is a wrapper for ClientWithResponses.GetAllCompaniesWithResponse
func (c *Tvdb) GetAllCompanies(params *GetAllCompaniesParams) (*GetAllCompaniesResponse, error) {
	return c.c.GetAllCompaniesWithResponse(c.ctx, params)
}

// GetCompanyTypesWithResponse is a wrapper for ClientWithResponses.GetCompanyTypesWithResponse
func (c *Tvdb) GetCompanyTypes() (*GetCompanyTypesResponse, error) {
	return c.c.GetCompanyTypesWithResponse(c.ctx)
}

// GetCompanyWithResponse is a wrapper for ClientWithResponses.GetCompanyWithResponse
func (c *Tvdb) GetCompany(id float32) (*GetCompanyResponse, error) {
	return c.c.GetCompanyWithResponse(c.ctx, id)
}

// GetAllContentRatingsWithResponse is a wrapper for ClientWithResponses.GetAllContentRatingsWithResponse
func (c *Tvdb) GetAllContentRatings() (*GetAllContentRatingsResponse, error) {
	return c.c.GetAllContentRatingsWithResponse(c.ctx)
}

// GetAllCountriesWithResponse is a wrapper for ClientWithResponses.GetAllCountriesWithResponse
func (c *Tvdb) GetAllCountries() (*GetAllCountriesResponse, error) {
	return c.c.GetAllCountriesWithResponse(c.ctx)
}

// GetEntityTypesWithResponse is a wrapper for ClientWithResponses.GetEntityTypesWithResponse
func (c *Tvdb) GetEntityTypes() (*GetEntityTypesResponse, error) {
	return c.c.GetEntityTypesWithResponse(c.ctx)
}

// GetAllEpisodesWithResponse is a wrapper for ClientWithResponses.GetAllEpisodesWithResponse
func (c *Tvdb) GetAllEpisodes(params *GetAllEpisodesParams) (*GetAllEpisodesResponse, error) {
	return c.c.GetAllEpisodesWithResponse(c.ctx, params)
}

// GetEpisodeBaseWithResponse is a wrapper for ClientWithResponses.GetEpisodeBaseWithResponse
func (c *Tvdb) GetEpisodeBase(id float32) (*GetEpisodeBaseResponse, error) {
	return c.c.GetEpisodeBaseWithResponse(c.ctx, id)
}

// GetEpisodeExtendedWithResponse is a wrapper for ClientWithResponses.GetEpisodeExtendedWithResponse
func (c *Tvdb) GetEpisodeExtended(id float32, params *GetEpisodeExtendedParams) (*GetEpisodeExtendedResponse, error) {
	return c.c.GetEpisodeExtendedWithResponse(c.ctx, id, params)
}

// GetEpisodeTranslationWithResponse is a wrapper for ClientWithResponses.GetEpisodeTranslationWithResponse
func (c *Tvdb) GetEpisodeTranslation(id float32, language string) (*GetEpisodeTranslationResponse, error) {
	return c.c.GetEpisodeTranslationWithResponse(c.ctx, id, language)
}

// GetAllGendersWithResponse is a wrapper for ClientWithResponses.GetAllGendersWithResponse
func (c *Tvdb) GetAllGenders() (*GetAllGendersResponse, error) {
	return c.c.GetAllGendersWithResponse(c.ctx)
}

// GetAllGenresWithResponse is a wrapper for ClientWithResponses.GetAllGenresWithResponse
func (c *Tvdb) GetAllGenres() (*GetAllGenresResponse, error) {
	return c.c.GetAllGenresWithResponse(c.ctx)
}

// GetGenreBaseWithResponse is a wrapper for ClientWithResponses.GetGenreBaseWithResponse
func (c *Tvdb) GetGenreBase(id float32) (*GetGenreBaseResponse, error) {
	return c.c.GetGenreBaseWithResponse(c.ctx, id)
}

// GetAllInspirationTypesWithResponse is a wrapper for ClientWithResponses.GetAllInspirationTypesWithResponse
func (c *Tvdb) GetAllInspirationTypes() (*GetAllInspirationTypesResponse, error) {
	return c.c.GetAllInspirationTypesWithResponse(c.ctx)
}

// GetAllLanguagesWithResponse is a wrapper for ClientWithResponses.GetAllLanguagesWithResponse
func (c *Tvdb) GetAllLanguages() (*GetAllLanguagesResponse, error) {
	return c.c.GetAllLanguagesWithResponse(c.ctx)
}

// GetAllListsWithResponse is a wrapper for ClientWithResponses.GetAllListsWithResponse
func (c *Tvdb) GetAllLists(params *GetAllListsParams) (*GetAllListsResponse, error) {
	return c.c.GetAllListsWithResponse(c.ctx, params)
}

// GetListBySlugWithResponse is a wrapper for ClientWithResponses.GetListBySlugWithResponse
func (c *Tvdb) GetListBySlug(slug string) (*GetListBySlugResponse, error) {
	return c.c.GetListBySlugWithResponse(c.ctx, slug)
}

// GetListWithResponse is a wrapper for ClientWithResponses.GetListWithResponse
func (c *Tvdb) GetList(id float32) (*GetListResponse, error) {
	return c.c.GetListWithResponse(c.ctx, id)
}

// GetListExtendedWithResponse is a wrapper for ClientWithResponses.GetListExtendedWithResponse
func (c *Tvdb) GetListExtended(id float32) (*GetListExtendedResponse, error) {
	return c.c.GetListExtendedWithResponse(c.ctx, id)
}

// GetListTranslationWithResponse is a wrapper for ClientWithResponses.GetListTranslationWithResponse
func (c *Tvdb) GetListTranslation(id float32, language string) (*GetListTranslationResponse, error) {
	return c.c.GetListTranslationWithResponse(c.ctx, id, language)
}

// GetAllMovieWithResponse is a wrapper for ClientWithResponses.GetAllMovieWithResponse
func (c *Tvdb) GetAllMovie(params *GetAllMovieParams) (*GetAllMovieResponse, error) {
	return c.c.GetAllMovieWithResponse(c.ctx, params)
}

// GetMoviesFilterWithResponse is a wrapper for ClientWithResponses.GetMoviesFilterWithResponse
func (c *Tvdb) GetMoviesFilter(params *GetMoviesFilterParams) (*GetMoviesFilterResponse, error) {
	return c.c.GetMoviesFilterWithResponse(c.ctx, params)
}

// GetMovieBaseBySlugWithResponse is a wrapper for ClientWithResponses.GetMovieBaseBySlugWithResponse
func (c *Tvdb) GetMovieBaseBySlug(slug string) (*GetMovieBaseBySlugResponse, error) {
	return c.c.GetMovieBaseBySlugWithResponse(c.ctx, slug)
}

// GetAllMovieStatusesWithResponse is a wrapper for ClientWithResponses.GetAllMovieStatusesWithResponse
func (c *Tvdb) GetAllMovieStatuses() (*GetAllMovieStatusesResponse, error) {
	return c.c.GetAllMovieStatusesWithResponse(c.ctx)
}

// GetMovieBaseWithResponse is a wrapper for ClientWithResponses.GetMovieBaseWithResponse
func (c *Tvdb) GetMovieBase(id float32) (*GetMovieBaseResponse, error) {
	return c.c.GetMovieBaseWithResponse(c.ctx, id)
}

// GetMovieExtendedWithResponse is a wrapper for ClientWithResponses.GetMovieExtendedWithResponse
func (c *Tvdb) GetMovieExtended(id float32, params *GetMovieExtendedParams) (*GetMovieExtendedResponse, error) {
	return c.c.GetMovieExtendedWithResponse(c.ctx, id, params)
}

// GetMovieTranslationWithResponse is a wrapper for ClientWithResponses.GetMovieTranslationWithResponse
func (c *Tvdb) GetMovieTranslation(id float32, language string) (*GetMovieTranslationResponse, error) {
	return c.c.GetMovieTranslationWithResponse(c.ctx, id, language)
}

// GetAllPeopleWithResponse is a wrapper for ClientWithResponses.GetAllPeopleWithResponse
func (c *Tvdb) GetAllPeople(params *GetAllPeopleParams) (*GetAllPeopleResponse, error) {
	return c.c.GetAllPeopleWithResponse(c.ctx, params)
}

// GetAllPeopleTypesWithResponse is a wrapper for ClientWithResponses.GetAllPeopleTypesWithResponse
func (c *Tvdb) GetAllPeopleTypes() (*GetAllPeopleTypesResponse, error) {
	return c.c.GetAllPeopleTypesWithResponse(c.ctx)
}

// GetPeopleBaseWithResponse is a wrapper for ClientWithResponses.GetPeopleBaseWithResponse
func (c *Tvdb) GetPeopleBase(id float32) (*GetPeopleBaseResponse, error) {
	return c.c.GetPeopleBaseWithResponse(c.ctx, id)
}

// GetPeopleExtendedWithResponse is a wrapper for ClientWithResponses.GetPeopleExtendedWithResponse
func (c *Tvdb) GetPeopleExtended(id float32, params *GetPeopleExtendedParams) (*GetPeopleExtendedResponse, error) {
	return c.c.GetPeopleExtendedWithResponse(c.ctx, id, params)
}

// GetPeopleTranslationWithResponse is a wrapper for ClientWithResponses.GetPeopleTranslationWithResponse
func (c *Tvdb) GetPeopleTranslation(id float32, language string) (*GetPeopleTranslationResponse, error) {
	return c.c.GetPeopleTranslationWithResponse(c.ctx, id, language)
}

// GetSearchResultsWithResponse is a wrapper for ClientWithResponses.GetSearchResultsWithResponse
func (c *Tvdb) GetSearchResults(params *GetSearchResultsParams) (*GetSearchResultsResponse, error) {
	return c.c.GetSearchResultsWithResponse(c.ctx, params)
}

// GetSearchResultsByRemoteIdWithResponse is a wrapper for ClientWithResponses.GetSearchResultsByRemoteIdWithResponse
func (c *Tvdb) GetSearchResultsByRemoteId(remoteId string) (*GetSearchResultsByRemoteIdResponse, error) {
	return c.c.GetSearchResultsByRemoteIdWithResponse(c.ctx, remoteId)
}

// GetAllSeasonsWithResponse is a wrapper for ClientWithResponses.GetAllSeasonsWithResponse
func (c *Tvdb) GetAllSeasons(params *GetAllSeasonsParams) (*GetAllSeasonsResponse, error) {
	return c.c.GetAllSeasonsWithResponse(c.ctx, params)
}

// GetSeasonTypesWithResponse is a wrapper for ClientWithResponses.GetSeasonTypesWithResponse
func (c *Tvdb) GetSeasonTypes() (*GetSeasonTypesResponse, error) {
	return c.c.GetSeasonTypesWithResponse(c.ctx)
}

// GetSeasonBaseWithResponse is a wrapper for ClientWithResponses.GetSeasonBaseWithResponse
func (c *Tvdb) GetSeasonBase(id float32) (*GetSeasonBaseResponse, error) {
	return c.c.GetSeasonBaseWithResponse(c.ctx, id)
}

// GetSeasonExtendedWithResponse is a wrapper for ClientWithResponses.GetSeasonExtendedWithResponse
func (c *Tvdb) GetSeasonExtended(id float32) (*GetSeasonExtendedResponse, error) {
	return c.c.GetSeasonExtendedWithResponse(c.ctx, id)
}

// GetSeasonTranslationWithResponse is a wrapper for ClientWithResponses.GetSeasonTranslationWithResponse
func (c *Tvdb) GetSeasonTranslation(id float32, language string) (*GetSeasonTranslationResponse, error) {
	return c.c.GetSeasonTranslationWithResponse(c.ctx, id, language)
}

// GetAllSeriesWithResponse is a wrapper for ClientWithResponses.GetAllSeriesWithResponse
func (c *Tvdb) GetAllSeries(params *GetAllSeriesParams) (*GetAllSeriesResponse, error) {
	return c.c.GetAllSeriesWithResponse(c.ctx, params)
}

// GetSeriesFilterWithResponse is a wrapper for ClientWithResponses.GetSeriesFilterWithResponse
func (c *Tvdb) GetSeriesFilter(params *GetSeriesFilterParams) (*GetSeriesFilterResponse, error) {
	return c.c.GetSeriesFilterWithResponse(c.ctx, params)
}

// GetSeriesBaseBySlugWithResponse is a wrapper for ClientWithResponses.GetSeriesBaseBySlugWithResponse
func (c *Tvdb) GetSeriesBaseBySlug(slug string) (*GetSeriesBaseBySlugResponse, error) {
	return c.c.GetSeriesBaseBySlugWithResponse(c.ctx, slug)
}

// GetAllSeriesStatusesWithResponse is a wrapper for ClientWithResponses.GetAllSeriesStatusesWithResponse
func (c *Tvdb) GetAllSeriesStatuses() (*GetAllSeriesStatusesResponse, error) {
	return c.c.GetAllSeriesStatusesWithResponse(c.ctx)
}

// GetSeriesBaseWithResponse is a wrapper for ClientWithResponses.GetSeriesBaseWithResponse
func (c *Tvdb) GetSeriesBase(id float32) (*GetSeriesBaseResponse, error) {
	return c.c.GetSeriesBaseWithResponse(c.ctx, id)
}

// GetSeriesArtworksWithResponse is a wrapper for ClientWithResponses.GetSeriesArtworksWithResponse
func (c *Tvdb) GetSeriesArtworks(id float32, params *GetSeriesArtworksParams) (*GetSeriesArtworksResponse, error) {
	return c.c.GetSeriesArtworksWithResponse(c.ctx, id, params)
}

// GetSeriesEpisodesWithResponse is a wrapper for ClientWithResponses.GetSeriesEpisodesWithResponse
func (c *Tvdb) GetSeriesEpisodes(id float32, seasonType string, params *GetSeriesEpisodesParams) (*GetSeriesEpisodesResponse, error) {
	return c.c.GetSeriesEpisodesWithResponse(c.ctx, id, seasonType, params)
}

// GetSeriesSeasonEpisodesTranslatedWithResponse is a wrapper for ClientWithResponses.GetSeriesSeasonEpisodesTranslatedWithResponse
func (c *Tvdb) GetSeriesSeasonEpisodesTranslated(id float32, seasonType string, lang string, params *GetSeriesSeasonEpisodesTranslatedParams) (*GetSeriesSeasonEpisodesTranslatedResponse, error) {
	return c.c.GetSeriesSeasonEpisodesTranslatedWithResponse(c.ctx, id, seasonType, lang, params)
}

// GetSeriesExtendedWithResponse is a wrapper for ClientWithResponses.GetSeriesExtendedWithResponse
func (c *Tvdb) GetSeriesExtended(id float32, params *GetSeriesExtendedParams) (*GetSeriesExtendedResponse, error) {
	return c.c.GetSeriesExtendedWithResponse(c.ctx, id, params)
}

// GetSeriesNextAiredWithResponse is a wrapper for ClientWithResponses.GetSeriesNextAiredWithResponse
func (c *Tvdb) GetSeriesNextAired(id float32) (*GetSeriesNextAiredResponse, error) {
	return c.c.GetSeriesNextAiredWithResponse(c.ctx, id)
}

// GetSeriesTranslationWithResponse is a wrapper for ClientWithResponses.GetSeriesTranslationWithResponse
func (c *Tvdb) GetSeriesTranslation(id float32, language string) (*GetSeriesTranslationResponse, error) {
	return c.c.GetSeriesTranslationWithResponse(c.ctx, id, language)
}

// GetAllSourceTypesWithResponse is a wrapper for ClientWithResponses.GetAllSourceTypesWithResponse
func (c *Tvdb) GetAllSourceTypes() (*GetAllSourceTypesResponse, error) {
	return c.c.GetAllSourceTypesWithResponse(c.ctx)
}

// UpdatesWithResponse is a wrapper for ClientWithResponses.UpdatesWithResponse
func (c *Tvdb) Updates(params *UpdatesParams) (*UpdatesResponse, error) {
	return c.c.UpdatesWithResponse(c.ctx, params)
}

// GetUserInfoWithResponse is a wrapper for ClientWithResponses.GetUserInfoWithResponse
func (c *Tvdb) GetUserInfo() (*GetUserInfoResponse, error) {
	return c.c.GetUserInfoWithResponse(c.ctx)
}

// GetUserFavoritesWithResponse is a wrapper for ClientWithResponses.GetUserFavoritesWithResponse
func (c *Tvdb) GetUserFavorites() (*GetUserFavoritesResponse, error) {
	return c.c.GetUserFavoritesWithResponse(c.ctx)
}

// CreateUserFavoritesWithResponse is a wrapper for ClientWithResponses.CreateUserFavoritesWithResponse
func (c *Tvdb) CreateUserFavorites(body CreateUserFavoritesJSONRequestBody) (*CreateUserFavoritesResponse, error) {
	return c.c.CreateUserFavoritesWithResponse(c.ctx, body)
}

// GetUserInfoByIdWithResponse is a wrapper for ClientWithResponses.GetUserInfoByIdWithResponse
func (c *Tvdb) GetUserInfoById(id float32) (*GetUserInfoByIdResponse, error) {
	return c.c.GetUserInfoByIdWithResponse(c.ctx, id)
}

// Use Login() instead
//
//	func (c *Tvdb) PostLoginWithBodyWithResponse(contentType string, body io.Reader) (*PostLoginResponse, error) {
//		return c.c.PostLoginWithBodyWithResponse(c.ctx, contentType, body)
//	}
//
//	func (c *Tvdb) PostLoginWithResponse(body PostLoginJSONRequestBody) (*PostLoginResponse, error) {
//		return c.c.PostLoginWithResponse(c.ctx, body)
//	}
