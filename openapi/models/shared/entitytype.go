// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// EntityType - Entity Type record
type EntityType struct {
	HasSpecials *bool   `json:"hasSpecials,omitempty"`
	ID          *int64  `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
}

func (o *EntityType) GetHasSpecials() *bool {
	if o == nil {
		return nil
	}
	return o.HasSpecials
}

func (o *EntityType) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *EntityType) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}