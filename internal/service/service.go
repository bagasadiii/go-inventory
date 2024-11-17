package service

import "inventory-app/internal/model"

type AccServiceMethod interface {
	RegisterService(acc *model.Account)error
	LoginService(username, password string)(string, error)
	UpdateAccountService(acc *model.Account)(*model.Account,error)
	GetAccountByUsername(username string)(*model.Account, error)
	DeleteAccountService(username, password string)error
}