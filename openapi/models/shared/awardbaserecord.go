// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AwardBaseRecord - base award record
type AwardBaseRecord struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

func (o *AwardBaseRecord) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AwardBaseRecord) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
