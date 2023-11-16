// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ParentCompany - A parent company record
type ParentCompany struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	// A company relationship
	Relation *CompanyRelationShip `json:"relation,omitempty"`
}

func (o *ParentCompany) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ParentCompany) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ParentCompany) GetRelation() *CompanyRelationShip {
	if o == nil {
		return nil
	}
	return o.Relation
}
