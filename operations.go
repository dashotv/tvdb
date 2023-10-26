package tvdb

func (s *artwork) GetArtworkBase(ctx context.Context, id float64) (*operations.GetArtworkBaseResponse, error) {
func (s *artwork) GetArtworkExtended(ctx context.Context, id float64) (*operations.GetArtworkExtendedResponse, error) {
func (s *artworkStatuses) GetAllArtworkStatuses(ctx context.Context) (*operations.GetAllArtworkStatusesResponse, error) {
func (s *artworkTypes) GetAllArtworkTypes(ctx context.Context) (*operations.GetAllArtworkTypesResponse, error) {
func (s *awardCategories) GetAwardCategory(ctx context.Context, id float64) (*operations.GetAwardCategoryResponse, error) {
func (s *awardCategories) GetAwardCategoryExtended(ctx context.Context, id float64) (*operations.GetAwardCategoryExtendedResponse, error) {
func (s *awards) GetAllAwards(ctx context.Context) (*operations.GetAllAwardsResponse, error) {
func (s *awards) GetAward(ctx context.Context, id float64) (*operations.GetAwardResponse, error) {
func (s *awards) GetAwardExtended(ctx context.Context, id float64) (*operations.GetAwardExtendedResponse, error) {
func (s *characters) GetCharacterBase(ctx context.Context, id float64) (*operations.GetCharacterBaseResponse, error) {
func (s *companies) GetAllCompanies(ctx context.Context, page *float64) (*operations.GetAllCompaniesResponse, error) {
func (s *companies) GetCompany(ctx context.Context, id float64) (*operations.GetCompanyResponse, error) {
func (s *companies) GetCompanyTypes(ctx context.Context) (*operations.GetCompanyTypesResponse, error) {
func (s *contentRatings) GetAllContentRatings(ctx context.Context) (*operations.GetAllContentRatingsResponse, error) {
func (s *countries) GetAllCountries(ctx context.Context) (*operations.GetAllCountriesResponse, error) {
func (s *entityTypes) GetEntityTypes(ctx context.Context) (*operations.GetEntityTypesResponse, error) {
func (s *episodes) GetAllEpisodes(ctx context.Context, page *float64) (*operations.GetAllEpisodesResponse, error) {
func (s *episodes) GetEpisodeBase(ctx context.Context, id float64) (*operations.GetEpisodeBaseResponse, error) {
func (s *episodes) GetEpisodeExtended(ctx context.Context, id float64, meta *operations.GetEpisodeExtendedMeta) (*operations.GetEpisodeExtendedResponse, error) {
func (s *episodes) GetEpisodeTranslation(ctx context.Context, id float64, language string) (*operations.GetEpisodeTranslationResponse, error) {
func (s *favorites) CreateUserFavorites(ctx context.Context, request *shared.FavoriteRecord) (*operations.CreateUserFavoritesResponse, error) {
func (s *favorites) GetUserFavorites(ctx context.Context) (*operations.GetUserFavoritesResponse, error) {
func (s *genders) GetAllGenders(ctx context.Context) (*operations.GetAllGendersResponse, error) {
func (s *genres) GetAllGenres(ctx context.Context) (*operations.GetAllGenresResponse, error) {
func (s *genres) GetGenreBase(ctx context.Context, id float64) (*operations.GetGenreBaseResponse, error) {
func (s *inspirationTypes) GetAllInspirationTypes(ctx context.Context) (*operations.GetAllInspirationTypesResponse, error) {
func (s *languages) GetAllLanguages(ctx context.Context) (*operations.GetAllLanguagesResponse, error) {
func (s *lists) GetAllLists(ctx context.Context, page *float64) (*operations.GetAllListsResponse, error) {
func (s *lists) GetList(ctx context.Context, id float64) (*operations.GetListResponse, error) {
func (s *lists) GetListBySlug(ctx context.Context, slug string) (*operations.GetListBySlugResponse, error) {
func (s *lists) GetListExtended(ctx context.Context, id float64) (*operations.GetListExtendedResponse, error) {
func (s *lists) GetListTranslation(ctx context.Context, id float64, language string) (*operations.GetListTranslationResponse, error) {
func (s *login) PostLogin(ctx context.Context, request operations.PostLoginRequestBody) (*operations.PostLoginResponse, error) {
func (s *movies) GetAllMovie(ctx context.Context, page *float64) (*operations.GetAllMovieResponse, error) {
func (s *movies) GetMovieBase(ctx context.Context, id float64) (*operations.GetMovieBaseResponse, error) {
func (s *movies) GetMovieBaseBySlug(ctx context.Context, slug string) (*operations.GetMovieBaseBySlugResponse, error) {
func (s *movies) GetMovieExtended(ctx context.Context, id float64, meta *operations.GetMovieExtendedMeta, short *bool) (*operations.GetMovieExtendedResponse, error) {
func (s *movies) GetMovieTranslation(ctx context.Context, id float64, language string) (*operations.GetMovieTranslationResponse, error) {
func (s *movies) GetMoviesFilter(ctx context.Context, request operations.GetMoviesFilterRequest) (*operations.GetMoviesFilterResponse, error) {
func (s *movieStatuses) GetAllMovieStatuses(ctx context.Context) (*operations.GetAllMovieStatusesResponse, error) {
func (s *people) GetAllPeople(ctx context.Context, page *float64) (*operations.GetAllPeopleResponse, error) {
func (s *people) GetPeopleBase(ctx context.Context, id float64) (*operations.GetPeopleBaseResponse, error) {
func (s *people) GetPeopleExtended(ctx context.Context, id float64, meta *operations.GetPeopleExtendedMeta) (*operations.GetPeopleExtendedResponse, error) {
func (s *people) GetPeopleTranslation(ctx context.Context, id float64, language string) (*operations.GetPeopleTranslationResponse, error) {
func (s *peopleTypes) GetAllPeopleTypes(ctx context.Context) (*operations.GetAllPeopleTypesResponse, error) {
func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
func (s *search) GetSearchResults(ctx context.Context, request operations.GetSearchResultsRequest) (*operations.GetSearchResultsResponse, error) {
func (s *search) GetSearchResultsByRemoteID(ctx context.Context, remoteID string) (*operations.GetSearchResultsByRemoteIDResponse, error) {
func (s *seasons) GetAllSeasons(ctx context.Context, page *float64) (*operations.GetAllSeasonsResponse, error) {
func (s *seasons) GetSeasonBase(ctx context.Context, id float64) (*operations.GetSeasonBaseResponse, error) {
func (s *seasons) GetSeasonExtended(ctx context.Context, id float64) (*operations.GetSeasonExtendedResponse, error) {
func (s *seasons) GetSeasonTranslation(ctx context.Context, id float64, language string) (*operations.GetSeasonTranslationResponse, error) {
func (s *seasons) GetSeasonTypes(ctx context.Context) (*operations.GetSeasonTypesResponse, error) {
func (s *series) GetAllSeries(ctx context.Context, page *float64) (*operations.GetAllSeriesResponse, error) {
func (s *series) GetSeriesArtworks(ctx context.Context, id float64, lang *string, type_ *int64) (*operations.GetSeriesArtworksResponse, error) {
func (s *series) GetSeriesBase(ctx context.Context, id float64) (*operations.GetSeriesBaseResponse, error) {
func (s *series) GetSeriesBaseBySlug(ctx context.Context, slug string) (*operations.GetSeriesBaseBySlugResponse, error) {
func (s *series) GetSeriesEpisodes(ctx context.Context, request operations.GetSeriesEpisodesRequest) (*operations.GetSeriesEpisodesResponse, error) {
func (s *series) GetSeriesExtended(ctx context.Context, id float64, meta *operations.GetSeriesExtendedMeta, short *bool) (*operations.GetSeriesExtendedResponse, error) {
func (s *series) GetSeriesFilter(ctx context.Context, request operations.GetSeriesFilterRequest) (*operations.GetSeriesFilterResponse, error) {
func (s *series) GetSeriesNextAired(ctx context.Context, id float64) (*operations.GetSeriesNextAiredResponse, error) {
func (s *series) GetSeriesSeasonEpisodesTranslated(ctx context.Context, id float64, lang string, page int64, seasonType string) (*operations.GetSeriesSeasonEpisodesTranslatedResponse, error) {
func (s *series) GetSeriesTranslation(ctx context.Context, id float64, language string) (*operations.GetSeriesTranslationResponse, error) {
func (s *seriesStatuses) GetAllSeriesStatuses(ctx context.Context) (*operations.GetAllSeriesStatusesResponse, error) {
func (s *sourceTypes) GetAllSourceTypes(ctx context.Context) (*operations.GetAllSourceTypesResponse, error) {
func (s *updates) Updates(ctx context.Context, since float64, action *operations.UpdatesAction, page *float64, type_ *operations.UpdatesType) (*operations.UpdatesResponse, error) {
func (s *userInfo) GetUserInfo(ctx context.Context) (*operations.GetUserInfoResponse, error) {
func (s *userInfo) GetUserInfoByID(ctx context.Context, id float64) (*operations.GetUserInfoByIDResponse, error) {
