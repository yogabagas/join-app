package sql

import (
	"context"
	"database/sql"
	"github/yogabagas/print-in/domain/model"
	"github/yogabagas/print-in/pkg/cache"
	"github/yogabagas/print-in/service/users/repository"
	"strings"
)

const (
	insertUsers = `INSERT INTO users (uid, first_name, last_name, email, birthdate, username, password, created_by, updated_by) 
	VALUES (?,?,?,?,?,?,?,?,?)`
	selectByEmail = `select a.user_uid, a.role_uid, u.password  from users u join authz a on u.uid  = a.user_uid where email = ? `
)

type UsersRepositoryImpl struct {
	db    DBExecutor
	cache cache.Cache
}

func NewUsersRepository(db DBExecutor, cache cache.Cache) repository.UsersRepository {
	return &UsersRepositoryImpl{db: db, cache: cache}
}

func (ur *UsersRepositoryImpl) CreateUsers(ctx context.Context, req *model.User) error {

	_, err := ur.db.ExecContext(ctx, insertUsers, req.UID, req.FirstName, req.LastName, req.Email, req.Birthdate, req.Username,
		req.Password, req.CreatedBy, req.UpdatedBy)
	if err != nil && !strings.Contains(err.Error(), "duplicate") {
		return err
	}

	return nil
}

func (ur *UsersRepositoryImpl) ReadUserByEmail(ctx context.Context, email string) (resp *model.Session, err error) {
	resp = &model.Session{}
	err = ur.db.QueryRowContext(ctx, selectByEmail, email).
		Scan(&resp.UserUID, &resp.RoleUID, &resp.Password)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return resp, err
}

func (ur *UsersRepositoryImpl) CreateSession(ctx context.Context, userUUID string) error {
	return ur.cache.Set(ctx, "user_uuid:"+userUUID, true, 1440)
}

func (ur *UsersRepositoryImpl) DeleteSession(ctx context.Context, userUUID string) error {
	return ur.cache.Delete(ctx, "user_uuid:"+userUUID)
}
