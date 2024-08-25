package dataminer

import draftcasemodel "lexicon/lexicon-beneficial-ownership-dataminer/models/draft_model"

type IndonesiaSupremeCourtDataminer struct {
}

func NewIndonesiaSupremeCourtDataminer() Dataminer {
	return IndonesiaSupremeCourtDataminer{}
}

func (d IndonesiaSupremeCourtDataminer) MineMetadata() (draftcasemodel.DraftCaseModel, error) {
	return draftcasemodel.EmptyDraft, nil
}
