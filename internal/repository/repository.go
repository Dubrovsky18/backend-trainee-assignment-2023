package repository

import (
	"fmt"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/config"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
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
	AddDelSlugToUser(uuidUser int, listForUser models.AddRemoveUserSlug) []error
	DeleteUser(uuidUser int) error
}

type TemplateRepositoryImpl struct {
	db *gorm.DB
}

func (t TemplateRepositoryImpl) CreateSlug(slug models.Slug) error {
	var existingSlug models.Slug
	result := t.db.Table(tableSlug).Where("name = ?", slug.NameSlug).First(&existingSlug)

	if existingSlug.ID == 0 {
		result = t.db.Table(tableSlug).Create(&slug)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func (t TemplateRepositoryImpl) GetSlugs() ([]models.Slug, error) {
	var slugs []models.Slug
	result := t.db.Table(tableSlug).Find(&slugs)
	return slugs, result.Error
}

func (t TemplateRepositoryImpl) DeleteSlug(slug models.Slug) error {
	result := t.db.Table(tableSlug).Delete(&slug)
	return result.Error
}

func (t TemplateRepositoryImpl) CreateUser(user models.User) (int, error) {
	result := t.db.Table(tableUser).Create(&user)
	return user.Id, result.Error

}

func (t TemplateRepositoryImpl) GetUser(uuidUser int) (models.User, error) {
	var user models.User

	result := t.db.Table(tableRelation).Where("user_id = ?", uuidUser).Find(&user.Slugs)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (t TemplateRepositoryImpl) AddDelSlugToUser(uuidUser int, listForUser models.AddRemoveUserSlug) []error {
	var user models.User
	user.Id = uuidUser
	var errArray []error

	if err := t.db.Table(tableUser).Where("id = ?", user.Id).First(&user).Error; err != nil {
		return []error{err}
	} else {
		for _, addSegment := range listForUser.AddSegments {
			var segment models.Slug
			if err := t.db.Where("name_slug = ?", addSegment).First(&segment).Error; err != nil {
				errArray = append(errArray, err)
				continue
			} else {
				var relation models.RelationUserSlug
				if err := t.db.Table(tableRelation).Where("name_slug = ? AND user_id = ?", addSegment, user.Id).First(&relation).Error; err != nil {
					relation = models.RelationUserSlug{
						UserId:   user.Id,
						NameSlug: addSegment,
					}
					if err := t.db.Create(&relation).Error; err != nil {
						errArray = append(errArray, err)
					}
				} else {
					errArray = append(errArray, fmt.Errorf("Slug = %s with user = %s already in base\n", addSegment, user.Id))
				}
			}
		}
	}

	for _, removeSegment := range listForUser.RemoveSegments {
		var relation models.RelationUserSlug
		if err := t.db.Unscoped().Where("user_id = ? AND name_slug = ?", user.Id, removeSegment).Delete(&relation).Error; err != nil {
			errArray = append(errArray, err)
		}
	}
	return errArray
}

func (t TemplateRepositoryImpl) DeleteUser(uuidUser int) error {
	result := t.db.Table(tableUser).Delete(uuidUser)
	return result.Error

}

func NewTemplateRepository() *TemplateRepositoryImpl {
	cfg := config.GetConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName, cfg.Database.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&models.User{}, &models.Slug{})
	if err != nil {
		log.Fatalf("Migrate User and Slug not to do: %s", err.Error())
	}
	err = db.AutoMigrate(&models.RelationUserSlug{})
	if err != nil {
		log.Fatalf("Migrate Relation not to do: %s", err.Error())
	}
	pgSvc := &TemplateRepositoryImpl{db: db}
	return pgSvc
}
