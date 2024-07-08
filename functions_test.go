package tvdb

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dashotv/tvdb/openapi/models/operations"
)

func TestClient_GetArtworkBase(t *testing.T) {
	c := testClient(t)
	var id float64 = artwork_id_float64
	r, err := c.GetArtworkBase(id)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetArtworkExtended(t *testing.T) {
	c := testClient(t)
	var id float64 = artwork_id_float64
	r, err := c.GetArtworkExtended(id)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetAllArtworkStatuses(t *testing.T) {
	c := testClient(t)

	r, err := c.GetAllArtworkStatuses()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetAllArtworkTypes(t *testing.T) {
	c := testClient(t)

	r, err := c.GetAllArtworkTypes()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetAwardCategory(t *testing.T) {
	c := testClient(t)
	var id float64 = awardcategories_id_float64
	r, err := c.GetAwardCategory(id)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetAwardCategoryExtended(t *testing.T) {
	c := testClient(t)
	var id float64 = awardcategories_id_float64
	r, err := c.GetAwardCategoryExtended(id)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetAllAwards(t *testing.T) {
	c := testClient(t)

	r, err := c.GetAllAwards()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetAward(t *testing.T) {
	c := testClient(t)
	var id float64 = awards_id_float64
	r, err := c.GetAward(id)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetAwardExtended(t *testing.T) {
	c := testClient(t)
	var id float64 = awards_id_float64
	r, err := c.GetAwardExtended(id)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetCharacterBase(t *testing.T) {
	c := testClient(t)
	var id float64 = characters_id_float64
	r, err := c.GetCharacterBase(id)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetAllCompanies(t *testing.T) {
	c := testClient(t)
	var page *float64 = nil
	r, err := c.GetAllCompanies(page)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetCompany(t *testing.T) {
	c := testClient(t)
	var id float64 = companies_id_float64
	r, err := c.GetCompany(id)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetCompanyTypes(t *testing.T) {
	c := testClient(t)

	r, err := c.GetCompanyTypes()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetAllContentRatings(t *testing.T) {
	c := testClient(t)

	r, err := c.GetAllContentRatings()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetAllCountries(t *testing.T) {
	c := testClient(t)

	r, err := c.GetAllCountries()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetEntityTypes(t *testing.T) {
	c := testClient(t)

	r, err := c.GetEntityTypes()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetAllEpisodes(t *testing.T) {
	c := testClient(t)
	var page *float64 = nil
	r, err := c.GetAllEpisodes(page)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetEpisodeBase(t *testing.T) {
	c := testClient(t)
	var id float64 = episodes_id_float64
	r, err := c.GetEpisodeBase(id)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetEpisodeExtended(t *testing.T) {
	c := testClient(t)
	var id float64 = episodes_id_float64
	var meta *operations.Meta = nil
	r, err := c.GetEpisodeExtended(id, meta)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetEpisodeTranslation(t *testing.T) {
	c := testClient(t)
	var id float64 = episodes_id_float64
	var language string = episodes_language_string
	r, err := c.GetEpisodeTranslation(id, language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetAllGenders(t *testing.T) {
	c := testClient(t)

	r, err := c.GetAllGenders()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetAllGenres(t *testing.T) {
	c := testClient(t)

	r, err := c.GetAllGenres()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetGenreBase(t *testing.T) {
	c := testClient(t)
	var id float64 = genres_id_float64
	r, err := c.GetGenreBase(id)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetAllInspirationTypes(t *testing.T) {
	c := testClient(t)

	r, err := c.GetAllInspirationTypes()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetAllLanguages(t *testing.T) {
	c := testClient(t)

	r, err := c.GetAllLanguages()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetAllLists(t *testing.T) {
	c := testClient(t)
	var page *float64 = nil
	r, err := c.GetAllLists(page)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetList(t *testing.T) {
	c := testClient(t)
	var id float64 = lists_id_float64
	r, err := c.GetList(id)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetListBySlug(t *testing.T) {
	c := testClient(t)
	var slug string = lists_slug_string
	r, err := c.GetListBySlug(slug)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetListExtended(t *testing.T) {
	c := testClient(t)
	var id float64 = lists_id_float64
	r, err := c.GetListExtended(id)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetListTranslation(t *testing.T) {
	c := testClient(t)
	var id float64 = lists_id_float64
	var language string = lists_language_string
	r, err := c.GetListTranslation(id, language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_PostLogin(t *testing.T) {
	c := testClient(t)
	var request operations.PostLoginRequestBody = operations_PostLoginRequestBody
	r, err := c.PostLogin(request)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetAllMovie(t *testing.T) {
	c := testClient(t)
	var page *float64 = nil
	r, err := c.GetAllMovie(page)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetMovieBase(t *testing.T) {
	c := testClient(t)
	var id float64 = movies_id_float64
	r, err := c.GetMovieBase(id)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetMovieBaseBySlug(t *testing.T) {
	c := testClient(t)
	var slug string = movies_slug_string
	r, err := c.GetMovieBaseBySlug(slug)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetMovieExtended(t *testing.T) {
	c := testClient(t)
	var id float64 = movies_id_float64
	var meta *operations.QueryParamMeta = nil
	var short *bool = nil
	r, err := c.GetMovieExtended(id, meta, short)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetMovieTranslation(t *testing.T) {
	c := testClient(t)
	var id float64 = movies_id_float64
	var language string = movies_language_string
	r, err := c.GetMovieTranslation(id, language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetMoviesFilter(t *testing.T) {
	c := testClient(t)
	var request operations.GetMoviesFilterRequest = operations_GetMoviesFilterRequest
	r, err := c.GetMoviesFilter(request)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetAllMovieStatuses(t *testing.T) {
	c := testClient(t)

	r, err := c.GetAllMovieStatuses()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetAllPeople(t *testing.T) {
	c := testClient(t)
	var page *float64 = nil
	r, err := c.GetAllPeople(page)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetPeopleBase(t *testing.T) {
	c := testClient(t)
	var id float64 = people_id_float64
	r, err := c.GetPeopleBase(id)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetPeopleExtended(t *testing.T) {
	c := testClient(t)
	var id float64 = people_id_float64
	var meta *operations.GetPeopleExtendedQueryParamMeta = nil
	r, err := c.GetPeopleExtended(id, meta)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetPeopleTranslation(t *testing.T) {
	c := testClient(t)
	var id float64 = people_id_float64
	var language string = people_language_string
	r, err := c.GetPeopleTranslation(id, language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetAllPeopleTypes(t *testing.T) {
	c := testClient(t)

	r, err := c.GetAllPeopleTypes()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetSearchResults(t *testing.T) {
	c := testClient(t)
	var request operations.GetSearchResultsRequest = operations_GetSearchResultsRequest
	r, err := c.GetSearchResults(request)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetSearchResultsByRemoteID(t *testing.T) {
	c := testClient(t)
	var remoteID string = search_remoteID_string
	r, err := c.GetSearchResultsByRemoteID(remoteID)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetAllSeasons(t *testing.T) {
	c := testClient(t)
	var page *float64 = nil
	r, err := c.GetAllSeasons(page)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetSeasonBase(t *testing.T) {
	c := testClient(t)
	var id float64 = seasons_id_float64
	r, err := c.GetSeasonBase(id)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetSeasonExtended(t *testing.T) {
	c := testClient(t)
	var id float64 = seasons_id_float64
	r, err := c.GetSeasonExtended(id)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetSeasonTypes(t *testing.T) {
	c := testClient(t)

	r, err := c.GetSeasonTypes()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetAllSeries(t *testing.T) {
	c := testClient(t)
	var page *float64 = nil
	r, err := c.GetAllSeries(page)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetSeriesArtworks(t *testing.T) {
	c := testClient(t)
	var id float64 = series_id_float64
	var lang *string = nil
	var type_ *int64 = nil
	r, err := c.GetSeriesArtworks(id, lang, type_)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetSeriesBase(t *testing.T) {
	c := testClient(t)
	var id float64 = series_id_float64
	r, err := c.GetSeriesBase(id)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetSeriesBaseBySlug(t *testing.T) {
	c := testClient(t)
	var slug string = series_slug_string
	r, err := c.GetSeriesBaseBySlug(slug)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetSeriesEpisodes(t *testing.T) {
	c := testClient(t)
	var request operations.GetSeriesEpisodesRequest = operations_GetSeriesEpisodesRequest
	r, err := c.GetSeriesEpisodes(request)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetSeriesExtended(t *testing.T) {
	c := testClient(t)
	var id float64 = series_id_float64
	var meta *operations.GetSeriesExtendedQueryParamMeta = nil
	var short *bool = nil
	r, err := c.GetSeriesExtended(id, meta, short)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetSeriesFilter(t *testing.T) {
	c := testClient(t)
	var request operations.GetSeriesFilterRequest = operations_GetSeriesFilterRequest
	r, err := c.GetSeriesFilter(request)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetSeriesNextAired(t *testing.T) {
	c := testClient(t)
	var id float64 = series_id_float64
	r, err := c.GetSeriesNextAired(id)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetSeriesSeasonEpisodesTranslated(t *testing.T) {
	c := testClient(t)
	var id float64 = series_id_float64
	var lang string = series_lang_string
	var page int64 = series_page_int64
	var seasonType string = series_seasonType_string
	r, err := c.GetSeriesSeasonEpisodesTranslated(id, lang, page, seasonType)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetSeriesTranslation(t *testing.T) {
	c := testClient(t)
	var id float64 = series_id_float64
	var language string = series_language_string
	r, err := c.GetSeriesTranslation(id, language)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetAllSeriesStatuses(t *testing.T) {
	c := testClient(t)

	r, err := c.GetAllSeriesStatuses()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_GetAllSourceTypes(t *testing.T) {
	c := testClient(t)

	r, err := c.GetAllSourceTypes()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestClient_Updates(t *testing.T) {
	c := testClient(t)
	var since int64 = updates_since_int64
	var action *operations.Action = nil
	var page *float64 = nil
	var type_ *operations.Type = nil
	r, err := c.Updates(since, action, page, type_)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}
