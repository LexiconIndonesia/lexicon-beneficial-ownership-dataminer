package dataminer

import (
	"fmt"
	"lexicon/lexicon-beneficial-ownership-dataminer/common"
	"lexicon/lexicon-beneficial-ownership-dataminer/models"
	draftcasemodel "lexicon/lexicon-beneficial-ownership-dataminer/models/draft_model"
	singaporesupremecourtmodel "lexicon/lexicon-beneficial-ownership-dataminer/models/singapore_supreme_court_model"
	"lexicon/lexicon-beneficial-ownership-dataminer/services"

	"github.com/golang-module/carbon/v2"
	"github.com/guregu/null"
	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
)

type SingaporeSupremeCourtDataminer struct {
}

// Start implements Dataminer.
func (s SingaporeSupremeCourtDataminer) Start() {
	log.Info().Msg("Starting Singapore Supreme Court Dataminer")

	log.Info().Msg("Getting current extractions")
	extract, err := s.getCurrentExtractions()

	log.Info().Msg(fmt.Sprintf("Extraction length: %d", len(extract)))
	log.Info().Msg("Mine metadata")

	if err != nil {
		log.Error().Err(err).Msg("Error getting current extractions")
	}

	s.mineMetadata(&extract)

	log.Info().Msg("Finished Singapore Supreme Court Dataminer")

}

// getCurrentExtractions implements Dataminer.
func (s SingaporeSupremeCourtDataminer) getCurrentExtractions() ([]models.Extraction[singaporesupremecourtmodel.Metadata], error) {
	return services.GetSingaporeSupremeCourtExtraction()
}

// mineMetadata implements Dataminer.
func (s SingaporeSupremeCourtDataminer) mineMetadata(e *[]models.Extraction[singaporesupremecourtmodel.Metadata]) error {

	for _, extraction := range *e {
		log.Info().Msg("Mine metadata of extraction: " + extraction.Metadata.Title)
		m := draftcasemodel.DraftCaseModel{
			ID:                 ulid.Make(),
			Subject:            extraction.Metadata.Defendant,
			SubjectType:        draftcasemodel.Individual,
			CaseDate:           null.StringFrom(extraction.Metadata.DecisionDate),
			DecisionNumber:     null.StringFrom(extraction.Metadata.Number),
			Source:             common.SingaporeSupremeCourt,
			Link:               extraction.RawPageLink.String,
			Nation:             "Singapore",
			CaseType:           draftcasemodel.Verdict,
			Year:               extraction.Metadata.Year,
			Summary:            extraction.Metadata.Verdict,
			SummaryFormatted:   extraction.Metadata.VerdictMarkdown,
			SummaryEn:          extraction.Metadata.Verdict,
			SummaryFormattedEn: extraction.Metadata.VerdictMarkdown,

			CreatedAt: carbon.Now().ToDateTimeStruct(),
			UpdatedAt: carbon.Now().ToDateTimeStruct(),
		}

		services.SaveDraftCase(m)

	}
	return nil
}

func NewSingaporeSupremeCourtDataminer() Dataminer[singaporesupremecourtmodel.Metadata] {
	return SingaporeSupremeCourtDataminer{}
}
