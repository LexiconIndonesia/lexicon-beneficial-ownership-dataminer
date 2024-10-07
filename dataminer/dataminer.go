package dataminer

import (
	"lexicon/lexicon-beneficial-ownership-dataminer/models"
)

type Dataminer[T any] interface {
	Start()
	mineMetadata(e *[]models.Extraction[T]) error
	getCurrentExtractions() ([]models.Extraction[T], error)
}
