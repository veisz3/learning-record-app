package domain

import "time"

type LearningRecord struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Content   string    `json:"content"`
	Duration  int       `json:"duration"`
	CreatedAt time.Time `json:"created_at"`
}
