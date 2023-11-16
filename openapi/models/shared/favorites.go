// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Favorites - User favorites record
type Favorites struct {
	Artwork  []int64 `json:"artwork,omitempty"`
	Episodes []int64 `json:"episodes,omitempty"`
	Lists    []int64 `json:"lists,omitempty"`
	Movies   []int64 `json:"movies,omitempty"`
	People   []int64 `json:"people,omitempty"`
	Series   []int64 `json:"series,omitempty"`
}

func (o *Favorites) GetArtwork() []int64 {
	if o == nil {
		return nil
	}
	return o.Artwork
}

func (o *Favorites) GetEpisodes() []int64 {
	if o == nil {
		return nil
	}
	return o.Episodes
}

func (o *Favorites) GetLists() []int64 {
	if o == nil {
		return nil
	}
	return o.Lists
}

func (o *Favorites) GetMovies() []int64 {
	if o == nil {
		return nil
	}
	return o.Movies
}

func (o *Favorites) GetPeople() []int64 {
	if o == nil {
		return nil
	}
	return o.People
}

func (o *Favorites) GetSeries() []int64 {
	if o == nil {
		return nil
	}
	return o.Series
}
