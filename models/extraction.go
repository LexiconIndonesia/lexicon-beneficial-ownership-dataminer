package models

import (
	indonesiasupremecourtmodel "lexicon/lexicon-beneficial-ownership-dataminer/models/indonesia_supreme_court_model"

	"github.com/golang-module/carbon/v2"
	"github.com/guregu/null"
)

type Extraction[T any] struct {
	Id            string
	UrlFrontierId string
	SiteContent   null.String
	ArtifactLink  null.String
	RawPageLink   null.String
	Metadata      T
	CreatedAt     carbon.DateTime
	UpdatedAt     carbon.DateTime
	Language      string
}

func NewIndonesiaSupremeCourtExtraction() Extraction[indonesiasupremecourtmodel.Metadata] {
	return Extraction[indonesiasupremecourtmodel.Metadata]{
		Id:            "",
		UrlFrontierId: "",
		SiteContent:   null.StringFrom(""),
		ArtifactLink:  null.StringFrom(""),
		RawPageLink:   null.StringFrom(""),
		Metadata:      indonesiasupremecourtmodel.EmptyMetadata,
		CreatedAt:     carbon.Now().ToDateTimeStruct(),
		UpdatedAt:     carbon.Now().ToDateTimeStruct(),
		Language:      "id",
	}
}
