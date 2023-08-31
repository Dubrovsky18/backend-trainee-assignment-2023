package repository

import (
	"testing"

	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Table(name string) *gorm.DB {
	args := m.Called(name)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Find(dest interface{}) *gorm.DB {
	args := m.Called(dest)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Create(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Delete(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Unscoped() *gorm.DB {
	args := m.Called()
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Save(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func TestCreateSlug(t *testing.T) {
	mockDB := new(MockDB)
	repo := &TemplateRepositoryImpl{db: mockDB}

	slug := models.Slug{NameSlug: "segment1"}
	mockDB.On("Table", tableSlug).Return(mockDB)
	mockDB.On("Create", &slug).Return(&gorm.DB{})

	err := repo.CreateSlug(slug)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestGetSlugs(t *testing.T) {
	mockDB := new(MockDB)
	repo := &TemplateRepositoryImpl{db: mockDB}

	expectedSlugs := []models.Slug{
		{NameSlug: "segment1"},
		{NameSlug: "segment2"},
	}
	mockDB.On("Find", mock.Anything).Return(mockDB)
	mockDB.On("Find", &[]models.Slug{}).Run(func(args mock.Arguments) {
		slugsPtr := args.Get(0).(*[]models.Slug)
		*slugsPtr = expectedSlugs
	})

	slugs, err := repo.GetSlugs()

	assert.NoError(t, err)
	assert.Equal(t, expectedSlugs, slugs)
	mockDB.AssertExpectations(t)
}
