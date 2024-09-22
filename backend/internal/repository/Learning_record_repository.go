package repository

import (
	"github.com/veise3/learning-record-app/internal/domain"
	"gorm.io/gorm"
)

type LearningRecordRepository interface {
	Create(record *domain.LearningRecord) error
	GetAll() ([]*domain.LearningRecord, error)
	Update(record *domain.LearningRecord) error
	Delete(id uint) error
}

type learningRecordRepository struct {
	db *gorm.DB
}

func NewLearningRecordRepository(db *gorm.DB) LearningRecordRepository {
	return &learningRecordRepository{db: db}
}

func (r *learningRecordRepository) Create(record *domain.LearningRecord) error {
	return r.db.Create(record).Error
}

func (r *learningRecordRepository) GetAll() ([]*domain.LearningRecord, error) {
	var records []*domain.LearningRecord
	err := r.db.Find(&records).Error
	return records, err
}

func (r *learningRecordRepository) Update(record *domain.LearningRecord) error {
	return r.db.Save(record).Error
}

func (r *learningRecordRepository) Delete(id uint) error {
	return r.db.Delete(&domain.LearningRecord{}, id).Error
}
