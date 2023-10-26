// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Language - language record
type Language struct {
	ID         *string `json:"id,omitempty"`
	Name       *string `json:"name,omitempty"`
	NativeName *string `json:"nativeName,omitempty"`
	ShortCode  *string `json:"shortCode,omitempty"`
}

func (o *Language) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Language) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Language) GetNativeName() *string {
	if o == nil {
		return nil
	}
	return o.NativeName
}

func (o *Language) GetShortCode() *string {
	if o == nil {
		return nil
	}
	return o.ShortCode
}
