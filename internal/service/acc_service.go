package service

import (
	"context"
	"inventory-app/internal/helper"
	"inventory-app/internal/model"
	"inventory-app/internal/repository"
	"time"

	"github.com/google/uuid"
)

type AccService struct {
	repo repository.AccRepoMethod
}
func NewAccService(repo repository.AccRepoMethod)AccServiceMethod{
	return &AccService{
		repo: repo,
	}
}
func(a *AccService)RegisterService(acc *model.Account)error{
	hashedPassword, err := helper.GenerateHash(acc.Password)
	if err != nil {
		return err
	}
	newAcc := &model.Account{
		UserID: uuid.New(),
		Username: acc.Username,
		Email: acc.Email,
		Password: hashedPassword,
		CreatedAt: time.Now(),
	}
	if err := a.repo.CreateAcc(context.Background(), newAcc); err != nil {
		return err
	}
	return nil
}
func(a *AccService)LoginService(username, password string)(string, error){
	acc, err := a.repo.FindByUsername(context.Background(), username)
	if err != nil {
		return "", err
	}
	if err := helper.CheckPassword(acc.Password, password); err != nil {
		return "", err
	}
	token, err := helper.GenerateToken(username)
	if err != nil {
		return "", err
	}
	return token, nil
}
func(a *AccService)UpdateAccountService(acc *model.Account)(*model.Account,error){
	newAcc, err := a.repo.UpdateAcc(context.Background(), acc)
	if err != nil {
		return nil,err
	}
	return newAcc, err
}
func(a *AccService)DeleteAccountService(username, password string)error{
	acc, err := a.repo.FindByUsername(context.Background(), username)
	if err != nil {
		return err
	}
	if err := helper.CheckPassword(acc.Password, password); err != nil {
		return err
	}
	if err := a.repo.DeleteAcc(context.Background(), acc.UserID.String()); err != nil {
		return err
	}
	return nil
}
func(a *AccService)GetAccountByUsername(username string)(*model.Account, error){
	acc, err := a.repo.FindByUsername(context.Background(), username)
	if err != nil {
		return nil,err
	}
	return acc, nil
}