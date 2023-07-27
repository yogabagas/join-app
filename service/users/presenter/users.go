package presenter

import (
	"context"
	"fmt"
	"github/yogabagas/join-app/domain/model"
	"github/yogabagas/join-app/domain/service"
	"github/yogabagas/join-app/shared/util"
	"time"
)

type UsersPresenterImpl struct{}

type UsersPresenter interface {
	GetUsersWithPagination(ctx context.Context, req service.GetUsersWithPaginationReq, userResp *model.ReadUsersWithPaginationResp,
		count *model.CountUsersResp) (service.GetUsersWithPaginationResp, error)
}

func NewUsersPresenter() UsersPresenter {
	return &UsersPresenterImpl{}
}

func (up *UsersPresenterImpl) GetUsersWithPagination(ctx context.Context, req service.GetUsersWithPaginationReq,
	userResp *model.ReadUsersWithPaginationResp, count *model.CountUsersResp) (resp service.GetUsersWithPaginationResp, err error) {

	if userResp != nil {
		for _, v := range userResp.Users {

			user := service.UserResp{
				Fullname:  fmt.Sprintf("%s %s", v.FirstName, v.LastName),
				Username:  v.Username,
				Birthdate: v.Birthdate.Format(time.DateOnly),
				Email:     v.Email,
				Role:      v.RoleName,
			}
			resp.Users = append(resp.Users, user)
		}

		resp.Pagination.Page = req.Page
		resp.Pagination.PerPage = userResp.PerPage
		resp.Pagination.TotalData = count.Total
		resp.Pagination.TotalPage = util.GetTotalPage(count.Total, userResp.PerPage)
	}

	return

}
