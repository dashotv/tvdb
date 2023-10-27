package tvdb

import "github.com/dashotv/tvdb/openapi/models/operations"

func String(s string) *string    { return &s }
func Bool(b bool) *bool          { return &b }
func Int(i int) *int             { return &i }
func Int64(i int64) *int64       { return &i }
func Float32(f float32) *float32 { return &f }
func Float64(f float64) *float64 { return &f }

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
