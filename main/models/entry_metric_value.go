package models

import "errors"

type EntryMetricValue struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	EntryMetricID uint           `gorm:"column:entry_metric_id;not null" json:"entry_metric_id"`
	EntryMetric   *EntryMetric   `gorm:"foreignKey:EntryMetricID" json:"-"`
	UnitID        uint           `gorm:"column:unit_id;not null" json:"unit_id"`
	Unit          DictionaryItem `gorm:"foreignKey:UnitID" json:"unit,omitempty"`
	Value         int            `gorm:"not null" json:"value"`
}

func (EntryMetricValue) TableName() string {
	return "entry_metric_value"
}

func NewEntryMetricValue(unitID uint, value int) (*EntryMetricValue, error) {
	if unitID == 0 {
		return nil, errors.New("unit is required")
	}

	if value <= 0 {
		return nil, errors.New("value must be positive")
	}

	return &EntryMetricValue{
		UnitID: unitID,
		Value:  value,
	}, nil
}

func (v *EntryMetricValue) ChangeValue(newValue int) error {
	if newValue <= 0 {
		return errors.New("value must be positive")
	}

	v.Value = newValue
	return nil
}
