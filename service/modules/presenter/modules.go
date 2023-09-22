package presenter

import (
	"context"
	"github/yogabagas/join-app/domain/model"
	"github/yogabagas/join-app/domain/service"
	"github/yogabagas/join-app/shared/util"
)

type ModulesPresenterImpl struct{}

type ModulesPresenter interface {
	GetModulesWithPagination(ctx context.Context, req service.GetModulesWithPaginationReq, userResp *model.ReadModulesWithPaginationResp,
		count *model.CountModulesResp) (resp service.GetModulesWithPaginationResp, err error)
}

func NewModulesPresenter() ModulesPresenter {
	return &ModulesPresenterImpl{}
}

func (up *ModulesPresenterImpl) GetModulesWithPagination(ctx context.Context, req service.GetModulesWithPaginationReq, modulesResp *model.ReadModulesWithPaginationResp,
	count *model.CountModulesResp) (resp service.GetModulesWithPaginationResp, err error) {

	if modulesResp != nil {
		for _, v := range modulesResp.Modules {

			module := service.ModuleResp{
				Name:        v.Name,
				Description: v.Description,
				File:        v.File,
			}
			resp.Modules = append(resp.Modules, module)
		}

		resp.Pagination.Page = req.Page
		resp.Pagination.PerPage = modulesResp.PerPage
		resp.Pagination.TotalData = count.Total
		resp.Pagination.TotalPage = util.GetTotalPage(count.Total, modulesResp.PerPage)
	}

	return

}
