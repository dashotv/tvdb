package tvdb

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tvdbURL = os.Getenv("TVDB_API_URL")
var tvdbKey = os.Getenv("TVDB_API_KEY")
var tvdbToken = os.Getenv("TVDB_API_TOKEN")

func testTvdb(t *testing.T) *Tvdb {
	c := New(tvdbURL)
	assert.NotNil(t, c)

	err := c.SetAuthToken(tvdbToken)
	assert.NoError(t, err)

	return c
}

func TestEnv(t *testing.T) {
	assert.NotNil(t, tvdbURL, "TVDB_API_URL is not set")
	assert.NotEmpty(t, tvdbURL, "TVDB_API_URL is empty")

	assert.NotNil(t, tvdbKey, "TVDB_API_KEY is not set")
	assert.NotEmpty(t, tvdbKey, "TVDB_API_KEY is empty")

	assert.NotNil(t, tvdbToken, "TVDB_API_TOKEN is not set")
	assert.NotEmpty(t, tvdbToken, "TVDB_API_TOKEN is empty")
}

func TestTvdb_Login(t *testing.T) {
	c := New(tvdbURL)
	assert.NotNil(t, c)

	err := c.Login(tvdbKey)
	assert.NoError(t, err)

	_, err = c.GetSearchResults(&GetSearchResultsParams{Query: String("The Simpsons")})
	assert.NoError(t, err)
}

func TestTvdb_SetAuthToken(t *testing.T) {
	c := New(tvdbURL)
	assert.NotNil(t, c)

	err := c.SetAuthToken(tvdbToken)
	assert.NoError(t, err)

	_, err = c.GetSearchResults(&GetSearchResultsParams{Query: String("The Simpsons")})
	assert.NoError(t, err)
}

func TestTvdb_Search(t *testing.T) {
	c := New(tvdbURL)
	assert.NotNil(t, c)

	err := c.Login(tvdbKey)
	assert.NoError(t, err)

	p := &GetSearchResultsParams{Query: String("The Simpsons")}
	r, err := c.GetSearchResults(p)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 38, len(*r.JSON200.Data))
}

func TestTvdb_GetAllArtworkStatuses(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetAllArtworkStatuses()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetAllArtworkTypes(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetAllArtworkTypes()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetArtwork(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetArtwork(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetArtworkExtended(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetArtworkExtended(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetAllAwards(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetAllAwards()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetAwardCategory(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetAwardCategory(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetAwardCategoryExtended(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetAwardCategoryExtended(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetAward(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetAward(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetAwardExtended(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetAwardExtended(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetCharacterBase(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetCharacter(5)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetAllCompanies(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetAllCompanies(&GetAllCompaniesParams{Page: Float32(1)})
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetCompanyTypes(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetCompanyTypes()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetCompany(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetCompany(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetAllContentRatings(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetAllContentRatings()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetAllCountries(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetAllCountries()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetEntityTypes(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetEntityTypes()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetAllEpisodes(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetAllEpisodes(&GetAllEpisodesParams{Page: Float32(1)})
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetEpisodeBase(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetEpisode(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetEpisodeExtended(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetEpisodeExtended(1, &GetEpisodeExtendedParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetEpisodeTranslation(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetEpisodeTranslation(1, "en")
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetAllGenders(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetAllGenders()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetAllGenres(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetAllGenres()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetGenreBase(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetGenre(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetAllInspirationTypes(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetAllInspirationTypes()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetAllLanguages(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetAllLanguages()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetAllLists(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetAllLists(&GetAllListsParams{Page: Float32(1)})
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetListBySlug(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetListBySlug("blarg")
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetList(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetList(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetListExtended(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetListExtended(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetListTranslation(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetListTranslation(1, "en")
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetAllMovie(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetAllMovie(&GetAllMovieParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetMoviesFilter(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetMoviesFilter(&GetMoviesFilterParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetMovieBaseBySlug(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetMovieBySlug("the-simpsons-movie-2007")
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetAllMovieStatuses(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetAllMovieStatuses()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetMovieBase(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetMovie(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetMovieExtended(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetMovieExtended(1, &GetMovieExtendedParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetMovieTranslation(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetMovieTranslation(1, "en")
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetAllPeople(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetAllPeople(&GetAllPeopleParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetAllPeopleTypes(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetAllPeopleTypes()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetPeopleBase(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetPeople(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetPeopleExtended(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetPeopleExtended(1, &GetPeopleExtendedParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetPeopleTranslation(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetPeopleTranslation(1, "en")
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

// func TestTvdb_GetSearchResults up at the top

func TestTvdb_GetSearchResultsByRemoteId(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetSearchResultsByRemoteId("tt0436992")
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetAllSeasons(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetAllSeasons(&GetAllSeasonsParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetSeasonTypes(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetSeasonTypes()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetSeasonBase(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetSeason(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetSeasonExtended(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetSeasonExtended(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetSeasonTranslation(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetSeasonTranslation(1, "en")
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetAllSeries(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetAllSeries(&GetAllSeriesParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetSeriesFilter(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetSeriesFilter(&GetSeriesFilterParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetSeriesBaseBySlug(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetSeriesBySlug("doctor-who-2005")
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetAllSeriesStatuses(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetAllSeriesStatuses()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetSeriesBase(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetSeries(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetSeriesArtworks(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetSeriesArtworks(1, &GetSeriesArtworksParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetSeriesEpisodes(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetSeriesEpisodes(1, "?", &GetSeriesEpisodesParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetSeriesSeasonEpisodesTranslated(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetSeriesSeasonEpisodesTranslated(1, "?", "en", &GetSeriesSeasonEpisodesTranslatedParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetSeriesExtended(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetSeriesExtended(1, &GetSeriesExtendedParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetSeriesNextAired(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetSeriesNextAired(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetSeriesTranslation(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetSeriesTranslation(1, "en")
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetAllSourceTypes(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetAllSourceTypes()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_Updates(t *testing.T) {
	c := testTvdb(t)
	r, err := c.Updates(&UpdatesParams{})
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetUserInfo(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetUserInfo()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestTvdb_GetUserFavorites(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetUserFavorites()
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

// func TestTvdb_CreateUserFavorites(t *testing.T) {
// 	c := testTvdb(t)
// 	r, err := c.CreateUserFavorites(FavoriteRecord{ItemID: 1, ItemType: "series"})
// 	assert.NoError(t, err)
// 	assert.NotNil(t, r)
// }

func TestTvdb_GetUserInfoById(t *testing.T) {
	c := testTvdb(t)
	r, err := c.GetUserInfoById(1)
	assert.NoError(t, err)
	assert.NotNil(t, r)
}
