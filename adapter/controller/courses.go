package controller

import (
	"context"
	"github/yogabagas/join-app/domain/service"
	"github/yogabagas/join-app/service/courses/usecase"
)

type CoursesControllerImpl struct {
	coursesSvc usecase.CoursesService
}

type CoursesController interface {
	CreateCourses(ctx context.Context, req service.CreateCoursesReq) error
}

func NewCoursesController(coursesService usecase.CoursesService) CoursesController {
	return &CoursesControllerImpl{coursesSvc: coursesService}
}

func (cs *CoursesControllerImpl) CreateCourses(ctx context.Context, req service.CreateCoursesReq) error {
	return cs.coursesSvc.CreateCourses(ctx, req)
}
