package i_controller

import (
	"context"
	"gf-admin/api/co_v1"
	model "gf-admin/internal/model"
	"github.com/SupenBysz/gf-admin-community/api_v1"
)

type IDept interface {
	iModule
	// GetDeptById 根据id获取部门信息
	GetDeptById(ctx context.Context, req *co_v1.GetDeptByIdReq) (*model.DeptRes, error)

	// HasDeptByName 部门名称是否存在
	HasDeptByName(ctx context.Context, req *co_v1.HasDeptByNameReq) (api_v1.BoolRes, error)

	// QueryDeptList 查询部门列表
	QueryDeptList(ctx context.Context, req *co_v1.QueryDeptListReq) (*model.DeptListRes, error)

	// CreateDept 创建部门信息
	CreateDept(ctx context.Context, req *co_v1.CreateDeptReq) (*model.DeptRes, error)

	// UpdateDept 更新部门信息
	UpdateDept(ctx context.Context, req *co_v1.UpdateDeptReq) (*model.DeptRes, error)

	// DeleteDept 删除部门信息
	DeleteDept(ctx context.Context, req *co_v1.DeleteDeptReq) (api_v1.BoolRes, error)

	// GetDeptListTree 获取部门树形结构
	GetDeptListTree(ctx context.Context, req *co_v1.GetDeptListTreeByIdReq) (*model.DeptTreeSelectRes, error)
}
