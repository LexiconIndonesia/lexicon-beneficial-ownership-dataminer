package repositories

import (
	"context"
	draftcasemodel "lexicon/lexicon-beneficial-ownership-dataminer/models/draft_model"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

func SaveDraftCase(context context.Context, tx pgx.Tx, d draftcasemodel.DraftCaseModel) error {
	sql := `INSERT INTO draft_cases
(id, subject, subject_type, person_in_charge, benificiary_ownership, case_date, decision_number, "source", link, nation, punishment_start, punishment_end, "type", "year", summary, summary_formatted, created_at, updated_at, summary_en, summary_formatted_en) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20);`
	_, err := tx.Exec(context, sql, d.ID.String(), d.Subject, d.SubjectType, d.PersonInCharge, d.BenificiaryOwnership, d.CaseDate, d.DecisionNumber, d.Source, d.Link, d.Nation, d.PunishmentStart, d.PunishmentEnd, d.CaseType, d.Year, d.Summary, d.SummaryFormatted, d.CreatedAt, d.UpdatedAt, d.SummaryEn, d.SummaryFormattedEn)

	if err != nil {
		log.Err(err).Msg("Error saving draft case")
		return err
	}

	return nil

}
