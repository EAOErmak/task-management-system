package requests

import "go-learn/main/models"

type EntryMetricValueCreateRequest struct {
	UnitID uint `json:"unit_id" binding:"required"`
	Value  int  `json:"value" binding:"required"`
}

type EntryMetricValueResponse struct {
	ID            uint                    `json:"id"`
	EntryMetricID uint                    `json:"entry_metric_id"`
	UnitID        uint                    `json:"unit_id"`
	Unit          *DictionaryItemResponse `json:"unit,omitempty"`
	Value         int                     `json:"value"`
}

func NewEntryMetricValueResponse(value models.EntryMetricValue) EntryMetricValueResponse {
	response := EntryMetricValueResponse{
		ID:            value.ID,
		EntryMetricID: value.EntryMetricID,
		UnitID:        value.UnitID,
		Value:         value.Value,
	}

	if value.Unit.ID != 0 {
		unit := NewDictionaryItemResponse(value.Unit)
		response.Unit = &unit
	}

	return response
}

func NewEntryMetricValueResponses(values []models.EntryMetricValue) []EntryMetricValueResponse {
	responses := make([]EntryMetricValueResponse, 0, len(values))
	for _, value := range values {
		responses = append(responses, NewEntryMetricValueResponse(value))
	}

	return responses
}
