package requests

import "go-learn/main/models"

type EntryMetricCreateRequest struct {
	MetricTypeID uint                            `json:"metric_type_id" binding:"required"`
	Values       []EntryMetricValueCreateRequest `json:"values"`
}

type EntryMetricResponse struct {
	ID           uint                       `json:"id"`
	MetricTypeID uint                       `json:"metric_type_id"`
	MetricType   *DictionaryItemResponse    `json:"metric_type,omitempty"`
	DiaryEntryID uint                       `json:"diary_entry_id"`
	Values       []EntryMetricValueResponse `json:"values,omitempty"`
}

func NewEntryMetricResponse(metric models.EntryMetric) EntryMetricResponse {
	response := EntryMetricResponse{
		ID:           metric.ID,
		MetricTypeID: metric.MetricTypeID,
		DiaryEntryID: metric.DiaryEntryID,
		Values:       NewEntryMetricValueResponses(metric.Values),
	}

	if metric.MetricType.ID != 0 {
		metricType := NewDictionaryItemResponse(metric.MetricType)
		response.MetricType = &metricType
	}

	return response
}

func NewEntryMetricResponses(metrics []models.EntryMetric) []EntryMetricResponse {
	responses := make([]EntryMetricResponse, 0, len(metrics))
	for _, metric := range metrics {
		responses = append(responses, NewEntryMetricResponse(metric))
	}

	return responses
}
