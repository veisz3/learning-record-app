package usecase

import (
	"github.com/veise3/learning-record-app/internal/domain"
	"github.com/veise3/learning-record-app/internal/repository"
)

type LearningRecordUseCase interface {
	CreateLearningRecord(record *domain.LearningRecord) error
	GetLearningRecords() ([]*domain.LearningRecord, error)
	UpdateLearningRecord(record *domain.LearningRecord) error
	DeleteLearningRecord(id uint) error
}

type learningRecordUseCase struct {
	repo repository.LearningRecordRepository
}

func NewLearningRecordUseCase(repo repository.LearningRecordRepository) LearningRecordUseCase {
	return &learningRecordUseCase{repo: repo}
}

func (u *learningRecordUseCase) CreateLearningRecord(record *domain.LearningRecord) error {
	return u.repo.Create(record)
}

func (u *learningRecordUseCase) GetLearningRecords() ([]*domain.LearningRecord, error) {
	return u.repo.GetAll()
}

func (u *learningRecordUseCase) UpdateLearningRecord(record *domain.LearningRecord) error {
	return u.repo.Update(record)
}

func (u *learningRecordUseCase) DeleteLearningRecord(id uint) error {
	return u.repo.Delete(id)
}
