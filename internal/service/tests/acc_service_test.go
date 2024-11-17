package tests

import (
	"context"
	"inventory-app/internal/mocks"
	"inventory-app/internal/model"
	"inventory-app/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterService_Success(t *testing.T) {
	mockRepo := new(mocks.AccRepoMethod)
	serv := service.NewAccService(mockRepo)

	acc := &model.Account{
		Username: "testuser",
		Email:    "testuser@example.com",
		Password: "password123",
	}
	_ = mockRepo.CreateAcc(context.Background(), acc)
	err2 := serv.RegisterService(acc)

	assert.NoError(t, err2)
	mockRepo.AssertExpectations(t)
}