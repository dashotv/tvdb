package tvdb

import (
	"github.com/pkg/errors"

	"github.com/dashotv/tvdb/openapi/models/operations"
)

// GetArtworkBase wraps the generated openapi.SDK.Artwork.GetArtworkBase call
func (c *Client) GetArtworkBase(id float64) (*GetArtworkBaseResponse, error) {
	r, err := c.sdk.Artwork.GetArtworkBase(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetArtworkBase200ApplicationJSONObject, nil
}

// GetArtworkExtended wraps the generated openapi.SDK.Artwork.GetArtworkExtended call
func (c *Client) GetArtworkExtended(id float64) (*GetArtworkExtendedResponse, error) {
	r, err := c.sdk.Artwork.GetArtworkExtended(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetArtworkExtended200ApplicationJSONObject, nil
}

// GetAllArtworkStatuses wraps the generated openapi.SDK.ArtworkStatuses.GetAllArtworkStatuses call
func (c *Client) GetAllArtworkStatuses() (*GetAllArtworkStatusesResponse, error) {
	r, err := c.sdk.ArtworkStatuses.GetAllArtworkStatuses(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllArtworkStatuses200ApplicationJSONObject, nil
}

// GetAllArtworkTypes wraps the generated openapi.SDK.ArtworkTypes.GetAllArtworkTypes call
func (c *Client) GetAllArtworkTypes() (*GetAllArtworkTypesResponse, error) {
	r, err := c.sdk.ArtworkTypes.GetAllArtworkTypes(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllArtworkTypes200ApplicationJSONObject, nil
}

// GetAwardCategory wraps the generated openapi.SDK.AwardCategories.GetAwardCategory call
func (c *Client) GetAwardCategory(id float64) (*GetAwardCategoryResponse, error) {
	r, err := c.sdk.AwardCategories.GetAwardCategory(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAwardCategory200ApplicationJSONObject, nil
}

// GetAwardCategoryExtended wraps the generated openapi.SDK.AwardCategories.GetAwardCategoryExtended call
func (c *Client) GetAwardCategoryExtended(id float64) (*GetAwardCategoryExtendedResponse, error) {
	r, err := c.sdk.AwardCategories.GetAwardCategoryExtended(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAwardCategoryExtended200ApplicationJSONObject, nil
}

// GetAllAwards wraps the generated openapi.SDK.Awards.GetAllAwards call
func (c *Client) GetAllAwards() (*GetAllAwardsResponse, error) {
	r, err := c.sdk.Awards.GetAllAwards(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllAwards200ApplicationJSONObject, nil
}

// GetAward wraps the generated openapi.SDK.Awards.GetAward call
func (c *Client) GetAward(id float64) (*GetAwardResponse, error) {
	r, err := c.sdk.Awards.GetAward(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAward200ApplicationJSONObject, nil
}

// GetAwardExtended wraps the generated openapi.SDK.Awards.GetAwardExtended call
func (c *Client) GetAwardExtended(id float64) (*GetAwardExtendedResponse, error) {
	r, err := c.sdk.Awards.GetAwardExtended(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAwardExtended200ApplicationJSONObject, nil
}

// GetCharacterBase wraps the generated openapi.SDK.Characters.GetCharacterBase call
func (c *Client) GetCharacterBase(id float64) (*GetCharacterBaseResponse, error) {
	r, err := c.sdk.Characters.GetCharacterBase(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetCharacterBase200ApplicationJSONObject, nil
}

// GetAllCompanies wraps the generated openapi.SDK.Companies.GetAllCompanies call
func (c *Client) GetAllCompanies(page *float64) (*GetAllCompaniesResponse, error) {
	r, err := c.sdk.Companies.GetAllCompanies(c.ctx, page)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllCompanies200ApplicationJSONObject, nil
}

// GetCompany wraps the generated openapi.SDK.Companies.GetCompany call
func (c *Client) GetCompany(id float64) (*GetCompanyResponse, error) {
	r, err := c.sdk.Companies.GetCompany(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetCompany200ApplicationJSONObject, nil
}

// GetCompanyTypes wraps the generated openapi.SDK.Companies.GetCompanyTypes call
func (c *Client) GetCompanyTypes() (*GetCompanyTypesResponse, error) {
	r, err := c.sdk.Companies.GetCompanyTypes(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetCompanyTypes200ApplicationJSONObject, nil
}

// GetAllContentRatings wraps the generated openapi.SDK.ContentRatings.GetAllContentRatings call
func (c *Client) GetAllContentRatings() (*GetAllContentRatingsResponse, error) {
	r, err := c.sdk.ContentRatings.GetAllContentRatings(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllContentRatings200ApplicationJSONObject, nil
}

// GetAllCountries wraps the generated openapi.SDK.Countries.GetAllCountries call
func (c *Client) GetAllCountries() (*GetAllCountriesResponse, error) {
	r, err := c.sdk.Countries.GetAllCountries(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllCountries200ApplicationJSONObject, nil
}

// GetEntityTypes wraps the generated openapi.SDK.EntityTypes.GetEntityTypes call
func (c *Client) GetEntityTypes() (*GetEntityTypesResponse, error) {
	r, err := c.sdk.EntityTypes.GetEntityTypes(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetEntityTypes200ApplicationJSONObject, nil
}

// GetAllEpisodes wraps the generated openapi.SDK.Episodes.GetAllEpisodes call
func (c *Client) GetAllEpisodes(page *float64) (*GetAllEpisodesResponse, error) {
	r, err := c.sdk.Episodes.GetAllEpisodes(c.ctx, page)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllEpisodes200ApplicationJSONObject, nil
}

// GetEpisodeBase wraps the generated openapi.SDK.Episodes.GetEpisodeBase call
func (c *Client) GetEpisodeBase(id float64) (*GetEpisodeBaseResponse, error) {
	r, err := c.sdk.Episodes.GetEpisodeBase(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetEpisodeBase200ApplicationJSONObject, nil
}

// GetEpisodeExtended wraps the generated openapi.SDK.Episodes.GetEpisodeExtended call
func (c *Client) GetEpisodeExtended(id float64, meta *operations.GetEpisodeExtendedMeta) (*GetEpisodeExtendedResponse, error) {
	r, err := c.sdk.Episodes.GetEpisodeExtended(c.ctx, id, meta)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetEpisodeExtended200ApplicationJSONObject, nil
}

// GetEpisodeTranslation wraps the generated openapi.SDK.Episodes.GetEpisodeTranslation call
func (c *Client) GetEpisodeTranslation(id float64, language string) (*GetEpisodeTranslationResponse, error) {
	r, err := c.sdk.Episodes.GetEpisodeTranslation(c.ctx, id, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetEpisodeTranslation200ApplicationJSONObject, nil
}

// CreateUserFavorites wraps the generated openapi.SDK.Favorites.CreateUserFavorites call
// func (c *Client) CreateUserFavorites(request *shared.FavoriteRecord) (*CreateUserFavoritesResponse, error) {
//   r, err := c.sdk.Favorites.CreateUserFavorites(c.ctx, request)
//   if err != nil {
//           return nil, err
//   }
//   if r.StatusCode != 200 {
//           return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
//   }
//   return r.CreateUserFavorites200ApplicationJSONObject, nil
// }

// GetUserFavorites wraps the generated openapi.SDK.Favorites.GetUserFavorites call
func (c *Client) GetUserFavorites() (*GetUserFavoritesResponse, error) {
	r, err := c.sdk.Favorites.GetUserFavorites(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetUserFavorites200ApplicationJSONObject, nil
}

// GetAllGenders wraps the generated openapi.SDK.Genders.GetAllGenders call
func (c *Client) GetAllGenders() (*GetAllGendersResponse, error) {
	r, err := c.sdk.Genders.GetAllGenders(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllGenders200ApplicationJSONObject, nil
}

// GetAllGenres wraps the generated openapi.SDK.Genres.GetAllGenres call
func (c *Client) GetAllGenres() (*GetAllGenresResponse, error) {
	r, err := c.sdk.Genres.GetAllGenres(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllGenres200ApplicationJSONObject, nil
}

// GetGenreBase wraps the generated openapi.SDK.Genres.GetGenreBase call
func (c *Client) GetGenreBase(id float64) (*GetGenreBaseResponse, error) {
	r, err := c.sdk.Genres.GetGenreBase(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetGenreBase200ApplicationJSONObject, nil
}

// GetAllInspirationTypes wraps the generated openapi.SDK.InspirationTypes.GetAllInspirationTypes call
func (c *Client) GetAllInspirationTypes() (*GetAllInspirationTypesResponse, error) {
	r, err := c.sdk.InspirationTypes.GetAllInspirationTypes(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllInspirationTypes200ApplicationJSONObject, nil
}

// GetAllLanguages wraps the generated openapi.SDK.Languages.GetAllLanguages call
func (c *Client) GetAllLanguages() (*GetAllLanguagesResponse, error) {
	r, err := c.sdk.Languages.GetAllLanguages(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllLanguages200ApplicationJSONObject, nil
}

// GetAllLists wraps the generated openapi.SDK.Lists.GetAllLists call
func (c *Client) GetAllLists(page *float64) (*GetAllListsResponse, error) {
	r, err := c.sdk.Lists.GetAllLists(c.ctx, page)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllLists200ApplicationJSONObject, nil
}

// GetList wraps the generated openapi.SDK.Lists.GetList call
func (c *Client) GetList(id float64) (*GetListResponse, error) {
	r, err := c.sdk.Lists.GetList(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetList200ApplicationJSONObject, nil
}

// GetListBySlug wraps the generated openapi.SDK.Lists.GetListBySlug call
func (c *Client) GetListBySlug(slug string) (*GetListBySlugResponse, error) {
	r, err := c.sdk.Lists.GetListBySlug(c.ctx, slug)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetListBySlug200ApplicationJSONObject, nil
}

// GetListExtended wraps the generated openapi.SDK.Lists.GetListExtended call
func (c *Client) GetListExtended(id float64) (*GetListExtendedResponse, error) {
	r, err := c.sdk.Lists.GetListExtended(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetListExtended200ApplicationJSONObject, nil
}

// GetListTranslation wraps the generated openapi.SDK.Lists.GetListTranslation call
func (c *Client) GetListTranslation(id float64, language string) (*GetListTranslationResponse, error) {
	r, err := c.sdk.Lists.GetListTranslation(c.ctx, id, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetListTranslation200ApplicationJSONObject, nil
}

// PostLogin wraps the generated openapi.SDK.Login.PostLogin call
func (c *Client) PostLogin(request operations.PostLoginRequestBody) (*PostLoginResponse, error) {
	r, err := c.sdk.Login.PostLogin(c.ctx, request)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.PostLogin200ApplicationJSONObject, nil
}

// GetAllMovie wraps the generated openapi.SDK.Movies.GetAllMovie call
func (c *Client) GetAllMovie(page *float64) (*GetAllMovieResponse, error) {
	r, err := c.sdk.Movies.GetAllMovie(c.ctx, page)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllMovie200ApplicationJSONObject, nil
}

// GetMovieBase wraps the generated openapi.SDK.Movies.GetMovieBase call
func (c *Client) GetMovieBase(id float64) (*GetMovieBaseResponse, error) {
	r, err := c.sdk.Movies.GetMovieBase(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetMovieBase200ApplicationJSONObject, nil
}

// GetMovieBaseBySlug wraps the generated openapi.SDK.Movies.GetMovieBaseBySlug call
func (c *Client) GetMovieBaseBySlug(slug string) (*GetMovieBaseBySlugResponse, error) {
	r, err := c.sdk.Movies.GetMovieBaseBySlug(c.ctx, slug)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetMovieBaseBySlug200ApplicationJSONObject, nil
}

// GetMovieExtended wraps the generated openapi.SDK.Movies.GetMovieExtended call
func (c *Client) GetMovieExtended(id float64, meta *operations.GetMovieExtendedMeta, short *bool) (*GetMovieExtendedResponse, error) {
	r, err := c.sdk.Movies.GetMovieExtended(c.ctx, id, meta, short)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetMovieExtended200ApplicationJSONObject, nil
}

// GetMovieTranslation wraps the generated openapi.SDK.Movies.GetMovieTranslation call
func (c *Client) GetMovieTranslation(id float64, language string) (*GetMovieTranslationResponse, error) {
	r, err := c.sdk.Movies.GetMovieTranslation(c.ctx, id, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetMovieTranslation200ApplicationJSONObject, nil
}

// GetMoviesFilter wraps the generated openapi.SDK.Movies.GetMoviesFilter call
func (c *Client) GetMoviesFilter(request operations.GetMoviesFilterRequest) (*GetMoviesFilterResponse, error) {
	r, err := c.sdk.Movies.GetMoviesFilter(c.ctx, request)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetMoviesFilter200ApplicationJSONObject, nil
}

// GetAllMovieStatuses wraps the generated openapi.SDK.MovieStatuses.GetAllMovieStatuses call
func (c *Client) GetAllMovieStatuses() (*GetAllMovieStatusesResponse, error) {
	r, err := c.sdk.MovieStatuses.GetAllMovieStatuses(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllMovieStatuses200ApplicationJSONObject, nil
}

// GetAllPeople wraps the generated openapi.SDK.People.GetAllPeople call
func (c *Client) GetAllPeople(page *float64) (*GetAllPeopleResponse, error) {
	r, err := c.sdk.People.GetAllPeople(c.ctx, page)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllPeople200ApplicationJSONObject, nil
}

// GetPeopleBase wraps the generated openapi.SDK.People.GetPeopleBase call
func (c *Client) GetPeopleBase(id float64) (*GetPeopleBaseResponse, error) {
	r, err := c.sdk.People.GetPeopleBase(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetPeopleBase200ApplicationJSONObject, nil
}

// GetPeopleExtended wraps the generated openapi.SDK.People.GetPeopleExtended call
func (c *Client) GetPeopleExtended(id float64, meta *operations.GetPeopleExtendedMeta) (*GetPeopleExtendedResponse, error) {
	r, err := c.sdk.People.GetPeopleExtended(c.ctx, id, meta)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetPeopleExtended200ApplicationJSONObject, nil
}

// GetPeopleTranslation wraps the generated openapi.SDK.People.GetPeopleTranslation call
func (c *Client) GetPeopleTranslation(id float64, language string) (*GetPeopleTranslationResponse, error) {
	r, err := c.sdk.People.GetPeopleTranslation(c.ctx, id, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetPeopleTranslation200ApplicationJSONObject, nil
}

// GetAllPeopleTypes wraps the generated openapi.SDK.PeopleTypes.GetAllPeopleTypes call
func (c *Client) GetAllPeopleTypes() (*GetAllPeopleTypesResponse, error) {
	r, err := c.sdk.PeopleTypes.GetAllPeopleTypes(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllPeopleTypes200ApplicationJSONObject, nil
}

// GetSearchResults wraps the generated openapi.SDK.Search.GetSearchResults call
func (c *Client) GetSearchResults(request operations.GetSearchResultsRequest) (*GetSearchResultsResponse, error) {
	r, err := c.sdk.Search.GetSearchResults(c.ctx, request)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetSearchResults200ApplicationJSONObject, nil
}

// GetSearchResultsByRemoteID wraps the generated openapi.SDK.Search.GetSearchResultsByRemoteID call
func (c *Client) GetSearchResultsByRemoteID(remoteID string) (*GetSearchResultsByRemoteIDResponse, error) {
	r, err := c.sdk.Search.GetSearchResultsByRemoteID(c.ctx, remoteID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetSearchResultsByRemoteID200ApplicationJSONObject, nil
}

// GetAllSeasons wraps the generated openapi.SDK.Seasons.GetAllSeasons call
func (c *Client) GetAllSeasons(page *float64) (*GetAllSeasonsResponse, error) {
	r, err := c.sdk.Seasons.GetAllSeasons(c.ctx, page)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllSeasons200ApplicationJSONObject, nil
}

// GetSeasonBase wraps the generated openapi.SDK.Seasons.GetSeasonBase call
func (c *Client) GetSeasonBase(id float64) (*GetSeasonBaseResponse, error) {
	r, err := c.sdk.Seasons.GetSeasonBase(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetSeasonBase200ApplicationJSONObject, nil
}

// GetSeasonExtended wraps the generated openapi.SDK.Seasons.GetSeasonExtended call
func (c *Client) GetSeasonExtended(id float64) (*GetSeasonExtendedResponse, error) {
	r, err := c.sdk.Seasons.GetSeasonExtended(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetSeasonExtended200ApplicationJSONObject, nil
}

// GetSeasonTranslation wraps the generated openapi.SDK.Seasons.GetSeasonTranslation call
func (c *Client) GetSeasonTranslation(id float64, language string) (*GetSeasonTranslationResponse, error) {
	r, err := c.sdk.Seasons.GetSeasonTranslation(c.ctx, id, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetSeasonTranslation200ApplicationJSONObject, nil
}

// GetSeasonTypes wraps the generated openapi.SDK.Seasons.GetSeasonTypes call
func (c *Client) GetSeasonTypes() (*GetSeasonTypesResponse, error) {
	r, err := c.sdk.Seasons.GetSeasonTypes(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetSeasonTypes200ApplicationJSONObject, nil
}

// GetAllSeries wraps the generated openapi.SDK.Series.GetAllSeries call
func (c *Client) GetAllSeries(page *float64) (*GetAllSeriesResponse, error) {
	r, err := c.sdk.Series.GetAllSeries(c.ctx, page)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllSeries200ApplicationJSONObject, nil
}

// GetSeriesArtworks wraps the generated openapi.SDK.Series.GetSeriesArtworks call
func (c *Client) GetSeriesArtworks(id float64, lang *string, type_ *int64) (*GetSeriesArtworksResponse, error) {
	r, err := c.sdk.Series.GetSeriesArtworks(c.ctx, id, lang, type_)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetSeriesArtworks200ApplicationJSONObject, nil
}

// GetSeriesBase wraps the generated openapi.SDK.Series.GetSeriesBase call
func (c *Client) GetSeriesBase(id float64) (*GetSeriesBaseResponse, error) {
	r, err := c.sdk.Series.GetSeriesBase(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetSeriesBase200ApplicationJSONObject, nil
}

// GetSeriesBaseBySlug wraps the generated openapi.SDK.Series.GetSeriesBaseBySlug call
func (c *Client) GetSeriesBaseBySlug(slug string) (*GetSeriesBaseBySlugResponse, error) {
	r, err := c.sdk.Series.GetSeriesBaseBySlug(c.ctx, slug)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetSeriesBaseBySlug200ApplicationJSONObject, nil
}

// GetSeriesEpisodes wraps the generated openapi.SDK.Series.GetSeriesEpisodes call
func (c *Client) GetSeriesEpisodes(request operations.GetSeriesEpisodesRequest) (*GetSeriesEpisodesResponse, error) {
	r, err := c.sdk.Series.GetSeriesEpisodes(c.ctx, request)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetSeriesEpisodes200ApplicationJSONObject, nil
}

// GetSeriesExtended wraps the generated openapi.SDK.Series.GetSeriesExtended call
func (c *Client) GetSeriesExtended(id float64, meta *operations.GetSeriesExtendedMeta, short *bool) (*GetSeriesExtendedResponse, error) {
	r, err := c.sdk.Series.GetSeriesExtended(c.ctx, id, meta, short)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetSeriesExtended200ApplicationJSONObject, nil
}

// GetSeriesFilter wraps the generated openapi.SDK.Series.GetSeriesFilter call
func (c *Client) GetSeriesFilter(request operations.GetSeriesFilterRequest) (*GetSeriesFilterResponse, error) {
	r, err := c.sdk.Series.GetSeriesFilter(c.ctx, request)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetSeriesFilter200ApplicationJSONObject, nil
}

// GetSeriesNextAired wraps the generated openapi.SDK.Series.GetSeriesNextAired call
func (c *Client) GetSeriesNextAired(id float64) (*GetSeriesNextAiredResponse, error) {
	r, err := c.sdk.Series.GetSeriesNextAired(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetSeriesNextAired200ApplicationJSONObject, nil
}

// GetSeriesSeasonEpisodesTranslated wraps the generated openapi.SDK.Series.GetSeriesSeasonEpisodesTranslated call
func (c *Client) GetSeriesSeasonEpisodesTranslated(id float64, lang string, page int64, seasonType string) (*GetSeriesSeasonEpisodesTranslatedResponse, error) {
	r, err := c.sdk.Series.GetSeriesSeasonEpisodesTranslated(c.ctx, id, lang, page, seasonType)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetSeriesSeasonEpisodesTranslated200ApplicationJSONObject, nil
}

// GetSeriesTranslation wraps the generated openapi.SDK.Series.GetSeriesTranslation call
func (c *Client) GetSeriesTranslation(id float64, language string) (*GetSeriesTranslationResponse, error) {
	r, err := c.sdk.Series.GetSeriesTranslation(c.ctx, id, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetSeriesTranslation200ApplicationJSONObject, nil
}

// GetAllSeriesStatuses wraps the generated openapi.SDK.SeriesStatuses.GetAllSeriesStatuses call
func (c *Client) GetAllSeriesStatuses() (*GetAllSeriesStatusesResponse, error) {
	r, err := c.sdk.SeriesStatuses.GetAllSeriesStatuses(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllSeriesStatuses200ApplicationJSONObject, nil
}

// GetAllSourceTypes wraps the generated openapi.SDK.SourceTypes.GetAllSourceTypes call
func (c *Client) GetAllSourceTypes() (*GetAllSourceTypesResponse, error) {
	r, err := c.sdk.SourceTypes.GetAllSourceTypes(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllSourceTypes200ApplicationJSONObject, nil
}

// Updates wraps the generated openapi.SDK.Updates.Updates call
func (c *Client) Updates(since float64, action *operations.UpdatesAction, page *float64, type_ *operations.UpdatesType) (*UpdatesResponse, error) {
	r, err := c.sdk.Updates.Updates(c.ctx, since, action, page, type_)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.Updates200ApplicationJSONObject, nil
}

// GetUserInfo wraps the generated openapi.SDK.UserInfo.GetUserInfo call
func (c *Client) GetUserInfo() (*GetUserInfoResponse, error) {
	r, err := c.sdk.UserInfo.GetUserInfo(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetUserInfo200ApplicationJSONObject, nil
}

// GetUserInfoByID wraps the generated openapi.SDK.UserInfo.GetUserInfoByID call
func (c *Client) GetUserInfoByID(id float64) (*GetUserInfoByIDResponse, error) {
	r, err := c.sdk.UserInfo.GetUserInfoByID(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetUserInfoByID200ApplicationJSONObject, nil
}
