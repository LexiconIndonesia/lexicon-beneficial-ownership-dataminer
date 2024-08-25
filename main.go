package main

import (
	"context"
	"lexicon/lexicon-beneficial-ownership-dataminer/dataminer"

	"github.com/golang-module/carbon/v2"

	"github.com/rs/zerolog/log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	// INITIATE CONFIGURATION
	err := godotenv.Load()
	if err != nil {
		log.Error().Err(err).Msg("Error loading .env file")
	}
	cfg := defaultConfig()
	cfg.loadFromEnv()

	ctx := context.Background()

	carbon.SetDefault(carbon.Default{
		Layout:       carbon.ISO8601Layout,
		Timezone:     carbon.UTC,
		WeekStartsAt: carbon.Monday,
		Locale:       "en",
	})

	// INITIATE DATABASES
	// PGSQL
	crawlerDb, err := pgxpool.New(ctx, cfg.CrawlerDB.ConnStr())

	if err != nil {
		log.Error().Err(err).Msg("Unable to connect to PGSQL Database")
	}
	defer crawlerDb.Close()

	dataminer.SetCrawlerDB(crawlerDb)

	beneficialDb, err := pgxpool.New(ctx, cfg.BeneficialDB.ConnStr())
	if err != nil {
		log.Error().Err(err).Msg("Unable to connect to PGSQL Database")
	}
	defer beneficialDb.Close()

	dataminer.SetBeneficialDB(beneficialDb)

}
