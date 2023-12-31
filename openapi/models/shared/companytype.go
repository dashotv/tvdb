// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CompanyType - A company type record
type CompanyType struct {
	CompanyTypeID   *int64  `json:"companyTypeId,omitempty"`
	CompanyTypeName *string `json:"companyTypeName,omitempty"`
}

func (o *CompanyType) GetCompanyTypeID() *int64 {
	if o == nil {
		return nil
	}
	return o.CompanyTypeID
}

func (o *CompanyType) GetCompanyTypeName() *string {
	if o == nil {
		return nil
	}
	return o.CompanyTypeName
}
