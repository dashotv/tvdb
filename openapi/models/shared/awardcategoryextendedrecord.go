// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AwardCategoryExtendedRecord - extended award category record
type AwardCategoryExtendedRecord struct {
	AllowCoNominees *bool `json:"allowCoNominees,omitempty"`
	// base award record
	Award     *AwardBaseRecord         `json:"award,omitempty"`
	ForMovies *bool                    `json:"forMovies,omitempty"`
	ForSeries *bool                    `json:"forSeries,omitempty"`
	ID        *int64                   `json:"id,omitempty"`
	Name      *string                  `json:"name,omitempty"`
	Nominees  []AwardNomineeBaseRecord `json:"nominees,omitempty"`
}

func (o *AwardCategoryExtendedRecord) GetAllowCoNominees() *bool {
	if o == nil {
		return nil
	}
	return o.AllowCoNominees
}

func (o *AwardCategoryExtendedRecord) GetAward() *AwardBaseRecord {
	if o == nil {
		return nil
	}
	return o.Award
}

func (o *AwardCategoryExtendedRecord) GetForMovies() *bool {
	if o == nil {
		return nil
	}
	return o.ForMovies
}

func (o *AwardCategoryExtendedRecord) GetForSeries() *bool {
	if o == nil {
		return nil
	}
	return o.ForSeries
}

func (o *AwardCategoryExtendedRecord) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AwardCategoryExtendedRecord) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *AwardCategoryExtendedRecord) GetNominees() []AwardNomineeBaseRecord {
	if o == nil {
		return nil
	}
	return o.Nominees
}
