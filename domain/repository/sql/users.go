package sql

import (
	"context"
	"github/yogabagas/print-in/domain/model"
	"github/yogabagas/print-in/service/users/repository"
	"strings"
)

const (
	insertUsers = `INSERT INTO users (uid, first_name, last_name, email, birthdate, username, password, created_by, updated_by) 
	VALUES (?,?,?,?,?,?,?,?,?)`
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
