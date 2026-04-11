package database

import (
	"errors"

	"go-learn/main/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	if dsn == "" {
		return errors.New("DATABASE_URL is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.DictionaryItem{},
		&models.DiaryEntry{},
		&models.EntryMetric{},
		&models.EntryMetricValue{},
	); err != nil {
		return err
	}

	if err := dropLegacyIndexes(db); err != nil {
		return err
	}

	DB = db
	return nil
}

func dropLegacyIndexes(db *gorm.DB) error {
	indexes := []struct {
		model any
		name  string
	}{
		{model: &models.DictionaryItem{}, name: "udx_dictionary_type_label"},
		{model: &models.DiaryEntry{}, name: "idx_diary_started"},
		{model: &models.EntryMetric{}, name: "idx_metric_type"},
		{model: &models.EntryMetric{}, name: "idx_metric_entry"},
		{model: &models.EntryMetricValue{}, name: "udx_metric_unit"},
	}

	for _, index := range indexes {
		if !db.Migrator().HasIndex(index.model, index.name) {
			continue
		}

		if err := db.Migrator().DropIndex(index.model, index.name); err != nil {
			return err
		}
	}

	return nil
}
