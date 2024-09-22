package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/veise3/learning-record-app/internal/domain"
	"github.com/veise3/learning-record-app/internal/usecase"
)

type LearningRecordHandler struct {
	useCase usecase.LearningRecordUseCase
}

func NewLearningRecordHandler(useCase usecase.LearningRecordUseCase) *LearningRecordHandler {
	return &LearningRecordHandler{useCase: useCase}
}

func (h *LearningRecordHandler) CreateLearningRecord(c *gin.Context) {
	var record domain.LearningRecord
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.useCase.CreateLearningRecord(&record); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, record)
}

func (h *LearningRecordHandler) GetLearningRecords(c *gin.Context) {
	records, err := h.useCase.GetLearningRecords()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, records)
}

func (h *LearningRecordHandler) UpdateLearningRecord(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var record domain.LearningRecord
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record.ID = uint(id)
	if err := h.useCase.UpdateLearningRecord(&record); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, record)
}

func (h *LearningRecordHandler) DeleteLearningRecord(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.useCase.DeleteLearningRecord(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Record deleted successfully"})
}
