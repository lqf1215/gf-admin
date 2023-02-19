package co_model

import (
	"gf-admin/internal/model/entity"
	"github.com/SupenBysz/gf-admin-community/sys_model"
)

type Dept struct {
	Id       int64  `json:"id"           description:"ID，保持与USERID一致"`
	ParentID int64  `p:"parentId"  v:"required#父级不能为空" description:"父部门id"`
	Name     string `p:"name"  v:"required#部门名称不能为空" description:"部门名称"`
	Sort     int    `p:"sort"  v:"required#排序不能为空"  description:"显示顺序"`
	Leader   string `p:"leader" description:"负责人"`
	Mobile   string `p:"mobile"    description:"联系电话"`
	Email    string `p:"email"  v:"email#邮箱格式不正确"  description:"邮箱"`
	State    int    `p:"state"  v:"required#状态必须"  description:"部门状态（0正常 1停用）"`
}

type DeptRes struct {
	*entity.CompanyDept
}

type DeptSearchRes struct {
	DeptList []*entity.CompanyDept `json:"deptList"`
}
type DeptTreeRes struct {
	*Dept
	Children []*DeptTreeRes `json:"children"`
}

type DeptTreeSelectRes struct {
	Deps []*DeptTreeRes `json:"deps"`
}

type DeptListRes sys_model.CollectRes[*DeptRes]
