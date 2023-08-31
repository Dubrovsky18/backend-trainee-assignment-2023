package userSlug_test

import (
	"bytes"
	"encoding/json"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/web/controllers/apiv1/userSlug"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/models"
	"github.com/stretchr/testify/assert"
)

type MockService struct {
}

func (m *MockService) CreateSlug(slug models.Slug) error {
	return nil
}

func (m *MockService) GetSlugs() ([]models.Slug, error) {
	return []models.Slug{{NameSlug: "slug1"}, {NameSlug: "slug2"}}, nil
}

func (m *MockService) GetSegmentHistory(userID, yearStart, yearFinish int, monthStart, monthFinish time.Month) (string, error) {
	return "test-file.csv", nil
}

func TestCreateSlug(t *testing.T) {
	mockService := new(MockService)
	controller := userSlug.NewController(mockService)

	reqBody := []byte(`{"name_slug": "test-slug"}`)
	req, _ := http.NewRequest("POST", "/slug/create", bytes.NewBuffer(reqBody))
	recorder := httptest.NewRecorder()

	controller.CreateSlug(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetSlugs(t *testing.T) {
	mockService := new(MockService)
	controller := userSlug.NewController(mockService)

	req, _ := http.NewRequest("GET", "/slug/get_all", nil)
	recorder := httptest.NewRecorder()

	controller.GetSlugs(recorder, req)

	expectedSlugs := []models.Slug{{NameSlug: "slug1"}, {NameSlug: "slug2"}}
	var slugs []models.Slug
	err := json.Unmarshal(recorder.Body.Bytes(), &slugs)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, expectedSlugs, slugs)
}

func TestGetSegmentsHistory(t *testing.T) {
	mockService := new(MockService)
	controller := userSlug.NewController(mockService)

	reqBody := []byte(`{
		"user_id": 1,
		"year_start": 2023,
		"year_finish": 2023,
		"month_start": 1,
		"month_finish": 2
	}`)
	req, _ := http.NewRequest("POST", "/api/v1/users/extra/history/1", bytes.NewBuffer(reqBody))
	recorder := httptest.NewRecorder()

	controller.GetSegmentsHistory(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "test-file.csv", recorder.Header().Get("Content-Disposition"))
}
