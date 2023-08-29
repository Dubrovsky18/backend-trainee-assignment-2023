package services

import (
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/models"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/repository"
)

type TemplateServiceUserSlug interface {
	CreateSlug(slug models.Slug) error
	GetSlugs() ([]models.Slug, error)
	DeleteSlug(slug models.Slug) error

	AddDelSlugToUser(uuidUser int, listForUser models.AddRemoveUserSlug) error

	GetUser(uuidUser int) (models.User, error)

	CreateUser(user models.User) (int, error)
	DeleteUser(uuidUser int) error
}

type TemplateServiceRepoImpl struct {
	repos repository.TemplateRepositoryUserSlug
}

func (t TemplateServiceRepoImpl) AddDelSlugToUser(uuidUser int, listForUser models.AddRemoveUserSlug) error {
	return t.repos.AddDelSlugToUser(uuidUser, listForUser)
}

func (t TemplateServiceRepoImpl) CreateSlug(slug models.Slug) error {
	return t.repos.CreateSlug(slug)
}

func (t TemplateServiceRepoImpl) GetSlugs() ([]models.Slug, error) {
	return t.repos.GetSlugs()
}

func (t TemplateServiceRepoImpl) DeleteSlug(slug models.Slug) error {
	return t.repos.DeleteSlug(slug)
}

func (t TemplateServiceRepoImpl) GetUser(uuidUser int) (models.User, error) {
	return t.repos.GetUser(uuidUser)
}

func (t TemplateServiceRepoImpl) CreateUser(user models.User) (int, error) {
	return t.repos.CreateUser(user)
}

func (t TemplateServiceRepoImpl) DeleteUser(uuidUser int) error {
	return t.repos.DeleteUser(uuidUser)
}

func NewTemplateRepository(repos *repository.TemplateRepositoryImpl) *TemplateServiceRepoImpl {
	return &TemplateServiceRepoImpl{repos: repos}
}
