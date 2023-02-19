// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	model "gf-admin/internal/model"
	"gf-admin/internal/model/entity"

	"github.com/SupenBysz/gf-admin-community/sys_model"
)

type (
	IDept interface {
		GetDeptById(ctx context.Context, id int64) (*model.DeptRes, error)
		GetDeptByName(ctx context.Context, name string) (*model.DeptRes, error)
		HasDeptByName(ctx context.Context, name string, excludeIds ...int64) bool
		QueryDeptList(ctx context.Context, filter *sys_model.SearchParams) (*model.DeptListRes, error)
		CreateDept(ctx context.Context, info *model.Dept) (*model.DeptRes, error)
		UpdateDept(ctx context.Context, info *model.Dept) (*model.DeptRes, error)
		DeleteDept(ctx context.Context, id int64) (bool, error)
		FindSonByParentId(deptList []*entity.CompanyDept, deptId int64) []*entity.CompanyDept
		GetListTree(pid int64, list []*model.Dept) (deptTree []*model.DeptTreeRes)
		GetDeptListTree(ctx context.Context, parentId int64) (*model.DeptTreeSelectRes, error)
	}
)

var (
	localDept IDept
)

func Dept() IDept {
	if localDept == nil {
		panic("implement not found for interface IDept, forgot register?")
	}
	return localDept
}

func RegisterDept(i IDept) {
	localDept = i
}