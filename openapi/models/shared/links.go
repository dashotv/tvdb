// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Links for next, previous and current record
type Links struct {
	Next       *string `json:"next,omitempty"`
	PageSize   *int64  `json:"page_size,omitempty"`
	Prev       *string `json:"prev,omitempty"`
	Self       *string `json:"self,omitempty"`
	TotalItems *int64  `json:"total_items,omitempty"`
}

func (o *Links) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *Links) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *Links) GetPrev() *string {
	if o == nil {
		return nil
	}
	return o.Prev
}

func (o *Links) GetSelf() *string {
	if o == nil {
		return nil
	}
	return o.Self
}

func (o *Links) GetTotalItems() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalItems
}
