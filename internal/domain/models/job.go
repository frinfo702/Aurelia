package models

type Job struct {
	JobID          int    `json:"job_id"`
	CompanyID      int    `json:"company_id"`
	HiringType     string `json:"hiring_type"`
	TechnologyType string `json:"technology_type"`
	IncomeRange    int    `json:"income_range"`
	JobTag         string `json:"job_tag"`
	Requirements   string `json:"requirements"`
	UsedTechnology string `json:"used_technology"`
}
