package internal

import (
	"context"
	"gf-admin/api/co_v1"
	"gf-admin/internal/co_interface"
	"gf-admin/internal/co_interface/i_controller"
	enum "gf-admin/internal/enum"
	model "gf-admin/internal/model"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/utility/funs"
)

type DeptController struct {
	modules co_interface.IModules
}

var Dept = func(modules co_interface.IModules) i_controller.IDept {
	return &DeptController{
		modules: modules,
	}
}

func (c *DeptController) GetModules() co_interface.IModules {
	return c.modules
}

func (c *DeptController) GetDeptById(ctx context.Context, req *co_v1.GetDeptByIdReq) (*model.DeptRes, error) {
	return funs.CheckPermission(ctx,
		func() (*model.DeptRes, error) {
			return c.modules.Dept().GetDeptById(ctx, req.Id)
		},
		enum.Dept.PermissionType(c.modules).ViewDetail,
	)
}

// HasDeptByName 部门名称是否存在
func (c *DeptController) HasDeptByName(ctx context.Context, req *co_v1.HasDeptByNameReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.modules.Dept().HasDeptByName(ctx, req.Name) == true, nil
		},
	)
}

// QueryDeptList 查询部门列表
func (c *DeptController) QueryDeptList(ctx context.Context, req *co_v1.QueryDeptListReq) (*model.DeptListRes, error) {
	return funs.CheckPermission(ctx,
		func() (*model.DeptListRes, error) {
			return c.modules.Dept().QueryDeptList(ctx, &req.SearchParams)
		},
		enum.Dept.PermissionType(c.modules).List,
	)
}

// CreateDept 创建部门信息
func (c *DeptController) CreateDept(ctx context.Context, req *co_v1.CreateDeptReq) (*model.DeptRes, error) {
	//req.UnionMainId = sys_service.SysSession().Get(ctx).JwtClaimsUser.UnionMainId

	return funs.CheckPermission(ctx,
		func() (*model.DeptRes, error) {
			ret, err := c.modules.Dept().CreateDept(ctx, &req.Dept)
			return ret, err
		},
		enum.Dept.PermissionType(c.modules).Create,
	)
}

// UpdateDept 更新部门信息
func (c *DeptController) UpdateDept(ctx context.Context, req *co_v1.UpdateDeptReq) (*model.DeptRes, error) {
	return funs.CheckPermission(ctx,
		func() (*model.DeptRes, error) {
			ret, err := c.modules.Dept().UpdateDept(ctx, &req.Dept)
			return ret, err
		},
		enum.Dept.PermissionType(c.modules).Update,
	)
}

// DeleteDept 删除部门信息
func (c *DeptController) DeleteDept(ctx context.Context, req *co_v1.DeleteDeptReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := c.modules.Dept().DeleteDept(ctx, req.Id)
			return ret == true, err
		},
		enum.Dept.PermissionType(c.modules).Delete,
	)
}

// GetDeptListTree 获取部门数据结构数据
func (c *DeptController) GetDeptListTree(ctx context.Context, req *co_v1.GetDeptListTreeByIdReq) (res *model.DeptTreeSelectRes, err error) {
	return funs.CheckPermission(ctx,
		func() (*model.DeptTreeSelectRes, error) {
			ret, err := c.modules.Dept().GetDeptListTree(ctx, req.Id)
			return ret, err
		},
		enum.Dept.PermissionType(c.modules).ListTree,
	)
}
