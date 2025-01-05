package usecase

import (
	"Aurelia/internal/domain/models"
	"Aurelia/internal/domain/repository/postgresql"
	"errors"
)

type jobUsecase struct {
	jobRepo postgresql.JobRepository
}

func NewJobUsecase(r postgresql.JobRepository) JobUsecase {
	return &jobUsecase{jobRepo: r}
}

// GetJobs は全求人取得
func (uc *jobUsecase) GetJobs() ([]models.Job, error) {
	return uc.jobRepo.FindAll()
}

// GetJobByID はID指定で求人を取得
func (uc *jobUsecase) GetJobByID(id int) (*models.Job, error) {
	return uc.jobRepo.FindByID(id)
}

// CreateJob は新規求人を作成。バリデーション例を追加してみた
func (uc *jobUsecase) CreateJob(job *models.Job) error {
	if job.IncomeRange < 0 {
		return errors.New("income_range must be >= 0")
	}
	if job.HiringType == "" {
		return errors.New("hiring_type must not be empty")
	}
	// ほかのバリデーション…

	job.ApplicationStatus = "pending" // 例: 新規求人はpendingでスタート
	return uc.jobRepo.Insert(job)
}

// DeleteJob はID指定で求人を削除
func (uc *jobUsecase) DeleteJob(id int) error {
	return uc.jobRepo.Delete(id)
}
