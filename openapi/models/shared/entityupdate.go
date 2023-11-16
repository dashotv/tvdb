// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// EntityUpdate - entity update record
type EntityUpdate struct {
	EntityType        *string `json:"entityType,omitempty"`
	ExtraInfo         *string `json:"extraInfo,omitempty"`
	MergeToEntityType *string `json:"mergeToEntityType,omitempty"`
	MergeToID         *int64  `json:"mergeToId,omitempty"`
	Method            *string `json:"method,omitempty"`
	MethodInt         *int64  `json:"methodInt,omitempty"`
	RecordID          *int64  `json:"recordId,omitempty"`
	RecordType        *string `json:"recordType,omitempty"`
	// Only present for episodes records
	SeriesID  *int64 `json:"seriesId,omitempty"`
	TimeStamp *int64 `json:"timeStamp,omitempty"`
	UserID    *int64 `json:"userId,omitempty"`
}

func (o *EntityUpdate) GetEntityType() *string {
	if o == nil {
		return nil
	}
	return o.EntityType
}

func (o *EntityUpdate) GetExtraInfo() *string {
	if o == nil {
		return nil
	}
	return o.ExtraInfo
}

func (o *EntityUpdate) GetMergeToEntityType() *string {
	if o == nil {
		return nil
	}
	return o.MergeToEntityType
}

func (o *EntityUpdate) GetMergeToID() *int64 {
	if o == nil {
		return nil
	}
	return o.MergeToID
}

func (o *EntityUpdate) GetMethod() *string {
	if o == nil {
		return nil
	}
	return o.Method
}

func (o *EntityUpdate) GetMethodInt() *int64 {
	if o == nil {
		return nil
	}
	return o.MethodInt
}

func (o *EntityUpdate) GetRecordID() *int64 {
	if o == nil {
		return nil
	}
	return o.RecordID
}

func (o *EntityUpdate) GetRecordType() *string {
	if o == nil {
		return nil
	}
	return o.RecordType
}

func (o *EntityUpdate) GetSeriesID() *int64 {
	if o == nil {
		return nil
	}
	return o.SeriesID
}

func (o *EntityUpdate) GetTimeStamp() *int64 {
	if o == nil {
		return nil
	}
	return o.TimeStamp
}

func (o *EntityUpdate) GetUserID() *int64 {
	if o == nil {
		return nil
	}
	return o.UserID
}
