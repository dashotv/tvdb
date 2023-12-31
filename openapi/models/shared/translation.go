// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Translation - translation record
type Translation struct {
	Aliases   []string `json:"aliases,omitempty"`
	IsAlias   *bool    `json:"isAlias,omitempty"`
	IsPrimary *bool    `json:"isPrimary,omitempty"`
	Language  *string  `json:"language,omitempty"`
	Name      *string  `json:"name,omitempty"`
	Overview  *string  `json:"overview,omitempty"`
	// Only populated for movie translations.  We disallow taglines without a title.
	Tagline *string `json:"tagline,omitempty"`
}

func (o *Translation) GetAliases() []string {
	if o == nil {
		return nil
	}
	return o.Aliases
}

func (o *Translation) GetIsAlias() *bool {
	if o == nil {
		return nil
	}
	return o.IsAlias
}

func (o *Translation) GetIsPrimary() *bool {
	if o == nil {
		return nil
	}
	return o.IsPrimary
}

func (o *Translation) GetLanguage() *string {
	if o == nil {
		return nil
	}
	return o.Language
}

func (o *Translation) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Translation) GetOverview() *string {
	if o == nil {
		return nil
	}
	return o.Overview
}

func (o *Translation) GetTagline() *string {
	if o == nil {
		return nil
	}
	return o.Tagline
}
