package draftcasemodel

import (
	"github.com/golang-module/carbon/v2"
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
		return Verdict
	case "blacklist":
		return Blacklist
	case "sanction":
		return Sanction
	default:
		return 0
	}
}

const (
	Verdict CaseType = iota + 1
	Blacklist
	Sanction
)

type CaseStatus int

const (
	Deleted CaseStatus = iota
	Validated
	Draft
)

func (c CaseStatus) String() string {
	return [...]string{"deleted", "validated", "draft"}[c]
}

type SubjectTypeInt int

const (
	Individual SubjectTypeInt = iota + 1
	Company
	Organization
)

func NewSubjectType(s string) SubjectTypeInt {

	switch s {
	case "individual":
		return Individual
	case "company":
		return Company
	case "organization":
		return Organization
	default:
		return 0
	}
}
func (s SubjectTypeInt) String() string {
	return [...]string{"individual", "company", "organization"}[s-1]
}

type DraftCaseModel struct {
	ID                   ulid.ULID       `json:"id"`
	Subject              string          `json:"subject"`
	SubjectType          SubjectTypeInt  `json:"subject_type"`
	PersonInCharge       null.String     `json:"person_in_charge"`
	BenificiaryOwnership null.String     `json:"benificiary_ownership"`
	CaseDate             null.String     `json:"case_date"`
	DecisionNumber       null.String     `json:"decision_number"`
	Source               string          `json:"source"`
	Link                 string          `json:"link"`
	Nation               string          `json:"nation"`
	PunishmentStart      null.String     `json:"punishment_start"`
	PunishmentEnd        null.String     `json:"punishment_end"`
	CaseType             CaseType        `json:"type"`
	Year                 string          `json:"year"`
	Summary              string          `json:"summary"`
	SummaryFormatted     string          `json:"summary_formatted"`
	SummaryEn            string          `json:"summary_en"`
	SummaryFormattedEn   string          `json:"summary_formatted_en"`
	CreatedAt            carbon.DateTime `json:"created_at"`
	UpdatedAt            carbon.DateTime `json:"updated_at"`
}

var EmptyDraft DraftCaseModel
