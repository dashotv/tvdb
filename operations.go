package tvdb

import (
	"context"

	"github.com/pkg/errors"

	"github.com/dashotv/tvdb/openapi/models/operations"
)

// type CreateUserFavorites = operations.CreateUserFavoritesResponse

type GetArtworkBase = operations.GetArtworkBase200ApplicationJSON
type GetArtworkExtended = operations.GetArtworkExtended200ApplicationJSON
type GetAllArtworkStatuses = operations.GetAllArtworkStatuses200ApplicationJSON
type GetAllArtworkTypes = operations.GetAllArtworkTypes200ApplicationJSON
type GetAwardCategory = operations.GetAwardCategory200ApplicationJSON
type GetAwardCategoryExtended = operations.GetAwardCategoryExtended200ApplicationJSON
type GetAllAwards = operations.GetAllAwards200ApplicationJSON
type GetAward = operations.GetAward200ApplicationJSON
type GetAwardExtended = operations.GetAwardExtended200ApplicationJSON
type GetCharacterBase = operations.GetCharacterBase200ApplicationJSON
type GetAllCompanies = operations.GetAllCompanies200ApplicationJSON
type GetCompany = operations.GetCompany200ApplicationJSON
type GetCompanyTypes = operations.GetCompanyTypes200ApplicationJSON
type GetAllContentRatings = operations.GetAllContentRatings200ApplicationJSON
type GetAllCountries = operations.GetAllCountries200ApplicationJSON
type GetEntityTypes = operations.GetEntityTypes200ApplicationJSON
type GetAllEpisodes = operations.GetAllEpisodes200ApplicationJSON
type GetEpisodeBase = operations.GetEpisodeBase200ApplicationJSON
type GetEpisodeExtended = operations.GetEpisodeExtended200ApplicationJSON
type GetEpisodeTranslation = operations.GetEpisodeTranslation200ApplicationJSON
type GetUserFavorites = operations.GetUserFavorites200ApplicationJSON
type GetAllGenders = operations.GetAllGenders200ApplicationJSON
type GetAllGenres = operations.GetAllGenres200ApplicationJSON
type GetGenreBase = operations.GetGenreBase200ApplicationJSON
type GetAllInspirationTypes = operations.GetAllInspirationTypes200ApplicationJSON
type GetAllLanguages = operations.GetAllLanguages200ApplicationJSON
type GetAllLists = operations.GetAllLists200ApplicationJSON
type GetList = operations.GetList200ApplicationJSON
type GetListBySlug = operations.GetListBySlug200ApplicationJSON
type GetListExtended = operations.GetListExtended200ApplicationJSON
type GetListTranslation = operations.GetListTranslation200ApplicationJSON
type PostLogin = operations.PostLogin200ApplicationJSON
type GetAllMovie = operations.GetAllMovie200ApplicationJSON
type GetMovieBase = operations.GetMovieBase200ApplicationJSON
type GetMovieBaseBySlug = operations.GetMovieBaseBySlug200ApplicationJSON
type GetMovieExtended = operations.GetMovieExtended200ApplicationJSON
type GetMovieTranslation = operations.GetMovieTranslation200ApplicationJSON
type GetMoviesFilter = operations.GetMoviesFilter200ApplicationJSON
type GetAllMovieStatuses = operations.GetAllMovieStatuses200ApplicationJSON
type GetAllPeople = operations.GetAllPeople200ApplicationJSON
type GetPeopleBase = operations.GetPeopleBase200ApplicationJSON
type GetPeopleExtended = operations.GetPeopleExtended200ApplicationJSON
type GetPeopleTranslation = operations.GetPeopleTranslation200ApplicationJSON
type GetAllPeopleTypes = operations.GetAllPeopleTypes200ApplicationJSON
type GetSearchResults = operations.GetSearchResults200ApplicationJSON
type GetSearchResultsByRemoteID = operations.GetSearchResultsByRemoteID200ApplicationJSON
type GetAllSeasons = operations.GetAllSeasons200ApplicationJSON
type GetSeasonBase = operations.GetSeasonBase200ApplicationJSON
type GetSeasonExtended = operations.GetSeasonExtended200ApplicationJSON
type GetSeasonTranslation = operations.GetSeasonTranslation200ApplicationJSON
type GetSeasonTypes = operations.GetSeasonTypes200ApplicationJSON
type GetAllSeries = operations.GetAllSeries200ApplicationJSON
type GetSeriesArtworks = operations.GetSeriesArtworks200ApplicationJSON
type GetSeriesBase = operations.GetSeriesBase200ApplicationJSON
type GetSeriesBaseBySlug = operations.GetSeriesBaseBySlug200ApplicationJSON
type GetSeriesEpisodes = operations.GetSeriesEpisodes200ApplicationJSON
type GetSeriesExtended = operations.GetSeriesExtended200ApplicationJSON
type GetSeriesFilter = operations.GetSeriesFilter200ApplicationJSON
type GetSeriesNextAired = operations.GetSeriesNextAired200ApplicationJSON
type GetSeriesSeasonEpisodesTranslated = operations.GetSeriesSeasonEpisodesTranslated200ApplicationJSON
type GetSeriesTranslation = operations.GetSeriesTranslation200ApplicationJSON
type GetAllSeriesStatuses = operations.GetAllSeriesStatuses200ApplicationJSON
type GetAllSourceTypes = operations.GetAllSourceTypes200ApplicationJSON
type Updates = operations.Updates200ApplicationJSON
type GetUserInfo = operations.GetUserInfo200ApplicationJSON
type GetUserInfoByID = operations.GetUserInfoByID200ApplicationJSON

