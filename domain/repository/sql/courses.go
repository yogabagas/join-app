package sql

import (
	"context"
	"github/yogabagas/join-app/domain/model"
	"github/yogabagas/join-app/service/courses/repository"
	"strings"
)

const (
	insertCourses = `INSERT INTO courses (uid, subject) VALUES (?,?)`
)

type CoursesRepositoryImpl struct {
	db DBExecutor
}

func NewCoursesRepository(db DBExecutor) repository.CoursesRepository {
	return &CoursesRepositoryImpl{db: db}
}

func (rr *CoursesRepositoryImpl) CreateCourses(ctx context.Context, req *model.Course) error {

	_, err := rr.db.ExecContext(ctx, insertCourses, req.UID, req.Subject)
	if err != nil && !strings.Contains(err.Error(), "duplicate") {
		return err
	}

	return nil
}
