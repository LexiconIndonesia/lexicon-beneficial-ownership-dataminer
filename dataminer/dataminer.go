package dataminer

import (
	draftcasemodel "lexicon/lexicon-beneficial-ownership-dataminer/models/draft_model"
)

type Dataminer interface {
	MineMetadata() (draftcasemodel.DraftCaseModel, error)
}
