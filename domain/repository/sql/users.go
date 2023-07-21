package sql

import (
	"context"
	"database/sql"
	"github/yogabagas/print-in/domain/model"
	"github/yogabagas/print-in/service/users/repository"
	"strings"
)

const (
	insertUsers = `INSERT INTO users (uid, first_name, last_name, email, birthdate, username, password, created_by, updated_by) 
	VALUES (?,?,?,?,?,?,?,?,?)`
	selectByEmail = `SELECT uid, first_name, last_name, email, birthdate, username, password from users WHERE email = ? `
)

type UsersRepositoryImpl struct {
	db DBExecutor
}

func NewUsersRepository(db DBExecutor) repository.UsersRepository {
	return &UsersRepositoryImpl{db: db}
}

func (ur *UsersRepositoryImpl) CreateUsers(ctx context.Context, req *model.User) error {

	_, err := ur.db.ExecContext(ctx, insertUsers, req.UID, req.FirstName, req.LastName, req.Email, req.Birthdate, req.Username,
		req.Password, req.CreatedBy, req.UpdatedBy)
	if err != nil && !strings.Contains(err.Error(), "duplicate") {
		return err
	}

	return nil
}

func (ur *UsersRepositoryImpl) FindByEmail(ctx context.Context, email string) (resp *model.User, err error) {
	resp = &model.User{}
	err = ur.db.QueryRowContext(ctx, selectByEmail, email).
		Scan(&resp.UID, &resp.FirstName, &resp.LastName, &resp.Email, &resp.Birthdate, &resp.Username, &resp.Password)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return resp, err
}
