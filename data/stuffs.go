package data

import (
	"time"
)

type Stuff struct {
	ID			int
	Name		string
	Description string
	CreatedOn	string
	UpdatedOn	string
	DeletedOn	string
}

func getStuffs() []*Stuff {
	return stuffList
}

var stuffList = []*Stuff{
	&Stuff{
		ID:				1,
		Name:			"iPhone",
		Description:	"Should be at least v11",
		CreatedOn: 		time.Now().UTC().String(),
		UpdatedOn: 		time.Now().UTC().String(),
	},
	&Stuff{
		ID:				2,
		Name:			"Monitor",
		Description:	"Either Samsung or LG",
		CreatedOn: 		time.Now().UTC().String(),
		UpdatedOn: 		time.Now().UTC().String(),
	},
}