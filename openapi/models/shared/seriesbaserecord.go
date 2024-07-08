// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// SeriesBaseRecord - The base record for a series. All series airs time like firstAired, lastAired, nextAired, etc. are in US EST for US series, and for all non-US series, the time of the show’s country capital or most populous city. For streaming services, is the official release time. See https://support.thetvdb.com/kb/faq.php?id=29.
type SeriesBaseRecord struct {
	Aliases              []Alias             `json:"aliases,omitempty"`
	AverageRuntime       *int64              `json:"averageRuntime,omitempty"`
	Country              *string             `json:"country,omitempty"`
	DefaultSeasonType    *int64              `json:"defaultSeasonType,omitempty"`
	Episodes             []EpisodeBaseRecord `json:"episodes,omitempty"`
	FirstAired           *string             `json:"firstAired,omitempty"`
	ID                   *int64              `json:"id,omitempty"`
	Image                *string             `json:"image,omitempty"`
	IsOrderRandomized    *bool               `json:"isOrderRandomized,omitempty"`
	LastAired            *string             `json:"lastAired,omitempty"`
	LastUpdated          *string             `json:"lastUpdated,omitempty"`
	Name                 *string             `json:"name,omitempty"`
	NameTranslations     []string            `json:"nameTranslations,omitempty"`
	NextAired            *string             `json:"nextAired,omitempty"`
	OriginalCountry      *string             `json:"originalCountry,omitempty"`
	OriginalLanguage     *string             `json:"originalLanguage,omitempty"`
	Overview             *string             `json:"overview,omitempty"`
	OverviewTranslations []string            `json:"overviewTranslations,omitempty"`
	Score                *float64            `json:"score,omitempty"`
	Slug                 *string             `json:"slug,omitempty"`
	// status record
	Status *Status `json:"status,omitempty"`
	Year   *string `json:"year,omitempty"`
}

func (o *SeriesBaseRecord) GetAliases() []Alias {
	if o == nil {
		return nil
	}
	return o.Aliases
}

func (o *SeriesBaseRecord) GetAverageRuntime() *int64 {
	if o == nil {
		return nil
	}
	return o.AverageRuntime
}

func (o *SeriesBaseRecord) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *SeriesBaseRecord) GetDefaultSeasonType() *int64 {
	if o == nil {
		return nil
	}
	return o.DefaultSeasonType
}

func (o *SeriesBaseRecord) GetEpisodes() []EpisodeBaseRecord {
	if o == nil {
		return nil
	}
	return o.Episodes
}

func (o *SeriesBaseRecord) GetFirstAired() *string {
	if o == nil {
		return nil
	}
	return o.FirstAired
}

func (o *SeriesBaseRecord) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SeriesBaseRecord) GetImage() *string {
	if o == nil {
		return nil
	}
	return o.Image
}

func (o *SeriesBaseRecord) GetIsOrderRandomized() *bool {
	if o == nil {
		return nil
	}
	return o.IsOrderRandomized
}

func (o *SeriesBaseRecord) GetLastAired() *string {
	if o == nil {
		return nil
	}
	return o.LastAired
}

func (o *SeriesBaseRecord) GetLastUpdated() *string {
	if o == nil {
		return nil
	}
	return o.LastUpdated
}

func (o *SeriesBaseRecord) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *SeriesBaseRecord) GetNameTranslations() []string {
	if o == nil {
		return nil
	}
	return o.NameTranslations
}

func (o *SeriesBaseRecord) GetNextAired() *string {
	if o == nil {
		return nil
	}
	return o.NextAired
}

func (o *SeriesBaseRecord) GetOriginalCountry() *string {
	if o == nil {
		return nil
	}
	return o.OriginalCountry
}

func (o *SeriesBaseRecord) GetOriginalLanguage() *string {
	if o == nil {
		return nil
	}
	return o.OriginalLanguage
}

func (o *SeriesBaseRecord) GetOverview() *string {
	if o == nil {
		return nil
	}
	return o.Overview
}

func (o *SeriesBaseRecord) GetOverviewTranslations() []string {
	if o == nil {
		return nil
	}
	return o.OverviewTranslations
}

func (o *SeriesBaseRecord) GetScore() *float64 {
	if o == nil {
		return nil
	}
	return o.Score
}

func (o *SeriesBaseRecord) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *SeriesBaseRecord) GetStatus() *Status {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *SeriesBaseRecord) GetYear() *string {
	if o == nil {
		return nil
	}
	return o.Year
}
