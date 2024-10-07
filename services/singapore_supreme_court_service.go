package services

import (
	"context"
	"lexicon/lexicon-beneficial-ownership-dataminer/common"
	"lexicon/lexicon-beneficial-ownership-dataminer/models"
	draftcasemodel "lexicon/lexicon-beneficial-ownership-dataminer/models/draft_model"
	singaporesupremecourtmodel "lexicon/lexicon-beneficial-ownership-dataminer/models/singapore_supreme_court_model"
	"lexicon/lexicon-beneficial-ownership-dataminer/repositories"

	"github.com/rs/zerolog/log"
)

func GetSingaporeSupremeCourtExtraction() ([]models.Extraction[singaporesupremecourtmodel.Metadata], error) {
	ctx := context.Background()
	tx, err := common.CrawlerDB.Begin(ctx)
	if err != nil {
		return nil, err
	}
	extraction, err := repositories.GetSingaporeSupremeCourtExtraction(ctx, tx)
	if err != nil {
		return nil, err
	}
	tx.Commit(ctx)

	return extraction, nil

}

func SaveDraftCase(draftcasemodel draftcasemodel.DraftCaseModel) error {
	ctx := context.Background()
	tx, err := common.BeneficialDB.Begin(ctx)
	if err != nil {
		return err
	}
	err = repositories.SaveDraftCase(ctx, tx, draftcasemodel)
	if err != nil {
		log.Error().Err(err).Msg("Error saving draft case")
		return err
	}
	tx.Commit(ctx)
	return nil
}
