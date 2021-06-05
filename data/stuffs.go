package data

import (
	"encoding/json"
	"io"
)

type Stuff struct {
	ID			int `json:"id"`
	Name		string `json:"name"`
	Description string `json:"description"`
}

type Stuffs []*Stuff

func (s *Stuffs) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(s)
}

func GetStuffs() []*Stuff {
	return stuffList
}

var stuffList = []*Stuff{
	&Stuff{
		ID:				1,
		Name:			"iPhone",
		Description:	"Should be at least v11",
	},
	&Stuff{
		ID:				2,
		Name:			"Monitor",
		Description:	"Either Samsung or LG",
	},
}