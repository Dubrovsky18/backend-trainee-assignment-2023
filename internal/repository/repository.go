package repository

import (
	"fmt"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/config"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	result := t.db.Create(&slug)
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
	result := t.db.Create(&user)
	return user.Id, result.Error
}

func (t TemplateRepositoryImpl) GetUser(uuidUser int) (models.User, error) {
	var user models.User
	result := t.db.Where("user_id = ? ", uuidUser).Find(&user.Slugs)
	return user, result.Error
}

func (t TemplateRepositoryImpl) AddDelSlugToUser(uuidUser int, listForUser models.AddRemoveUserSlug) error {
	var user models.User
	var slug models.Slug
	if err := t.db.First(&user, uuidUser).Error; err != nil {
		return err
	}

	for _, addSegment := range listForUser.AddSegments {
		var segment models.Slug
		if err := t.db.Where("name_slug = ?", addSegment).First(&segment).Error; err != nil {
			continue
		}
		user.Slugs = append(user.Slugs, segment.NameSlug)
	}

	for _, removeSegment := range listForUser.RemoveSegments {
		for i, segment := range user.Slugs {
			if segment == removeSegment {
				user.Slugs = append(user.Slugs[:i], user.Slugs[i+1:]...)
				break
			}
		}
	}

	if err := t.db.Save(&user).Error; err != nil {
		return err
	}
	if err := t.db.Save(&slug).Error; err != nil {
		return err
	}

	return nil

}

func (t TemplateRepositoryImpl) DeleteUser(uuidUser int) error {
	result := t.db.Delete(uuidUser)
	return result.Error
}

func NewTemplateRepository() *TemplateRepositoryImpl {
	cfg := config.GetConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName, cfg.Database.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&models.User{}, &models.Slug{})
	if err != nil {
		panic("failed to connect database")
	}
	pgSvc := &TemplateRepositoryImpl{db: db}
	return pgSvc
}
