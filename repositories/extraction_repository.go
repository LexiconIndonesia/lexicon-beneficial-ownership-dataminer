package repositories

import (
	"context"
	"encoding/json"
	"lexicon/lexicon-beneficial-ownership-dataminer/common"
	"lexicon/lexicon-beneficial-ownership-dataminer/models"
	singaporesupremecourtmodel "lexicon/lexicon-beneficial-ownership-dataminer/models/singapore_supreme_court_model"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

func GetSingaporeSupremeCourtExtraction(context context.Context, tx pgx.Tx) ([]models.Extraction[singaporesupremecourtmodel.Metadata], error) {
	sql := `select e.id, e.url_frontier_id, e.site_content, e.artifact_link, e.raw_page_link, e.metadata, e.created_at, e.updated_at, e.language from extraction e
inner join url_frontier uf
on e.url_frontier_id  = uf.id
where uf.crawler  = $1`
	var extraction []models.Extraction[singaporesupremecourtmodel.Metadata]
	row, err := tx.Query(context, sql, common.SingaporeSupremeCourtCrawlerName)

	if err != nil {
		return nil, err
	}

	for row.Next() {
		var e models.Extraction[singaporesupremecourtmodel.Metadata]
		var metadata string

		err := row.Scan(&e.Id, &e.UrlFrontierId, &e.SiteContent, &e.ArtifactLink, &e.RawPageLink, &metadata, &e.CreatedAt, &e.UpdatedAt, &e.Language)
		if err != nil {
			log.Error().Err(err).Msg("Error scanning row")
			return nil, err
		}

		err = json.Unmarshal([]byte(metadata), &e.Metadata)

		if err != nil {
			log.Error().Err(err).Msg("Error unmarshalling metadata")

			return nil, err
		}
		extraction = append(extraction, e)
	}
	return extraction, nil

}
