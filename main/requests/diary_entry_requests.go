package requests

import (
	"time"

	"go-learn/main/models"
)

type DiaryEntryCreateRequest struct {
	WhenStarted time.Time                  `json:"when_started" binding:"required"`
	WhenEnded   time.Time                  `json:"when_ended" binding:"required"`
	Mood        *int16                     `json:"mood"`
	Description string                     `json:"description" binding:"required"`
	Metrics     []EntryMetricCreateRequest `json:"metrics"`
}

type DiaryEntryResponse struct {
	ID          uint                  `json:"id"`
	UserID      uint                  `json:"user_id"`
	User        *UserResponse         `json:"user,omitempty"`
	Metrics     []EntryMetricResponse `json:"metrics,omitempty"`
	WhenStarted time.Time             `json:"when_started"`
	WhenEnded   time.Time             `json:"when_ended"`
	Duration    int                   `json:"duration"`
	Mood        *int16                `json:"mood,omitempty"`
	Description string                `json:"description"`
	CreatedAt   time.Time             `json:"created_at"`
	UpdatedAt   time.Time             `json:"updated_at"`
}

func NewDiaryEntryResponse(entry models.DiaryEntry) DiaryEntryResponse {
	response := DiaryEntryResponse{
		ID:          entry.ID,
		UserID:      entry.UserID,
		Metrics:     NewEntryMetricResponses(entry.Metrics),
		WhenStarted: entry.WhenStarted,
		WhenEnded:   entry.WhenEnded,
		Duration:    entry.Duration,
		Mood:        entry.Mood,
		Description: entry.Description,
		CreatedAt:   entry.CreatedAt,
		UpdatedAt:   entry.UpdatedAt,
	}

	if entry.User != nil {
		user := NewUserResponse(*entry.User)
		response.User = &user
	}

	return response
}

func NewDiaryEntryResponses(entries []models.DiaryEntry) []DiaryEntryResponse {
	responses := make([]DiaryEntryResponse, 0, len(entries))
	for _, entry := range entries {
		responses = append(responses, NewDiaryEntryResponse(entry))
	}

	return responses
}
