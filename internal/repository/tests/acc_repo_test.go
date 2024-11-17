package tests

import (
	"context"
	"errors"
	"inventory-app/internal/mocks"
	"inventory-app/internal/model"
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateAcc_Success(t *testing.T) {
	mockRepo := new(mocks.AccRepoMethod)
	ctx := context.Background()

	account := &model.Account{
		UserID:    uuid.New(),
		Username:  "testuser",
		Email:     "test@example.com",
		Password:  "securepassword",
		CreatedAt: time.Now(),
	}

	mockRepo.On("CreateAcc", ctx, account).Return(nil)

	err := mockRepo.CreateAcc(ctx, account)
	log.Println(account)
	require.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
func TestCreateAcc_Error(t *testing.T) {
	mockRepo := new(mocks.AccRepoMethod)
	ctx := context.Background()

	account := &model.Account{
		UserID:    uuid.New(),
		Username:  "testuser",
		Email:     "test@example.com",
		Password:  "securepassword",
		CreatedAt: time.Now(),
	}

	mockRepo.On("CreateAcc", ctx, account).Return(errors.New("database error"))

	err := mockRepo.CreateAcc(ctx, account)

	require.Error(t, err)
	assert.Equal(t, "database error", err.Error())
	mockRepo.AssertExpectations(t)
}
func TestDeleteAcc_Success(t *testing.T) {
	mockRepo := new(mocks.AccRepoMethod)
	ctx := context.Background()
	userID := "12345"

	mockRepo.On("DeleteAcc", ctx, userID).Return(nil)

	err := mockRepo.DeleteAcc(ctx, userID)

	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "DeleteAcc", ctx, userID)
	mockRepo.AssertExpectations(t)
}
func TestDeleteAcc_Error(t *testing.T) {
	mockRepo := new(mocks.AccRepoMethod)
	ctx := context.Background()
	userID := "12345"

	mockError := errors.New("failed to delete account")
	mockRepo.On("DeleteAcc", ctx, userID).Return(mockError)

	err := mockRepo.DeleteAcc(ctx, userID)

	assert.Error(t, err)
	assert.Equal(t, mockError, err)
	mockRepo.AssertCalled(t, "DeleteAcc", ctx, userID)
	mockRepo.AssertExpectations(t)
}
func TestFindByUsername_Success(t *testing.T) {
	mockRepo := new(mocks.AccRepoMethod)
	ctx := context.Background()
	username := "testuser"

	expectedAccount := &model.Account{
		UserID:    uuid.New(),
		Username:  username,
		Email:     "test@example.com",
		Password:  "hashedpassword",
		CreatedAt: time.Now(),
	}

	mockRepo.On("FindByUsername", ctx, username).Return(expectedAccount, nil)

	account, err := mockRepo.FindByUsername(ctx, username)

	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, expectedAccount, account)
	log.Println(expectedAccount, account)
	mockRepo.AssertCalled(t, "FindByUsername", ctx, username)
	mockRepo.AssertExpectations(t)
}

func TestFindByUsername_NoDataError(t *testing.T) {
	mockRepo := new(mocks.AccRepoMethod)
	ctx := context.Background()
	username := "nonexistentuser"

	mockError := errors.New("no data found")
	mockRepo.On("FindByUsername", ctx, username).Return(nil, mockError)

	account, err := mockRepo.FindByUsername(ctx, username)

	assert.Error(t, err)
	assert.Nil(t, account)
	assert.Equal(t, mockError, err)
	mockRepo.AssertCalled(t, "FindByUsername", ctx, username)
	mockRepo.AssertExpectations(t)
}

func TestFindByUsername_DBError(t *testing.T) {
	mockRepo := new(mocks.AccRepoMethod)
	ctx := context.Background()
	username := "testuser"

	mockError := errors.New("failed to get account database error")
	mockRepo.On("FindByUsername", ctx, username).Return(nil, mockError)

	account, err := mockRepo.FindByUsername(ctx, username)

	assert.Error(t, err)
	assert.Nil(t, account)
	assert.Equal(t, mockError, err)
	mockRepo.AssertCalled(t, "FindByUsername", ctx, username)
	mockRepo.AssertExpectations(t)
}

func TestUpdateAcc_Success(t *testing.T) {
	mockRepo := new(mocks.AccRepoMethod)
	ctx := context.Background()

	originalAcc := &model.Account{
		UserID:   uuid.New(),
		Username: "oldusername",
		Email:    "oldemail@example.com",
		Password: "oldpassword",
	}

	updatedAcc := &model.Account{
		UserID:   originalAcc.UserID,
		Username: "newusername",
		Email:    "newemail@example.com",
		Password: "newpassword",
	}
	log.Println(originalAcc)
	mockRepo.On("UpdateAcc", ctx, updatedAcc).Return(updatedAcc, nil)

	result, err := mockRepo.UpdateAcc(ctx, updatedAcc)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, updatedAcc.Username, result.Username)
	assert.Equal(t, updatedAcc.Email, result.Email)
	assert.Equal(t, updatedAcc.Password, result.Password)
	mockRepo.AssertCalled(t, "UpdateAcc", ctx, updatedAcc)
	mockRepo.AssertExpectations(t)
}


func TestUpdateAcc_NoRowsAffected_Error(t *testing.T) {
	mockRepo := new(mocks.AccRepoMethod)
	ctx := context.Background()

	acc := &model.Account{
		UserID:   uuid.New(),
		Username: "nonexistentuser",
		Email:    "nonexistent@example.com",
		Password: "password",
	}

	mockError := errors.New("no account found to update")
	mockRepo.On("UpdateAcc", ctx, acc).Return(nil, mockError)

	updatedAcc, err := mockRepo.UpdateAcc(ctx, acc)

	assert.Error(t, err)
	assert.Nil(t, updatedAcc)
	assert.Equal(t, mockError, err)
	mockRepo.AssertCalled(t, "UpdateAcc", ctx, acc)
	mockRepo.AssertExpectations(t)
}

func TestUpdateAcc_DBError(t *testing.T) {
	mockRepo := new(mocks.AccRepoMethod)
	ctx := context.Background()

	acc := &model.Account{
		UserID:   uuid.New(),
		Username: "userwitherror",
		Email:    "error@example.com",
		Password: "password",
	}

	mockError := errors.New("failed to update account: database error")
	mockRepo.On("UpdateAcc", ctx, acc).Return(nil, mockError)

	updatedAcc, err := mockRepo.UpdateAcc(ctx, acc)

	assert.Error(t, err)
	assert.Nil(t, updatedAcc)
	assert.Equal(t, mockError, err)
	mockRepo.AssertCalled(t, "UpdateAcc", ctx, acc)
	mockRepo.AssertExpectations(t)
}