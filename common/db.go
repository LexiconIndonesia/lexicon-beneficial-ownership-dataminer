package common

import (
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	CrawlerDB    *pgxpool.Pool
	BeneficialDB *pgxpool.Pool
)

func SetCrawlerDB(newPool *pgxpool.Pool) error {

	if newPool == nil {
		return errors.New("cannot assign nil database")
	}
	CrawlerDB = newPool
	return nil
}

func SetBeneficialDB(newPool *pgxpool.Pool) error {

	if newPool == nil {
		return errors.New("cannot assign nil database")
	}
	BeneficialDB = newPool
	return nil
}
