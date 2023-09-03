package usecase

import (
	"context"
	"github/yogabagas/join-app/domain/model"
	"github/yogabagas/join-app/domain/service"
	coursesRepo "github/yogabagas/join-app/service/courses/repository"
	"github/yogabagas/join-app/shared/util"
)

type CoursesServiceImpl struct {
	coursesRepo coursesRepo.CoursesRepository
}

type CoursesService interface {
	CreateCourses(ctx context.Context, req service.CreateCoursesReq) error
}

func NewCoursesService(coursesRepo coursesRepo.CoursesRepository) CoursesService {
	return &CoursesServiceImpl{
		coursesRepo: coursesRepo}
}

func (cs *CoursesServiceImpl) CreateCourses(ctx context.Context, req service.CreateCoursesReq) error {

	uID := util.NewULIDGenerate()

	return cs.coursesRepo.CreateCourses(ctx, &model.Course{
		UID:     uID,
		Subject: req.Subject,
	})
}
