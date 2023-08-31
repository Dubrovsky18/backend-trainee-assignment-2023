package services

import (
	"encoding/csv"
	"fmt"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/models"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/repository"
	"os"
	"strconv"
	"time"
)

type TemplateServiceUserSlug interface {
	CreateSlug(slug models.Slug) error
	GetSlugs() ([]models.Slug, error)
	DeleteSlug(slug models.Slug) error

	AddDelSlugToUser(uuidUser int, listForUser models.AddRemoveUserSlug) error

	GetUser(uuidUser int) (models.User, error)

	CreateUser(user models.User) (int, error)
	DeleteUser(uuidUser int) error

	GetSegmentHistory(userID, yearStart, yearFinish int, monthStart, monthFinish time.Month) (string, error)
}

type TemplateServiceRepoImpl struct {
	Repos repository.TemplateRepositoryUserSlug
}

func (t TemplateServiceRepoImpl) AddDelSlugToUser(uuidUser int, listForUser models.AddRemoveUserSlug) error {
	t.Repos.SaveSegmentHistory(uuidUser, listForUser.AddSegments, listForUser.RemoveSegments)
	return t.Repos.AddDelSlugToUser(uuidUser, listForUser)
}

func (t TemplateServiceRepoImpl) CreateSlug(slug models.Slug) error {
	return t.Repos.CreateSlug(slug)
}

func (t TemplateServiceRepoImpl) GetSlugs() ([]models.Slug, error) {
	return t.Repos.GetSlugs()
}

func (t TemplateServiceRepoImpl) DeleteSlug(slug models.Slug) error {
	return t.Repos.DeleteSlug(slug)
}

func (t TemplateServiceRepoImpl) GetUser(uuidUser int) (models.User, error) {
	return t.Repos.GetUser(uuidUser)
}

func (t TemplateServiceRepoImpl) CreateUser(user models.User) (int, error) {
	return t.Repos.CreateUser(user)
}

func (t TemplateServiceRepoImpl) DeleteUser(uuidUser int) error {
	return t.Repos.DeleteUser(uuidUser)
}

func (t TemplateServiceRepoImpl) GetSegmentHistory(userID, yearStart, yearFinish int, monthStart, monthFinish time.Month) (string, error) {

	history, err := t.Repos.GetSegmentHistory(userID, yearStart, yearFinish, monthStart, monthFinish)

	file, err := os.Create(fmt.Sprintf("files/report_%d.csv", userID))
	if err != nil {
		return "", err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"идентификатор пользователя", "сегмент", "операция", "дата и время"}
	if err := writer.Write(header); err != nil {
		return "", err
	}

	for _, entry := range history {
		row := []string{
			strconv.Itoa(entry.UserId),
			entry.SegmentSlug,
			entry.Operation,
			entry.Timestamp.Format("2006-01-02 15:04:05"),
		}
		if err := writer.Write(row); err != nil {
			return "", err
		}
	}

	return file.Name(), nil
}

func NewTemplateRepository(repos *repository.TemplateRepositoryImpl) *TemplateServiceRepoImpl {
	return &TemplateServiceRepoImpl{Repos: repos}
}