// GetArtworkBase wraps the generated openapi.SDK client call
func (c *Client) GetArtworkBase(id float64) (*GetArtworkBase, error) {
	r, err := c.sdk.Artwork.GetArtworkBase(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetArtworkBase200ApplicationJSONObject, nil
}

// GetArtworkExtended wraps the generated openapi.SDK client call
func (c *Client) GetArtworkExtended(id float64) (*GetArtworkExtended, error) {
	r, err := c.sdk.Artwork.GetArtworkExtended(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetArtworkExtended200ApplicationJSONObject, nil
}

// GetAllArtworkStatuses wraps the generated openapi.SDK client call
func (c *Client) GetAllArtworkStatuses(ctx context.Context) (*GetAllArtworkStatuses, error) {
	r, err := c.sdk.ArtworkStatuses.GetAllArtworkStatuses(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllArtworkStatuses200ApplicationJSONObject, nil
}

// GetAllArtworkTypes wraps the generated openapi.SDK client call
func (c *Client) GetAllArtworkTypes(ctx context.Context) (*GetAllArtworkTypes, error) {
	r, err := c.sdk.ArtworkTypes.GetAllArtworkTypes(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllArtworkTypes200ApplicationJSONObject, nil
}

// GetAwardCategory wraps the generated openapi.SDK client call
func (c *Client) GetAwardCategory(id float64) (*GetAwardCategory, error) {
	r, err := c.sdk.AwardCategories.GetAwardCategory(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAwardCategory200ApplicationJSONObject, nil
}

// GetAwardCategoryExtended wraps the generated openapi.SDK client call
func (c *Client) GetAwardCategoryExtended(id float64) (*GetAwardCategoryExtended, error) {
	r, err := c.sdk.AwardCategories.GetAwardCategoryExtended(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAwardCategoryExtended200ApplicationJSONObject, nil
}

// GetAllAwards wraps the generated openapi.SDK client call
func (c *Client) GetAllAwards(ctx context.Context) (*GetAllAwards, error) {
	r, err := c.sdk.Awards.GetAllAwards(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllAwards200ApplicationJSONObject, nil
}

// GetAward wraps the generated openapi.SDK client call
func (c *Client) GetAward(id float64) (*GetAward, error) {
	r, err := c.sdk.Awards.GetAward(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAward200ApplicationJSONObject, nil
}

// GetAwardExtended wraps the generated openapi.SDK client call
func (c *Client) GetAwardExtended(id float64) (*GetAwardExtended, error) {
	r, err := c.sdk.Awards.GetAwardExtended(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAwardExtended200ApplicationJSONObject, nil
}

// GetCharacterBase wraps the generated openapi.SDK client call
func (c *Client) GetCharacterBase(id float64) (*GetCharacterBase, error) {
	r, err := c.sdk.Characters.GetCharacterBase(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetCharacterBase200ApplicationJSONObject, nil
}

// GetAllCompanies wraps the generated openapi.SDK client call
func (c *Client) GetAllCompanies(page *float64) (*GetAllCompanies, error) {
	r, err := c.sdk.Companies.GetAllCompanies(c.ctx, page)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllCompanies200ApplicationJSONObject, nil
}

// GetCompany wraps the generated openapi.SDK client call
func (c *Client) GetCompany(id float64) (*GetCompany, error) {
	r, err := c.sdk.Companies.GetCompany(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetCompany200ApplicationJSONObject, nil
}

// GetCompanyTypes wraps the generated openapi.SDK client call
func (c *Client) GetCompanyTypes(ctx context.Context) (*GetCompanyTypes, error) {
	r, err := c.sdk.Companies.GetCompanyTypes(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetCompanyTypes200ApplicationJSONObject, nil
}

// GetAllContentRatings wraps the generated openapi.SDK client call
func (c *Client) GetAllContentRatings(ctx context.Context) (*GetAllContentRatings, error) {
	r, err := c.sdk.ContentRatings.GetAllContentRatings(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllContentRatings200ApplicationJSONObject, nil
}

// GetAllCountries wraps the generated openapi.SDK client call
func (c *Client) GetAllCountries(ctx context.Context) (*GetAllCountries, error) {
	r, err := c.sdk.Countries.GetAllCountries(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllCountries200ApplicationJSONObject, nil
}

// GetEntityTypes wraps the generated openapi.SDK client call
func (c *Client) GetEntityTypes(ctx context.Context) (*GetEntityTypes, error) {
	r, err := c.sdk.EntityTypes.GetEntityTypes(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetEntityTypes200ApplicationJSONObject, nil
}

// GetAllEpisodes wraps the generated openapi.SDK client call
func (c *Client) GetAllEpisodes(page *float64) (*GetAllEpisodes, error) {
	r, err := c.sdk.Episodes.GetAllEpisodes(c.ctx, page)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllEpisodes200ApplicationJSONObject, nil
}

// GetEpisodeBase wraps the generated openapi.SDK client call
func (c *Client) GetEpisodeBase(id float64) (*GetEpisodeBase, error) {
	r, err := c.sdk.Episodes.GetEpisodeBase(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetEpisodeBase200ApplicationJSONObject, nil
}

// GetEpisodeExtended wraps the generated openapi.SDK client call
func (c *Client) GetEpisodeExtended(id float64, meta *operations.GetEpisodeExtendedMeta) (*GetEpisodeExtended, error) {
	r, err := c.sdk.Episodes.GetEpisodeExtended(c.ctx, id, meta)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetEpisodeExtended200ApplicationJSONObject, nil
}

// GetEpisodeTranslation wraps the generated openapi.SDK client call
func (c *Client) GetEpisodeTranslation(id float64, language string) (*GetEpisodeTranslation, error) {
	r, err := c.sdk.Episodes.GetEpisodeTranslation(c.ctx, id, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetEpisodeTranslation200ApplicationJSONObject, nil
}

// CreateUserFavorites wraps the generated openapi.SDK client call
//
//	func (c *Client) CreateUserFavorites(request *shared.FavoriteRecord) (*CreateUserFavorites, error) {
//		r, err := c.sdk.Favorites.CreateUserFavorites(c.ctx, request)
//		if err != nil {
//			return nil, err
//		}
//		if r.StatusCode != 200 {
//			return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
//		}
//		return r.CreateUserFavorites200ApplicationJSONObject, nil
//	}
//
// GetUserFavorites wraps the generated openapi.SDK client call
func (c *Client) GetUserFavorites(ctx context.Context) (*GetUserFavorites, error) {
	r, err := c.sdk.Favorites.GetUserFavorites(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetUserFavorites200ApplicationJSONObject, nil
}

// GetAllGenders wraps the generated openapi.SDK client call
func (c *Client) GetAllGenders(ctx context.Context) (*GetAllGenders, error) {
	r, err := c.sdk.Genders.GetAllGenders(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllGenders200ApplicationJSONObject, nil
}

// GetAllGenres wraps the generated openapi.SDK client call
func (c *Client) GetAllGenres(ctx context.Context) (*GetAllGenres, error) {
	r, err := c.sdk.Genres.GetAllGenres(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllGenres200ApplicationJSONObject, nil
}

// GetGenreBase wraps the generated openapi.SDK client call
func (c *Client) GetGenreBase(id float64) (*GetGenreBase, error) {
	r, err := c.sdk.Genres.GetGenreBase(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetGenreBase200ApplicationJSONObject, nil
}

// GetAllInspirationTypes wraps the generated openapi.SDK client call
func (c *Client) GetAllInspirationTypes(ctx context.Context) (*GetAllInspirationTypes, error) {
	r, err := c.sdk.InspirationTypes.GetAllInspirationTypes(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllInspirationTypes200ApplicationJSONObject, nil
}

// GetAllLanguages wraps the generated openapi.SDK client call
func (c *Client) GetAllLanguages(ctx context.Context) (*GetAllLanguages, error) {
	r, err := c.sdk.Languages.GetAllLanguages(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllLanguages200ApplicationJSONObject, nil
}

// GetAllLists wraps the generated openapi.SDK client call
func (c *Client) GetAllLists(page *float64) (*GetAllLists, error) {
	r, err := c.sdk.Lists.GetAllLists(c.ctx, page)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllLists200ApplicationJSONObject, nil
}

// GetList wraps the generated openapi.SDK client call
func (c *Client) GetList(id float64) (*GetList, error) {
	r, err := c.sdk.Lists.GetList(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetList200ApplicationJSONObject, nil
}

// GetListBySlug wraps the generated openapi.SDK client call
func (c *Client) GetListBySlug(slug string) (*GetListBySlug, error) {
	r, err := c.sdk.Lists.GetListBySlug(c.ctx, slug)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetListBySlug200ApplicationJSONObject, nil
}

// GetListExtended wraps the generated openapi.SDK client call
func (c *Client) GetListExtended(id float64) (*GetListExtended, error) {
	r, err := c.sdk.Lists.GetListExtended(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetListExtended200ApplicationJSONObject, nil
}

// GetListTranslation wraps the generated openapi.SDK client call
func (c *Client) GetListTranslation(id float64, language string) (*GetListTranslation, error) {
	r, err := c.sdk.Lists.GetListTranslation(c.ctx, id, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetListTranslation200ApplicationJSONObject, nil
}

// PostLogin wraps the generated openapi.SDK client call
func (c *Client) PostLogin(request operations.PostLoginRequestBody) (*PostLogin, error) {
	r, err := c.sdk.Login.PostLogin(c.ctx, request)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.PostLogin200ApplicationJSONObject, nil
}

// GetAllMovie wraps the generated openapi.SDK client call
func (c *Client) GetAllMovie(page *float64) (*GetAllMovie, error) {
	r, err := c.sdk.Movies.GetAllMovie(c.ctx, page)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllMovie200ApplicationJSONObject, nil
}

// GetMovieBase wraps the generated openapi.SDK client call
func (c *Client) GetMovieBase(id float64) (*GetMovieBase, error) {
	r, err := c.sdk.Movies.GetMovieBase(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetMovieBase200ApplicationJSONObject, nil
}

// GetMovieBaseBySlug wraps the generated openapi.SDK client call
func (c *Client) GetMovieBaseBySlug(slug string) (*GetMovieBaseBySlug, error) {
	r, err := c.sdk.Movies.GetMovieBaseBySlug(c.ctx, slug)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetMovieBaseBySlug200ApplicationJSONObject, nil
}

// GetMovieExtended wraps the generated openapi.SDK client call
func (c *Client) GetMovieExtended(id float64, meta *operations.GetMovieExtendedMeta, short *bool) (*GetMovieExtended, error) {
	r, err := c.sdk.Movies.GetMovieExtended(c.ctx, id, meta, short)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetMovieExtended200ApplicationJSONObject, nil
}

// GetMovieTranslation wraps the generated openapi.SDK client call
func (c *Client) GetMovieTranslation(id float64, language string) (*GetMovieTranslation, error) {
	r, err := c.sdk.Movies.GetMovieTranslation(c.ctx, id, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetMovieTranslation200ApplicationJSONObject, nil
}

// GetMoviesFilter wraps the generated openapi.SDK client call
func (c *Client) GetMoviesFilter(request operations.GetMoviesFilterRequest) (*GetMoviesFilter, error) {
	r, err := c.sdk.Movies.GetMoviesFilter(c.ctx, request)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetMoviesFilter200ApplicationJSONObject, nil
}

// GetAllMovieStatuses wraps the generated openapi.SDK client call
func (c *Client) GetAllMovieStatuses(ctx context.Context) (*GetAllMovieStatuses, error) {
	r, err := c.sdk.MovieStatuses.GetAllMovieStatuses(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllMovieStatuses200ApplicationJSONObject, nil
}

// GetAllPeople wraps the generated openapi.SDK client call
func (c *Client) GetAllPeople(page *float64) (*GetAllPeople, error) {
	r, err := c.sdk.People.GetAllPeople(c.ctx, page)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllPeople200ApplicationJSONObject, nil
}

// GetPeopleBase wraps the generated openapi.SDK client call
func (c *Client) GetPeopleBase(id float64) (*GetPeopleBase, error) {
	r, err := c.sdk.People.GetPeopleBase(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetPeopleBase200ApplicationJSONObject, nil
}

// GetPeopleExtended wraps the generated openapi.SDK client call
func (c *Client) GetPeopleExtended(id float64, meta *operations.GetPeopleExtendedMeta) (*GetPeopleExtended, error) {
	r, err := c.sdk.People.GetPeopleExtended(c.ctx, id, meta)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetPeopleExtended200ApplicationJSONObject, nil
}

// GetPeopleTranslation wraps the generated openapi.SDK client call
func (c *Client) GetPeopleTranslation(id float64, language string) (*GetPeopleTranslation, error) {
	r, err := c.sdk.People.GetPeopleTranslation(c.ctx, id, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetPeopleTranslation200ApplicationJSONObject, nil
}

// GetAllPeopleTypes wraps the generated openapi.SDK client call
func (c *Client) GetAllPeopleTypes(ctx context.Context) (*GetAllPeopleTypes, error) {
	r, err := c.sdk.PeopleTypes.GetAllPeopleTypes(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllPeopleTypes200ApplicationJSONObject, nil
}

// TODO: figure this
// GetServerDetails wraps the generated openapi.SDK client call out
// func (c *Client) GetServerDetails() (string, map[string]string) {
// 	r, err := c.sdk.SdkConfiguration.GetServerDetails(c.ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if r.StatusCode != 200 {
// 		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
// 	}
// 	return r.Client200ApplicationJSONObject, nil
// }

// GetSearchResults wraps the generated openapi.SDK client call
func (c *Client) GetSearchResults(request operations.GetSearchResultsRequest) (*GetSearchResults, error) {
	r, err := c.sdk.Search.GetSearchResults(c.ctx, request)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetSearchResults200ApplicationJSONObject, nil
}

// GetSearchResultsByRemoteID wraps the generated openapi.SDK client call
func (c *Client) GetSearchResultsByRemoteID(remoteID string) (*GetSearchResultsByRemoteID, error) {
	r, err := c.sdk.Search.GetSearchResultsByRemoteID(c.ctx, remoteID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetSearchResultsByRemoteID200ApplicationJSONObject, nil
}

// GetAllSeasons wraps the generated openapi.SDK client call
func (c *Client) GetAllSeasons(page *float64) (*GetAllSeasons, error) {
	r, err := c.sdk.Seasons.GetAllSeasons(c.ctx, page)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllSeasons200ApplicationJSONObject, nil
}

// GetSeasonBase wraps the generated openapi.SDK client call
func (c *Client) GetSeasonBase(id float64) (*GetSeasonBase, error) {
	r, err := c.sdk.Seasons.GetSeasonBase(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetSeasonBase200ApplicationJSONObject, nil
}

// GetSeasonExtended wraps the generated openapi.SDK client call
func (c *Client) GetSeasonExtended(id float64) (*GetSeasonExtended, error) {
	r, err := c.sdk.Seasons.GetSeasonExtended(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetSeasonExtended200ApplicationJSONObject, nil
}

// GetSeasonTranslation wraps the generated openapi.SDK client call
func (c *Client) GetSeasonTranslation(id float64, language string) (*GetSeasonTranslation, error) {
	r, err := c.sdk.Seasons.GetSeasonTranslation(c.ctx, id, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetSeasonTranslation200ApplicationJSONObject, nil
}

// GetSeasonTypes wraps the generated openapi.SDK client call
func (c *Client) GetSeasonTypes(ctx context.Context) (*GetSeasonTypes, error) {
	r, err := c.sdk.Seasons.GetSeasonTypes(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetSeasonTypes200ApplicationJSONObject, nil
}

// GetAllSeries wraps the generated openapi.SDK client call
func (c *Client) GetAllSeries(page *float64) (*GetAllSeries, error) {
	r, err := c.sdk.Series.GetAllSeries(c.ctx, page)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllSeries200ApplicationJSONObject, nil
}

// GetSeriesArtworks wraps the generated openapi.SDK client call
func (c *Client) GetSeriesArtworks(id float64, lang *string, type_ *int64) (*GetSeriesArtworks, error) {
	r, err := c.sdk.Series.GetSeriesArtworks(c.ctx, id, lang, type_)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetSeriesArtworks200ApplicationJSONObject, nil
}

// GetSeriesBase wraps the generated openapi.SDK client call
func (c *Client) GetSeriesBase(id float64) (*GetSeriesBase, error) {
	r, err := c.sdk.Series.GetSeriesBase(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetSeriesBase200ApplicationJSONObject, nil
}

// GetSeriesBaseBySlug wraps the generated openapi.SDK client call
func (c *Client) GetSeriesBaseBySlug(slug string) (*GetSeriesBaseBySlug, error) {
	r, err := c.sdk.Series.GetSeriesBaseBySlug(c.ctx, slug)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetSeriesBaseBySlug200ApplicationJSONObject, nil
}

// GetSeriesEpisodes wraps the generated openapi.SDK client call
func (c *Client) GetSeriesEpisodes(request operations.GetSeriesEpisodesRequest) (*GetSeriesEpisodes, error) {
	r, err := c.sdk.Series.GetSeriesEpisodes(c.ctx, request)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetSeriesEpisodes200ApplicationJSONObject, nil
}

// GetSeriesExtended wraps the generated openapi.SDK client call
func (c *Client) GetSeriesExtended(id float64, meta *operations.GetSeriesExtendedMeta, short *bool) (*GetSeriesExtended, error) {
	r, err := c.sdk.Series.GetSeriesExtended(c.ctx, id, meta, short)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetSeriesExtended200ApplicationJSONObject, nil
}

// GetSeriesFilter wraps the generated openapi.SDK client call
func (c *Client) GetSeriesFilter(request operations.GetSeriesFilterRequest) (*GetSeriesFilter, error) {
	r, err := c.sdk.Series.GetSeriesFilter(c.ctx, request)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetSeriesFilter200ApplicationJSONObject, nil
}

// GetSeriesNextAired wraps the generated openapi.SDK client call
func (c *Client) GetSeriesNextAired(id float64) (*GetSeriesNextAired, error) {
	r, err := c.sdk.Series.GetSeriesNextAired(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetSeriesNextAired200ApplicationJSONObject, nil
}

// GetSeriesSeasonEpisodesTranslated wraps the generated openapi.SDK client call
func (c *Client) GetSeriesSeasonEpisodesTranslated(id float64, lang string, page int64, seasonType string) (*GetSeriesSeasonEpisodesTranslated, error) {
	r, err := c.sdk.Series.GetSeriesSeasonEpisodesTranslated(c.ctx, id, lang, page, seasonType)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetSeriesSeasonEpisodesTranslated200ApplicationJSONObject, nil
}

// GetSeriesTranslation wraps the generated openapi.SDK client call
func (c *Client) GetSeriesTranslation(id float64, language string) (*GetSeriesTranslation, error) {
	r, err := c.sdk.Series.GetSeriesTranslation(c.ctx, id, language)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetSeriesTranslation200ApplicationJSONObject, nil
}

// GetAllSeriesStatuses wraps the generated openapi.SDK client call
func (c *Client) GetAllSeriesStatuses(ctx context.Context) (*GetAllSeriesStatuses, error) {
	r, err := c.sdk.SeriesStatuses.GetAllSeriesStatuses(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllSeriesStatuses200ApplicationJSONObject, nil
}

// GetAllSourceTypes wraps the generated openapi.SDK client call
func (c *Client) GetAllSourceTypes(ctx context.Context) (*GetAllSourceTypes, error) {
	r, err := c.sdk.SourceTypes.GetAllSourceTypes(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetAllSourceTypes200ApplicationJSONObject, nil
}

// Updates wraps the generated openapi.SDK client call
func (c *Client) Updates(since float64, action *operations.UpdatesAction, page *float64, type_ *operations.UpdatesType) (*Updates, error) {
	r, err := c.sdk.Updates.Updates(c.ctx, since, action, page, type_)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.Updates200ApplicationJSONObject, nil
}

// GetUserInfo wraps the generated openapi.SDK client call
func (c *Client) GetUserInfo(ctx context.Context) (*GetUserInfo, error) {
	r, err := c.sdk.UserInfo.GetUserInfo(c.ctx)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetUserInfo200ApplicationJSONObject, nil
}

// GetUserInfoByID wraps the generated openapi.SDK client call
func (c *Client) GetUserInfoByID(id float64) (*GetUserInfoByID, error) {
	r, err := c.sdk.UserInfo.GetUserInfoByID(c.ctx, id)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.GetUserInfoByID200ApplicationJSONObject, nil
}
