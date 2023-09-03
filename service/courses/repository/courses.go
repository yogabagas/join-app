package repository

import (
	"context"
	"github/yogabagas/join-app/domain/model"
)

type CoursesRepository interface {
	CreateCourses(ctx context.Context, req *model.Course) error
}
