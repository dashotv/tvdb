// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Trailer - trailer record
type Trailer struct {
	ID       *int64  `json:"id,omitempty"`
	Language *string `json:"language,omitempty"`
	Name     *string `json:"name,omitempty"`
	Runtime  *int64  `json:"runtime,omitempty"`
	URL      *string `json:"url,omitempty"`
}

func (o *Trailer) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Trailer) GetLanguage() *string {
	if o == nil {
		return nil
	}
	return o.Language
}

func (o *Trailer) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Trailer) GetRuntime() *int64 {
	if o == nil {
		return nil
	}
	return o.Runtime
}

func (o *Trailer) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}
