// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// GenreBaseRecord - base genre record
type GenreBaseRecord struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Slug *string `json:"slug,omitempty"`
}

func (o *GenreBaseRecord) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GenreBaseRecord) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GenreBaseRecord) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}