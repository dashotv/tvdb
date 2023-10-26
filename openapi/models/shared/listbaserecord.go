// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ListBaseRecord - base list record
type ListBaseRecord struct {
	Aliases              []Alias     `json:"aliases,omitempty"`
	ID                   *int64      `json:"id,omitempty"`
	Image                *string     `json:"image,omitempty"`
	ImageIsFallback      *bool       `json:"imageIsFallback,omitempty"`
	IsOfficial           *bool       `json:"isOfficial,omitempty"`
	Name                 *string     `json:"name,omitempty"`
	NameTranslations     []string    `json:"nameTranslations,omitempty"`
	Overview             *string     `json:"overview,omitempty"`
	OverviewTranslations []string    `json:"overviewTranslations,omitempty"`
	RemoteIds            []RemoteID  `json:"remoteIds,omitempty"`
	Score                *int64      `json:"score,omitempty"`
	Tags                 []TagOption `json:"tags,omitempty"`
	URL                  *string     `json:"url,omitempty"`
}

func (o *ListBaseRecord) GetAliases() []Alias {
	if o == nil {
		return nil
	}
	return o.Aliases
}

func (o *ListBaseRecord) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ListBaseRecord) GetImage() *string {
	if o == nil {
		return nil
	}
	return o.Image
}

func (o *ListBaseRecord) GetImageIsFallback() *bool {
	if o == nil {
		return nil
	}
	return o.ImageIsFallback
}

func (o *ListBaseRecord) GetIsOfficial() *bool {
	if o == nil {
		return nil
	}
	return o.IsOfficial
}

func (o *ListBaseRecord) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ListBaseRecord) GetNameTranslations() []string {
	if o == nil {
		return nil
	}
	return o.NameTranslations
}

func (o *ListBaseRecord) GetOverview() *string {
	if o == nil {
		return nil
	}
	return o.Overview
}

func (o *ListBaseRecord) GetOverviewTranslations() []string {
	if o == nil {
		return nil
	}
	return o.OverviewTranslations
}

func (o *ListBaseRecord) GetRemoteIds() []RemoteID {
	if o == nil {
		return nil
	}
	return o.RemoteIds
}

func (o *ListBaseRecord) GetScore() *int64 {
	if o == nil {
		return nil
	}
	return o.Score
}

func (o *ListBaseRecord) GetTags() []TagOption {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *ListBaseRecord) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}
