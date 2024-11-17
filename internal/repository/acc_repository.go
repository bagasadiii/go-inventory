package repository

import (
	"context"
	"errors"
	"inventory-app/internal/model"

	"github.com/jackc/pgx/v5"
)

type AccRepository struct {
	conn *pgx.Conn
}
func NewAccRepository(conn *pgx.Conn)AccRepoMethod{
	return &AccRepository{
		conn: conn,
	}
}
func(a *AccRepository)CreateAcc(ctx context.Context,acc *model.Account)error{
	query := `insert into account (user_id, username, email, password, created_at) values ($1,$2,$3,$4,$5);`

	_, err := a.conn.Exec(ctx,query,acc.UserID,acc.Username,acc.Email,acc.Password,acc.CreatedAt)
	if err != nil {
		return errors.New("failed to create account" + err.Error())
	}
	return nil
}
func(a *AccRepository)DeleteAcc(ctx context.Context,id string)error{
	query := `delete from account where user_id=$1;`

	_, err := a.conn.Exec(ctx,query, id)
	if err != nil {
		return errors.New("failed to delete account" + err.Error())
	}
	return nil
}
func(a *AccRepository)UpdateAcc(ctx context.Context, acc *model.Account)(*model.Account, error){
	query := `update account set username=$1, email=$2, password=$3 where user_id=$4;`

	result, err := a.conn.Exec(ctx, query, acc.Username, acc.Email, acc.Password, acc.UserID)
	if err != nil {
		return nil, errors.New("failed to update account: " + err.Error())
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, errors.New("no account found to update")
	}
	return acc, nil
}
func(a *AccRepository)FindByUsername(ctx context.Context,username string)(*model.Account, error){
	query := `select user_id, username, email, password from account where username=$1`

	var acc model.Account
	err := a.conn.QueryRow(ctx, query, username).Scan(&acc.UserID,&acc.Username,&acc.Email,&acc.Password)
	if err != nil {
		if err == pgx.ErrNoRows{
			return nil, errors.New("no data found")
		}
		return nil, errors.New("failed to get account " + err.Error())
	}
	return &acc, nil
}