// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ContentRating - content rating record
type ContentRating struct {
	ContentType *string `json:"contentType,omitempty"`
	Country     *string `json:"country,omitempty"`
	Description *string `json:"description,omitempty"`
	FullName    *string `json:"fullName,omitempty"`
	ID          *int64  `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Order       *int64  `json:"order,omitempty"`
}

func (o *ContentRating) GetContentType() *string {
	if o == nil {
		return nil
	}
	return o.ContentType
}

func (o *ContentRating) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *ContentRating) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ContentRating) GetFullName() *string {
	if o == nil {
		return nil
	}
	return o.FullName
}

func (o *ContentRating) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ContentRating) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ContentRating) GetOrder() *int64 {
	if o == nil {
		return nil
	}
	return o.Order
}