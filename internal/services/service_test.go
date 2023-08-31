package services_test

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/services"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock репозитория для использования в тестах
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) SaveSegmentHistory(userID int, addSegments, removeSegments []string) {
	//TODO implement me
	panic("implement me")
}

func (m *MockRepository) TestGetSegmentHistory(t *testing.T) {
	mockRepo := new(MockRepository)
	service := services.NewTemplateRepository(mockRepo)

	expectedHistory := []models.UserSegmentHistory{
		{UserId: 1, SegmentSlug: "segment1", Operation: "add", Timestamp: time.Now()},
	}
	mockRepo.On("GetSegmentHistory", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(expectedHistory, nil)

	history, err := service.GetSegmentHistory(1, 2023, 2023, time.January, time.February)

	assert.NoError(t, err)
	assert.Equal(t, expectedHistory, history)
	mockRepo.AssertExpectations(t)
}

func (m *MockRepository) AddDelSlugToUser(uuidUser int, listForUser models.AddRemoveUserSlug) error {
	args := m.Called(uuidUser, listForUser)
	return args.Error(0)
}

func (m *MockRepository) CreateSlug(slug models.Slug) error {
	args := m.Called(slug)
	return args.Error(0)
}

func (m *MockRepository) GetSlugs() ([]models.Slug, error) {
	args := m.Called()
	return args.Get(0).([]models.Slug), args.Error(1)
}

func (m *MockRepository) DeleteSlug(slug models.Slug) error {
	args := m.Called(slug)
	return args.Error(0)
}

func (m *MockRepository) GetUser(uuidUser int) (models.User, error) {
	args := m.Called(uuidUser)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockRepository) CreateUser(user models.User) (int, error) {
	args := m.Called(user)
	return args.Int(0), args.Error(1)
}

func (m *MockRepository) DeleteUser(uuidUser int) error {
	args := m.Called(uuidUser)
	return args.Error(0)
}

func TestAddDelSlugToUser(t *testing.T) {
	repo := new(MockRepository)
	service := services.NewTemplateRepository(repo)

	listForUser := models.AddRemoveUserSlug{
		AddSegments: []string{"segment1"},
	}

	repo.On("AddDelSlugToUser", 1, listForUser).Return(nil)

	err := service.AddDelSlugToUser(1, listForUser)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestAddDelSlugToUser_RepositoryError(t *testing.T) {
	repo := new(MockRepository)
	service := services.NewTemplateRepository(repo)

	listForUser := models.AddRemoveUserSlug{
		AddSegments: []string{"segment1"},
	}

	expectedErr := errors.New("repository error")
	repo.On("AddDelSlugToUser", 1, listForUser).Return(expectedErr)

	err := service.AddDelSlugToUser(1, listForUser)

	assert.EqualError(t, err, expectedErr.Error())
	repo.AssertExpectations(t)
}

func TestCreateSlug(t *testing.T) {
	repo := new(MockRepository)
	service := services.NewTemplateRepository(repo)

	slug := models.Slug{NameSlug: "segment1"}
	repo.On("CreateSlug", slug).Return(nil)

	err := service.CreateSlug(slug)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestGetSlugs(t *testing.T) {
	repo := new(MockRepository)
	service := services.NewTemplateRepository(repo)

	expectedSlugs := []models.Slug{
		{NameSlug: "segment1"},
		{NameSlug: "segment2"},
	}

	repo.On("GetSlugs").Return(expectedSlugs, nil)

	slugs, err := service.GetSlugs()

	assert.NoError(t, err)
	assert.Equal(t, expectedSlugs, slugs)
	repo.AssertExpectations(t)
}

func TestDeleteSlug(t *testing.T) {
	repo := new(MockRepository)
	service := services.NewTemplateRepository(repo)

	slug := models.Slug{NameSlug: "segment1"}
	repo.On("DeleteSlug", slug).Return(nil)

	err := service.DeleteSlug(slug)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestGetUser(t *testing.T) {
	repo := new(MockRepository)
	service := services.NewTemplateRepository(repo)

	expectedUser := models.User{Id: 1}
	repo.On("GetUser", 1).Return(expectedUser, nil)

	user, err := service.GetUser(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	repo.AssertExpectations(t)
}

func TestCreateUser(t *testing.T) {
	repo := new(MockRepository)
	service := services.NewTemplateRepository(repo)

	user := models.User{Id: 1}
	repo.On("CreateUser", user).Return(1, nil)

	createdUserID, err := service.CreateUser(user)

	assert.NoError(t, err)
	assert.Equal(t, 1, createdUserID)
	repo.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	repo := new(MockRepository)
	service := services.NewTemplateRepository(repo)

	repo.On("DeleteUser", 1).Return(nil)

	err := service.DeleteUser(1)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestGetSegmentHistory(t *testing.T) {
	mockRepo := new(MockRepository)
	service := services.TemplateServiceRepoImpl{Repos: mockRepo}

	userID := 1
	yearStart := 2023
	yearFinish := 2023
	monthStart := time.January
	monthFinish := time.February

	history := []models.UserSegmentHistory{
		{UserId: userID, SegmentSlug: "segment1", Operation: "add", Timestamp: time.Now()},
		{UserId: userID, SegmentSlug: "segment2", Operation: "remove", Timestamp: time.Now()},
	}

	mockRepo.On("GetSegmentHistory", userID, yearStart, yearFinish, monthStart, monthFinish).
		Return(history, nil)

	expectedFilePath := fmt.Sprintf("files/report_%d.csv", userID)

	file, err := os.Create(expectedFilePath)
	assert.NoError(t, err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"идентификатор пользователя", "сегмент", "операция", "дата и время"}
	assert.NoError(t, writer.Write(header))

	for _, entry := range history {
		row := []string{
			strconv.Itoa(entry.UserId),
			entry.SegmentSlug,
			entry.Operation,
			entry.Timestamp.Format("2006-01-02 15:04:05"),
		}
		assert.NoError(t, writer.Write(row))
	}

	expectedResult := expectedFilePath

	result, err := service.GetSegmentHistory(userID, yearStart, yearFinish, monthStart, monthFinish)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)

	mockRepo.AssertExpectations(t)
}
