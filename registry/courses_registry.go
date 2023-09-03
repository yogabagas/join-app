package registry

import (
	"github/yogabagas/join-app/adapter/controller"
	"github/yogabagas/join-app/domain/repository/sql"
	"github/yogabagas/join-app/service/courses/repository"
	"github/yogabagas/join-app/service/courses/usecase"
)

func (m *module) NewCoursesRepository() repository.CoursesRepository {
	return sql.NewCoursesRepository(m.sqlDB)
}

func (m *module) NewCoursesRegistry() usecase.CoursesService {
	return usecase.NewCoursesService(m.NewCoursesRepository())
}

func (m *module) NewCoursesController() controller.CoursesController {
	return controller.NewCoursesController(m.NewCoursesRegistry())
}
