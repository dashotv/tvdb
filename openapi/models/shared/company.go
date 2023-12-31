// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Company - A company record
type Company struct {
	ActiveDate           *string  `json:"activeDate,omitempty"`
	Aliases              []Alias  `json:"aliases,omitempty"`
	Country              *string  `json:"country,omitempty"`
	ID                   *int64   `json:"id,omitempty"`
	InactiveDate         *string  `json:"inactiveDate,omitempty"`
	Name                 *string  `json:"name,omitempty"`
	NameTranslations     []string `json:"nameTranslations,omitempty"`
	OverviewTranslations []string `json:"overviewTranslations,omitempty"`
	// A parent company record
	ParentCompany      *ParentCompany `json:"parentCompany,omitempty"`
	PrimaryCompanyType *int64         `json:"primaryCompanyType,omitempty"`
	Slug               *string        `json:"slug,omitempty"`
	TagOptions         []TagOption    `json:"tagOptions,omitempty"`
}

func (o *Company) GetActiveDate() *string {
	if o == nil {
		return nil
	}
	return o.ActiveDate
}

func (o *Company) GetAliases() []Alias {
	if o == nil {
		return nil
	}
	return o.Aliases
}

func (o *Company) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *Company) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Company) GetInactiveDate() *string {
	if o == nil {
		return nil
	}
	return o.InactiveDate
}

func (o *Company) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Company) GetNameTranslations() []string {
	if o == nil {
		return nil
	}
	return o.NameTranslations
}

func (o *Company) GetOverviewTranslations() []string {
	if o == nil {
		return nil
	}
	return o.OverviewTranslations
}

func (o *Company) GetParentCompany() *ParentCompany {
	if o == nil {
		return nil
	}
	return o.ParentCompany
}

func (o *Company) GetPrimaryCompanyType() *int64 {
	if o == nil {
		return nil
	}
	return o.PrimaryCompanyType
}

func (o *Company) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *Company) GetTagOptions() []TagOption {
	if o == nil {
		return nil
	}
	return o.TagOptions
}
