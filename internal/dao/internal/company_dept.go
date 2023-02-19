// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl/dao_interface"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CompanyDeptDao is the data access object for table co_company_dept.
type CompanyDeptDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns CompanyDeptColumns // columns contains all the column names of Table for convenient usage.
}

// CompanyDeptColumns defines and stores column names for table co_company_dept.
type CompanyDeptColumns struct {
	Id        string // ID
	ParentId  string // 父部门id
	Ancestors string // 祖级列表
	Name      string // 部门名称
	Email     string // 邮箱
	State     string // 部门状态（0正常 1停用）
	Sort      string // 显示顺序
	Mobile    string // 联系电话
	Leader    string // 负责人
	CreatedBy string // 创建人
	UpdatedBy string // 修改人
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// companyDeptColumns holds the columns for table co_company_dept.
var companyDeptColumns = CompanyDeptColumns{
	Id:        "id",
	ParentId:  "parent_id",
	Ancestors: "ancestors",
	Name:      "name",
	Email:     "email",
	State:     "state",
	Sort:      "sort",
	Mobile:    "mobile",
	Leader:    "leader",
	CreatedBy: "created_by",
	UpdatedBy: "updated_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewCompanyDeptDao creates and returns a new DAO object for table data access.
func NewCompanyDeptDao(proxy ...dao_interface.IDao) *CompanyDeptDao {
	var dao *CompanyDeptDao
	if len(proxy) > 0 {
		dao = &CompanyDeptDao{
			group:   proxy[0].Group(),
			table:   proxy[0].Table(),
			columns: companyDeptColumns,
		}
		return dao
	}

	return &CompanyDeptDao{
		group:   "default",
		table:   "co_company_dept",
		columns: companyDeptColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CompanyDeptDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CompanyDeptDao) Table() string {
	return dao.table
}

// Group returns the configuration group name of database of current dao.
func (dao *CompanyDeptDao) Group() string {
	return dao.group
}

// Columns returns all column names of current dao.
func (dao *CompanyDeptDao) Columns() CompanyDeptColumns {
	return dao.columns
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CompanyDeptDao) Ctx(ctx context.Context, cacheOption ...*gdb.CacheOption) *gdb.Model {
	return dao.DaoConfig(ctx, cacheOption...).Model
}

func (dao *CompanyDeptDao) DaoConfig(ctx context.Context, cacheOption ...*gdb.CacheOption) dao_interface.DaoConfig {
	daoConfig := dao_interface.DaoConfig{
		Dao:   dao,
		DB:    dao.DB(),
		Table: dao.table,
		Group: dao.group,
		Model: dao.DB().Model(dao.Table()).Safe().Ctx(ctx),
	}

	if len(cacheOption) == 0 {
		daoConfig.CacheOption = daoctl.MakeDaoCache(dao.Table())
		daoConfig.Model = daoConfig.Model.Cache(*daoConfig.CacheOption)
	} else {
		if cacheOption[0] != nil {
			daoConfig.CacheOption = cacheOption[0]
			daoConfig.Model = daoConfig.Model.Cache(*daoConfig.CacheOption)
		}
	}

	daoConfig.Model = daoctl.RegisterDaoHook(daoConfig.Model)

	return daoConfig
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CompanyDeptDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
