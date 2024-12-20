package models

import (
	"time"
)

type Inquiry struct {
	CompanyName       string
	CompanyOverview   string
	WorkingPepleID    int
	CultureAndBenefit string
	EstablishDate     time.Time
	CompanyWebsite    string
	CompanyLocations  string
	CompanySize       string
	TotalRaised       string
	CompanyType       string
	CompanyMarkets    string
	IsAuthorized      bool
}
