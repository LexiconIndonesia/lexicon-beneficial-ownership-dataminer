package draftcasemodel

import (
	"github.com/guregu/null"
	"github.com/oklog/ulid/v2"
)

type CaseType int

func (c CaseType) String() string {
	return [...]string{"verdict", "blacklist", "sanction"}[c-1]
}

func newCaseType(s string) CaseType {
	switch s {
	case "verdict":
		return verdict
	case "blacklist":
		return blacklist
	case "sanction":
		return sanction
	default:
		return 0
	}
}

const (
	verdict CaseType = iota + 1
	blacklist
	sanction
)

type CaseStatus int

const (
	deleted CaseStatus = iota
	validated
	draft
)

func (c CaseStatus) String() string {
	return [...]string{"deleted", "validated", "draft"}[c]
}

type SubjectTypeInt int

const (
	individual SubjectTypeInt = iota + 1
	company
	organization
)

func newSubjectType(s string) SubjectTypeInt {

	switch s {
	case "individual":
		return individual
	case "company":
		return company
	case "organization":
		return organization
	default:
		return 0
	}
}
func (s SubjectTypeInt) String() string {
	return [...]string{"individual", "company", "organization"}[s-1]
}

type DraftCaseModel struct {
	ID                   ulid.ULID   `json:"id"`
	Subject              string      `json:"subject"`
	SubjectType          string      `json:"subject_type"`
	PersonInCharge       null.String `json:"person_in_charge"`
	BenificiaryOwnership null.String `json:"benificiary_ownership"`
	CaseDate             null.Time   `json:"date"`
	DecisionNumber       null.String `json:"decision_number"`
	Source               string      `json:"source"`
	Link                 string      `json:"link"`
	Nation               string      `json:"nation"`
	PunishmentStart      null.Time   `json:"punishment_start"`
	PunishmentEnd        null.Time   `json:"punishment_end"`
	CaseType             string      `json:"case_type"`
	Year                 string      `json:"year"`
	Summary              string      `json:"summary"`
	SummaryFormatted     string      `json:"summary_formatted"`
	CreatedAt            null.Time   `json:"created_at"`
	UpdatedAt            null.Time   `json:"updated_at"`
}

var EmptyDraft DraftCaseModel
