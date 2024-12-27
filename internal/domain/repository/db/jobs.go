package db

import (
	"Aurelia/internal/domain/models"
	"database/sql"
	"log"
)

type PostgresJobRepository struct {
	db *sql.DB
}

func NewPostgresJobRepository(db *sql.DB) *PostgresJobRepository {
	return &PostgresJobRepository{db: db}
}

func (r *PostgresJobRepository) FindAll() ([]models.Job, error) {
	query := "SELECT job_id, company_id, hiring_type, technology_type, income_range, job_tag, used_technology FROM jobs"
	rows, err := r.db.Query(query)
	if err != nil {
		log.Println("Error while querying jobs", err)
		return nil, err
	}
	defer rows.Close()

	var jobs []models.Job
	for rows.Next() {
		var job models.Job
		err := rows.Scan(&job.JobID, &job.CompanyID, &job.HiringType, &job.TechnologyType, &job.IncomeRange, &job.JobTag, &job.UsedTechnology)
		if err != nil {
			log.Println("Error while scanning rows", err)
			return nil, err
		}
		jobs = append(jobs, job)
	}

	return jobs, nil
}

func (r *PostgresJobRepository) FindByID(id int) (*models.Job, error) {
	var job models.Job
	query := "SELECT job_id, company_id, hiring_type, technology_type, income_range, job_tag, used_technology FROM jobs WHERE job_id = $1"
	row := r.db.QueryRow(query, id)
	err := row.Scan(&job.JobID, &job.CompanyID, &job.HiringType, &job.TechnologyType, &job.IncomeRange, &job.JobTag, &job.UsedTechnology)
	if err != nil {
		log.Println("Error while scanning row: ", err)
		return nil, err
	}
	return &job, nil
}

func (r *PostgresJobRepository) Save(job *models.Job) error {
	query := "INSERT INTO jobs (company_id, hiring_type, technology_type, income_range, job_tag, used_technology) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := r.db.Exec(query, job.CompanyID, job.HiringType, job.TechnologyType, job.IncomeRange, job.JobTag, job.UsedTechnology)
	if err != nil {
		log.Println("Error while inserting job", err)
		return err
	}
	return nil
}

func (r *PostgresJobRepository) Delete(id int) error {
	query := "DELETE FROM jobs WHERE job_id = $1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Println("Error while deleting job", err)
		return err
	}
	return nil
}
