package repository

import (
	"context"
	"inventory-app/internal/model"
)

type AccRepoMethod interface {
	CreateAcc(ctx context.Context,acc *model.Account) error
	DeleteAcc(ctx context.Context,id string) error
	UpdateAcc(ctx context.Context,acc *model.Account) (*model.Account, error)
	FindByUsername(ctx context.Context,username string) (*model.Account, error)
}