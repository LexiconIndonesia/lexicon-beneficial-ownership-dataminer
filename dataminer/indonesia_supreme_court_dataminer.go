package dataminer

import (
	"lexicon/lexicon-beneficial-ownership-dataminer/models"
	indonesiasupremecourtmodel "lexicon/lexicon-beneficial-ownership-dataminer/models/indonesia_supreme_court_model"
)

type IndonesiaSupremeCourtDataminer struct {
}

// Start implements Dataminer.
func (i IndonesiaSupremeCourtDataminer) Start() {
	panic("unimplemented")
}

// getCurrentExtractions implements Dataminer.
func (i IndonesiaSupremeCourtDataminer) getCurrentExtractions() ([]models.Extraction[indonesiasupremecourtmodel.Metadata], error) {
	panic("unimplemented")
}

// mineMetadata implements Dataminer.
func (i IndonesiaSupremeCourtDataminer) mineMetadata(e *[]models.Extraction[indonesiasupremecourtmodel.Metadata]) error {
	panic("unimplemented")
}

func NewIndonesiaSupremeCourtDataminer() Dataminer[indonesiasupremecourtmodel.Metadata] {
	return IndonesiaSupremeCourtDataminer{}
}
