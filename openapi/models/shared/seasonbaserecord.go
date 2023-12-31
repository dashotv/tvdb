// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// SeasonBaseRecord - season genre record
type SeasonBaseRecord struct {
	// Companies by type record
	Companies            *Companies `json:"companies,omitempty"`
	ID                   *int64     `json:"id,omitempty"`
	Image                *string    `json:"image,omitempty"`
	ImageType            *int64     `json:"imageType,omitempty"`
	LastUpdated          *string    `json:"lastUpdated,omitempty"`
	Name                 *string    `json:"name,omitempty"`
	NameTranslations     []string   `json:"nameTranslations,omitempty"`
	Number               *int64     `json:"number,omitempty"`
	OverviewTranslations []string   `json:"overviewTranslations,omitempty"`
	SeriesID             *int64     `json:"seriesId,omitempty"`
	// season type record
	Type *SeasonType `json:"type,omitempty"`
	Year *string     `json:"year,omitempty"`
}

func (o *SeasonBaseRecord) GetCompanies() *Companies {
	if o == nil {
		return nil
	}
	return o.Companies
}

func (o *SeasonBaseRecord) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SeasonBaseRecord) GetImage() *string {
	if o == nil {
		return nil
	}
	return o.Image
}

func (o *SeasonBaseRecord) GetImageType() *int64 {
	if o == nil {
		return nil
	}
	return o.ImageType
}

func (o *SeasonBaseRecord) GetLastUpdated() *string {
	if o == nil {
		return nil
	}
	return o.LastUpdated
}

func (o *SeasonBaseRecord) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *SeasonBaseRecord) GetNameTranslations() []string {
	if o == nil {
		return nil
	}
	return o.NameTranslations
}

func (o *SeasonBaseRecord) GetNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.Number
}

func (o *SeasonBaseRecord) GetOverviewTranslations() []string {
	if o == nil {
		return nil
	}
	return o.OverviewTranslations
}

func (o *SeasonBaseRecord) GetSeriesID() *int64 {
	if o == nil {
		return nil
	}
	return o.SeriesID
}

func (o *SeasonBaseRecord) GetType() *SeasonType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *SeasonBaseRecord) GetYear() *string {
	if o == nil {
		return nil
	}
	return o.Year
}
