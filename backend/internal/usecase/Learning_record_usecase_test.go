package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/veise3/learning-record-app/internal/domain"
)

type MockLearningRecordRepository struct {
	mock.Mock
}

func (m *MockLearningRecordRepository) Create(record *domain.LearningRecord) error {
	args := m.Called(record)
	return args.Error(0)
}

func (m *MockLearningRecordRepository) GetAll() ([]*domain.LearningRecord, error) {
	args := m.Called()
	return args.Get(0).([]*domain.LearningRecord), args.Error(1)
}

func (m *MockLearningRecordRepository) Update(record *domain.LearningRecord) error {
	args := m.Called(record)
	return args.Error(0)
}

func (m *MockLearningRecordRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func Test_learningRecordUseCase_GetLearningRecords(t *testing.T) {
	mockRepo := new(MockLearningRecordRepository)
	useCase := NewLearningRecordUseCase(mockRepo)

	record := &domain.LearningRecord{Content: "Test Content", Duration: 30}

	// 成功ケース
	mockRepo.On("Create", record).Return(nil).Once()
	err := useCase.CreateLearningRecord(record)
	assert.NoError(t, err)

	// エラーケース
	expectedError := errors.New("creation error")
	mockRepo.On("Create", record).Return(expectedError).Once()
	err = useCase.CreateLearningRecord(record)
	assert.Equal(t, expectedError, err)

	mockRepo.AssertExpectations(t)
}

func TestGetLearningRecords(t *testing.T) {
	mockRepo := new(MockLearningRecordRepository)
	useCase := NewLearningRecordUseCase(mockRepo)

	expectedRecords := []*domain.LearningRecord{
		{ID: 1, Content: "Test 1", Duration: 30},
		{ID: 2, Content: "Test 2", Duration: 45},
	}

	// 成功ケース
	mockRepo.On("GetAll").Return(expectedRecords, nil).Once()
	records, err := useCase.GetLearningRecords()
	assert.NoError(t, err)
	assert.Equal(t, expectedRecords, records)

	// 失敗ケース
	expectedError := errors.New("retrieval error")
	mockRepo.On("GetAll").Return(([]*domain.LearningRecord)(nil), expectedError).Once()
	records, err = useCase.GetLearningRecords()
	assert.Equal(t, expectedError, err)
	assert.Nil(t, records)

	mockRepo.AssertExpectations(t)
}

func TestUpdateLearningRecord(t *testing.T) {
	mockRepo := new(MockLearningRecordRepository)
	useCase := NewLearningRecordUseCase(mockRepo)

	record := &domain.LearningRecord{ID: 1, Content: "Updated Content", Duration: 60}

	// 成功ケース
	mockRepo.On("Update", record).Return(nil).Once()
	err := useCase.UpdateLearningRecord(record)
	assert.NoError(t, err)

	// エラーケース
	expectedError := errors.New("update error")
	mockRepo.On("Update", record).Return(expectedError).Once()
	err = useCase.UpdateLearningRecord(record)
	assert.Equal(t, expectedError, err)

	mockRepo.AssertExpectations(t)
}

func TestDeleteLearningRecord(t *testing.T) {
	mockRepo := new(MockLearningRecordRepository)
	useCase := NewLearningRecordUseCase(mockRepo)

	id := uint(1)

	// 成功ケース
	mockRepo.On("Delete", id).Return(nil).Once()
	err := useCase.DeleteLearningRecord(id)
	assert.NoError(t, err)

	// エラーケース
	expectedError := errors.New("deletion error")
	mockRepo.On("Delete", id).Return(expectedError).Once()
	err = useCase.DeleteLearningRecord(id)
	assert.Equal(t, expectedError, err)

	mockRepo.AssertExpectations(t)
}
