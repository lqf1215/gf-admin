package consts

import (
	"gf-admin/internal/co_interface"
	"gf-admin/internal/dao"
	model "gf-admin/internal/model"
	"gf-admin/internal/module"
	"gf-admin/internal/premission"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/utility/permission"
	cm_co_interface "github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_module"
)

type global struct {
	CompanyModules cm_co_interface.IModules
	ServiceModules co_interface.IModules
}

var (
	PermissionTree []*permission.SysPermissionTree

	// FinancialPermissionTree 财务服务权限树 (可选)
	FinancialPermissionTree []*permission.SysPermissionTree
	// DeptPermissionTree 部门 权限树 (可选)
	DeptPermissionTree []*permission.SysPermissionTree

	Global = global{
		CompanyModules: co_module.NewModules(
			&co_model.Config{
				AllowEmptyNo:                   true,
				IsCreateDefaultEmployeeAndRole: false,
				HardDeleteWaitAt:               0,
				KeyIndex:                       "Company",
				RoutePrefix:                    "/company",
				StoragePath:                    "./resources/company",
				UserType:                       sys_enum.User.Type.SuperAdmin,
				Identifier: co_model.Identifier{
					Company:  "company",
					Employee: "employee",
					Team:     "team",
				},
			},
			&co_dao.XDao{ // 以下为业务层实例化dao模型，如果不是使用默认模型时需要将自定义dao模型作为参数传进去
				Company:    co_dao.NewCompany(),
				Team:       co_dao.NewCompanyTeam(),
				Employee:   co_dao.NewCompanyEmployee(),
				TeamMember: co_dao.NewCompanyTeamMember(),
			},
		),

		ServiceModules: initServiceModules(),
	}
)

func initServiceModules() co_interface.IModules {
	modules := module.NewModules(
		&model.Config{
			AllowEmptyNo:                   true,
			IsCreateDefaultEmployeeAndRole: false,
			HardDeleteWaitAt:               0,
			KeyIndex:                       "Company",
			RoutePrefix:                    "/company",
			StoragePath:                    "./resources/company",
			UserType:                       sys_enum.User.Type.SuperAdmin,
			Identifier: model.Identifier{
				Dept: "dept",
				Post: "post",
			},
		},
		&dao.XDao{ // 以下为业务层实例化dao模型，如果不是使用默认模型时需要将自定义dao模型作为参数传进去
			Dept: dao.NewCompanyDept(),
			Post: dao.NewCompanyPost(),
		},
	)
	// 权限树追加权限
	PermissionTree = append(PermissionTree, premission.InitPermission(modules)...)
	return modules
}
