package handlers

import (
	"go-learn/main/models"
	"go-learn/main/requests"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllMineDiaryEntries(c *gin.Context) {
	claims, ok := currentAuthClaims(c)
	if !ok {
		return
	}

	db := diaryDB(c)
	if db == nil {
		return
	}

	query := preloadDiaryEntry(db).Where("user_id = ?", claims.UserID)

	var entries []models.DiaryEntry
	if err := query.Find(&entries).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, requests.NewDiaryEntryResponses(entries))
}

func GetAllDiaryEntriesForAllUsers(c *gin.Context) {
	db := diaryDB(c)
	if db == nil {
		return
	}

	var entries []models.DiaryEntry
	if err := preloadDiaryEntry(db).Find(&entries).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, requests.NewDiaryEntryResponses(entries))
}

func CreateDiaryEntry(c *gin.Context) {
	var req requests.DiaryEntryCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	claims, ok := currentAuthClaims(c)
	if !ok {
		return
	}

	db := diaryDB(c)
	if db == nil {
		return
	}

	entry, err := models.NewDiaryEntry(req.WhenStarted, req.WhenEnded, req.Mood, req.Description)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entry.UserID = claims.UserID

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(entry).Error; err != nil {
			return err
		}

		if err := createMetricRecords(tx, entry.ID, req.Metrics); err != nil {
			return err
		}

		return preloadDiaryEntry(tx).First(entry, entry.ID).Error
	}); err != nil {
		writeDiaryError(c, err)
		return
	}

	c.JSON(http.StatusCreated, requests.NewDiaryEntryResponse(*entry))
}

func GetDiaryEntryByID(c *gin.Context) {
	entryID, ok := parseUintParam(c, "id")
	if !ok {
		return
	}

	claims, ok := currentAuthClaims(c)
	if !ok {
		return
	}

	db := diaryDB(c)
	if db == nil {
		return
	}

	var entry models.DiaryEntry
	if err := preloadDiaryEntry(db).Where("user_id = ?", claims.UserID).First(&entry, entryID).Error; err != nil {
		writeDiaryError(c, err)
		return
	}

	c.JSON(http.StatusOK, requests.NewDiaryEntryResponse(entry))
}

func UpdateDiaryEntry(c *gin.Context) {
	entryID, ok := parseUintParam(c, "id")
	if !ok {
		return
	}

	claims, ok := currentAuthClaims(c)
	if !ok {
		return
	}

	var req requests.DiaryEntryCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	db := diaryDB(c)
	if db == nil {
		return
	}

	var entry models.DiaryEntry
	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ?", claims.UserID).First(&entry, entryID).Error; err != nil {
			return err
		}

		if err := entry.UpdateTime(req.WhenStarted, req.WhenEnded); err != nil {
			return err
		}

		if err := entry.UpdateMood(req.Mood); err != nil {
			return err
		}

		if err := entry.UpdateDescription(req.Description); err != nil {
			return err
		}

		if err := tx.Save(&entry).Error; err != nil {
			return err
		}

		return preloadDiaryEntry(tx).Where("user_id = ?", claims.UserID).First(&entry, entry.ID).Error
	}); err != nil {
		writeDiaryError(c, err)
		return
	}

	c.JSON(http.StatusOK, requests.NewDiaryEntryResponse(entry))
}

func DeleteDiaryEntry(c *gin.Context) {
	entryID, ok := parseUintParam(c, "id")
	if !ok {
		return
	}

	claims, ok := currentAuthClaims(c)
	if !ok {
		return
	}

	db := diaryDB(c)
	if db == nil {
		return
	}

	var entry models.DiaryEntry
	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := preloadDiaryEntry(tx).Where("user_id = ?", claims.UserID).First(&entry, entryID).Error; err != nil {
			return err
		}

		return tx.Delete(&entry).Error
	}); err != nil {
		writeDiaryError(c, err)
		return
	}

	c.JSON(http.StatusOK, requests.NewDiaryEntryResponse(entry))
}
