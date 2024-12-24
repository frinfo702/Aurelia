package models

import "time"

type Company struct {
	CompanyID         int       `json:"company_id"`
	CompanyName       string    `json:"company_name"`
	CompanyOverview   string    `json:"company_overview"`
	WorkingPeopleID   int       `json:"working_people_id"`
	CultureAndBenefit string    `json:"culture_and_benefit"`
	EstablishDate     time.Time `json:"establish_date"`
	CompanyWebsite    string    `json:"company_website"`
	CompanyLocations  string    `json:"company_locations"`
	CompanySize       string    `json:"company_size"`
	TotalRaised       string    `json:"total_raised"`
	CompanyType       string    `json:"company_type"`
	CompanyMarkets    string    `json:"company_markets"`
	IsAuthorized      bool      `json:"is_authorized"`
}
