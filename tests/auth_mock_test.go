package test

import (
	model "staycation/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	mock.Mock
}

func (m *UserRepoMock) Create(user *model.UserTest) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepoMock) FindByEmail(email string) (*model.UserTest, error) {
	args := m.Called(email)
	return args.Get(0).(*model.UserTest), args.Error(1)
}

func (m *UserRepoMock) FindByPhone(phone string) (*model.UserTest, error) {
	args := m.Called(phone)
	return args.Get(0).(*model.UserTest), args.Error(1)
}

func TestCreateUser(t *testing.T) {
	userRepoMock := new(UserRepoMock)
	user := &model.UserTest{
		ID:       1,
		Email:    "test@example.com",
		Password: "hashedpassword",
	}

	userRepoMock.On("Create", user).Return(nil)

	err := userRepoMock.Create(user)
	assert.Nil(t, err)
}

func TestFindByEmail(t *testing.T) {
	userRepoMock := new(UserRepoMock)
	user := &model.UserTest{
		ID:       1,
		Email:    "test@example.com",
		Password: "hashedpassword",
	}

	userRepoMock.On("FindByEmail", "test@example.com").Return(user, nil)

	result, err := userRepoMock.FindByEmail("test@example.com")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "test@example.com", result.Email)
}

func TestFindByPhone(t *testing.T) {
	userRepoMock := new(UserRepoMock)
	user := &model.UserTest{
		ID:       1,
		Phone:    "1234567890",
		Password: "hashedpassword",
	}

	userRepoMock.On("FindByPhone", "1234567890").Return(user, nil)

	result, err := userRepoMock.FindByPhone("1234567890")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "1234567890", result.Phone)
}
