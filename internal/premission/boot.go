package premission

import (
	"context"
	"gf-admin/internal/co_interface"
	enum "gf-admin/internal/enum"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/utility/permission"
)

// InitPermission 初始化权限树
func InitPermission(module co_interface.IModules) []*permission.SysPermissionTree {
	result := []*permission.SysPermissionTree{

		// 部门
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5948221667408321,
				Name:       module.T(context.TODO(), "{#CompanyName}{#DeptName}"),
				Identifier: module.GetConfig().Identifier.Dept,
				Type:       1,
				IsShow:     1,
			},
			Children: []*permission.SysPermissionTree{
				enum.Dept.PermissionType(module).ViewDetail,
				enum.Dept.PermissionType(module).List,
				enum.Dept.PermissionType(module).Create,
				enum.Dept.PermissionType(module).Update,
				enum.Dept.PermissionType(module).Delete,
				enum.Dept.PermissionType(module).SetState,
				enum.Dept.PermissionType(module).ListTree,
			},
		},

		// 部门
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5948221667408351,
				Name:       module.T(context.TODO(), "{#CompanyName}{#PostName}"),
				Identifier: module.GetConfig().Identifier.Post,
				Type:       1,
				IsShow:     1,
			},
			Children: []*permission.SysPermissionTree{
				enum.Post.PermissionType(module).ViewDetail,
				enum.Post.PermissionType(module).List,
				enum.Post.PermissionType(module).Create,
				enum.Post.PermissionType(module).Update,
				enum.Post.PermissionType(module).Delete,
				enum.Post.PermissionType(module).SetState,
			},
		},
	}
	return result
}
