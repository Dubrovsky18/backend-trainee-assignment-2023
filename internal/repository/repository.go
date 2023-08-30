package repository

import (
	"fmt"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/config"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	tableRelation = "relation_user_slugs"
	tableUser     = "users"
	tableSlug     = "slugs"
)

type TemplateRepositoryUserSlug interface {
	CreateSlug(slug models.Slug) error
	GetSlugs() ([]models.Slug, error)
	DeleteSlug(slug models.Slug) error

	CreateUser(user models.User) (int, error)
	GetUser(uuidUser int) (models.User, error)
	AddDelSlugToUser(uuidUser int, listForUser models.AddRemoveUserSlug) error
	DeleteUser(uuidUser int) error
}

type TemplateRepositoryImpl struct {
	db *gorm.DB
}

func (t TemplateRepositoryImpl) CreateSlug(slug models.Slug) error {
	result := t.db.Table(tableSlug).Clauses(clause.OnConflict{DoNothing: true}).Create(&slug)
	return result.Error
}

func (t TemplateRepositoryImpl) GetSlugs() ([]models.Slug, error) {
	var slugs []models.Slug
	result := t.db.Find(&slugs)
	return slugs, result.Error
}

func (t TemplateRepositoryImpl) DeleteSlug(slug models.Slug) error {
	result := t.db.Delete(&slug)
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

		relation.UserId = uuidUser
		relation.NameSlug = addSegment

		if err := t.db.Table(tableRelation).Clauses(clause.OnConflict{DoNothing: true}).Where("name_slug = ? AND user_id = ?", addSegment, uuidUser).Error; err != nil {
			return fmt.Errorf("already in system")
		}

		err := t.db.Table(tableRelation).Clauses(clause.OnConflict{DoNothing: true}).Create(&relation).Error
		if err != nil {
			return err
		}
	}

	for _, removeSegment := range listForUser.RemoveSegments {
		err := t.db.Table(tableRelation).Clauses(clause.OnConflict{DoNothing: true}).Where("name_slug = ? AND user_id = ?", removeSegment, uuidUser).Delete(&models.RelationUserSlug{UserId: uuidUser, NameSlug: removeSegment}).Error
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
	result := t.db.Unscoped().Table(tableUser).Delete(uuidUser)
	t.db.Unscoped().Table(tableRelation).Where("user_id = ?", uuidUser).Delete(uuidUser)
	return result.Error
}

func NewTemplateRepository() *TemplateRepositoryImpl {
	cfg := config.GetConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName, cfg.Database.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&models.User{}, &models.Slug{}, &models.RelationUserSlug{})
	if err != nil {
		panic("failed to connect database")
	}
	pgSvc := &TemplateRepositoryImpl{db: db}
	return pgSvc
}
