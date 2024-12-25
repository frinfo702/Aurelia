package db

import (
	"Aurelia/internal/models"
	"database/sql"
)

func SelectJobList(db *sql.DB, hiringType int, technologyType string, incomeRange int, jobTag string, usedTechnology string) ([]models.Job, error) {
	const query = `
	SELECT job_id, company_id, hiring_type, technology_type, income_range, 
		   job_tag, requirements, used_technology
	FROM jobs
	WHERE 
		($1 = 0 OR hiring_type = $1)
		AND ($2 = '' OR technology_type = $2)
		AND ($3 = 0 OR income_range = $3)
		AND ($4 = '' OR job_tag = $4)
		AND ($5 = '' OR used_technology = $5)
	`

	rows, err := db.Query(query, hiringType, technologyType, incomeRange, jobTag, usedTechnology)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var jobList []models.Job
	for rows.Next() {
		var job models.Job
		err := rows.Scan(
			&job.JobID,
			&job.CompanyID,
			&job.HiringType,
			&job.TechnologyType,
			&job.IncomeRange,
			&job.JobTag,
			&job.Requirements,
			&job.UsedTechnology,
		)
		if err != nil {
			return []models.Job{}, err
		}
		jobList = append(jobList, job)
	}
	return jobList, nil
}
