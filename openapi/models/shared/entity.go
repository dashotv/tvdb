// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Entity record
type Entity struct {
	MovieID  *int64 `json:"movieId,omitempty"`
	Order    *int64 `json:"order,omitempty"`
	SeriesID *int64 `json:"seriesId,omitempty"`
}

func (o *Entity) GetMovieID() *int64 {
	if o == nil {
		return nil
	}
	return o.MovieID
}

func (o *Entity) GetOrder() *int64 {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *Entity) GetSeriesID() *int64 {
	if o == nil {
		return nil
	}
	return o.SeriesID
}
