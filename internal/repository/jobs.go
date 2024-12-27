package repository

import (
	"Aurelia/internal/domain/models"
	"database/sql"
)

func SelectJobList(db *sql.DB, hiringType, technologyType string, incomeRange int, jobTag, usedTechnology string) ([]models.Job, error) {
	query := "SELECT job_id, company_id, hiring_type, technology_type, income_range, requirements, used_technology, application_deadline, location, remote_work, job_category FROM jobs WHERE 1=1"
	args := []interface{}{}

	if hiringType != "" {
		query += " AND hiring_type = $1"
		args = append(args, hiringType)
	}
	if technologyType != "" {
		query += " AND technology_type = $2"
		args = append(args, technologyType)
	}
	if incomeRange > 0 {
		query += " AND income_range = $3"
		args = append(args, incomeRange)
	}
	if jobTag != "" {
		query += " AND job_category = $4"
		args = append(args, jobTag)
	}
	if usedTechnology != "" {
		query += " AND used_technology = $5"
		args = append(args, usedTechnology)
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []models.Job
	for rows.Next() {
		var job models.Job
		err := rows.Scan(&job.JobID, &job.CompanyID, &job.HiringType, &job.TechnologyType, &job.IncomeRange, &job.Requirements, &job.UsedTechnology)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}
