package repository

import (
	"errors"
	"fmt"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/config"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/models"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

const (
	tableRelation = "relation_user_slugs"
	tableUser     = "users"
	tableSlug     = "slugs"
	tableHistory  = "user_segment_histories"
)

type TemplateRepositoryUserSlug interface {
	CreateSlug(slug models.Slug) error
	GetSlugs() ([]models.Slug, error)
	DeleteSlug(slug models.Slug) error

	CreateUser(user models.User) (int, error)
	GetUser(uuidUser int) (models.User, error)
	AddDelSlugToUser(uuidUser int, listForUser models.AddRemoveUserSlug) error
	DeleteUser(uuidUser int) error

	SaveSegmentHistory(userID int, addSegments, removeSegments []string)
	GetSegmentHistory(userId, yearStart, yearFinish int, monthStart, monthFinish time.Month) ([]models.UserSegmentHistory, error)
}

type TemplateRepositoryImpl struct {
	db *gorm.DB
}

func (t TemplateRepositoryImpl) CreateSlug(slug models.Slug) error {
	var existingSlug models.Slug
	if err := t.db.Table(tableSlug).Where("name_slug = ?", slug.NameSlug).First(&existingSlug).Error; err == nil {
		existingSlug.NameSlug = slug.NameSlug
		err = t.db.Save(&existingSlug).Error
		return err
	} else if err == gorm.ErrRecordNotFound {
		newRecord := models.Slug{
			NameSlug: slug.NameSlug,
		}
		if err := t.db.Create(&newRecord).Error; err != nil {
			return err
		}
		result := t.db.Table(tableSlug).Clauses(clause.OnConflict{DoNothing: true}).Create(&slug)
		return result.Error
	}
	return errors.New("nothing to do")
}

func (t TemplateRepositoryImpl) GetSlugs() ([]models.Slug, error) {
	var slugs []models.Slug
	result := t.db.Table(tableSlug).Find(&slugs)
	return slugs, result.Error
}

func (t TemplateRepositoryImpl) DeleteSlug(slug models.Slug) error {
	result := t.db.Unscoped().Table(tableSlug).Where("name_slug = ?", slug.NameSlug).Delete(slug.NameSlug)
	if result.Error != nil {
		return result.Error
	}
	result = t.db.Unscoped().Table(tableRelation).Where("name_slug = ?", slug.NameSlug).Delete(slug.NameSlug)
	return result.Error
}

func (t TemplateRepositoryImpl) CreateUser(user models.User) (int, error) {
	result := t.db.Table(tableUser).Clauses(clause.OnConflict{DoNothing: true}).Create(&user)
	return user.Id, result.Error
}

func (t TemplateRepositoryImpl) GetUser(uuidUser int) (models.User, error) {
	var relation []models.RelationUserSlug
	var user models.User
	result := t.db.Table(tableRelation).Where("user_id = ? ", uuidUser).Find(&relation)

	for _, segment := range relation {
		user.Slugs = append(user.Slugs, segment.NameSlug)
	}

	return user, result.Error
}

func (t TemplateRepositoryImpl) AddDelSlugToUser(uuidUser int, listForUser models.AddRemoveUserSlug) error {
	var user models.User
	var relation models.RelationUserSlug
	if err := t.db.Table(tableUser).Clauses(clause.OnConflict{DoNothing: true}).Find(&user, uuidUser).Error; err != nil {
		return err
	}

	for _, addSegment := range listForUser.AddSegments {
		var segment models.Slug
		if err := t.db.Table(tableSlug).Where("name_slug = ?", addSegment).First(&segment).Error; err != nil {
			return err
		}

		var existingRelation models.RelationUserSlug
		if err := t.db.Table(tableRelation).Where("name_slug = ? AND user_id = ?", addSegment, uuidUser).First(&existingRelation).Error; err == nil {
			continue
		} else if err == gorm.ErrRecordNotFound {

			relation.UserId = uuidUser
			relation.NameSlug = addSegment

			err := t.db.Table(tableRelation).Clauses(clause.OnConflict{DoNothing: true}).Create(&relation).Error
			if err != nil {
				return err
			}
		}
	}

	for _, removeSegment := range listForUser.RemoveSegments {
		err := t.db.Unscoped().Table(tableRelation).Clauses(clause.OnConflict{DoNothing: true}).Where("name_slug = ? AND user_id = ?", removeSegment, uuidUser).Delete(&models.RelationUserSlug{UserId: uuidUser, NameSlug: removeSegment}).Error
		if err != nil {
			return err
		}
	}

	if err := t.db.Save(relation).Error; err != nil {
		return err
	}

	return nil

}

func (t TemplateRepositoryImpl) DeleteUser(uuidUser int) error {
	result := t.db.Unscoped().Table(tableUser).Where("id = ?", uuidUser).Delete(uuidUser)
	if result.Error != nil {
		return result.Error
	}
	result = t.db.Unscoped().Table(tableRelation).Where("user_id = ?", uuidUser).Delete(uuidUser)
	return result.Error
}

func (t TemplateRepositoryImpl) SaveSegmentHistory(userID int, addSegments, removeSegments []string) {
	now := time.Now()
	for _, segment := range addSegments {
		history := models.UserSegmentHistory{
			UserId:      userID,
			SegmentSlug: segment,
			Operation:   "add",
			Timestamp:   now,
		}
		if err := t.db.Table(tableHistory).Create(&history).Error; err != nil {
			logger.Info(err.Error())
		}
	}
	for _, segment := range removeSegments {
		history := models.UserSegmentHistory{
			UserId:      userID,
			SegmentSlug: segment,
			Operation:   "remove",
			Timestamp:   now,
		}
		if err := t.db.Table(tableHistory).Create(&history).Error; err != nil {
			logger.Info(err.Error())
		}
	}
}

func (t TemplateRepositoryImpl) GetSegmentHistory(userId, yearStart, yearFinish int, monthStart, monthFinish time.Month) ([]models.UserSegmentHistory, error) {
	startOfMonth := time.Date(yearStart, monthStart, 1, 0, 0, 0, 0, time.UTC)
	endOfMonth := startOfMonth.AddDate(yearFinish, int(monthFinish), 0).Add(-time.Nanosecond)
	var history []models.UserSegmentHistory
	if err := t.db.Table(tableHistory).Where("user_id = ? AND timestamp BETWEEN ? AND ?", userId, startOfMonth, endOfMonth).Find(&history).Error; err != nil {
		return []models.UserSegmentHistory{}, err
	}
	return history, nil
}

func NewTemplateRepository() *TemplateRepositoryImpl {
	cfg := config.GetConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName, cfg.Database.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&models.User{}, &models.Slug{}, &models.RelationUserSlug{}, &models.UserSegmentHistory{})
	if err != nil {
		panic("failed to connect database")
	}
	pgSvc := &TemplateRepositoryImpl{db: db}
	return pgSvc
}
