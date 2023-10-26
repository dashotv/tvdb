// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Inspiration - Movie inspiration record
type Inspiration struct {
	ID       *int64  `json:"id,omitempty"`
	Type     *string `json:"type,omitempty"`
	TypeName *string `json:"type_name,omitempty"`
	URL      *string `json:"url,omitempty"`
}

func (o *Inspiration) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Inspiration) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Inspiration) GetTypeName() *string {
	if o == nil {
		return nil
	}
	return o.TypeName
}

func (o *Inspiration) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}