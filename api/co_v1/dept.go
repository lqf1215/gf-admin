package co_v1

import (
	model "gf-admin/internal/model"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/frame/g"
)

type GetDeptByIdReq struct {
	g.Meta `method:"post" summary:"根据ID获取部门|信息" tags:"部门"`
	Id     int64 `json:"id" v:"required#ID校验失败" dc:"部门ID"`
}

type HasDeptByNameReq struct {
	g.Meta `method:"post" summary:"判断名称是否存在" tags:"部门"`
	Name   string `json:"name" v:"required#名称不能为空" dc:"名称"`
}

type QueryDeptListReq struct {
	g.Meta `method:"post" summary:"查询部门|列表" tags:"部门"`
	sys_model.SearchParams
}

type CreateDeptReq struct {
	g.Meta `method:"post" summary:"创建部门|信息" tags:"部门"`
	model.Dept
}

type UpdateDeptReq struct {
	g.Meta `method:"post" summary:"更新部门|信息" tags:"部门"`
	model.Dept
}

type DeleteDeptReq struct {
	g.Meta `method:"post" summary:"删除部门|信息" tags:"部门"`
	Id     int64 `json:"id" v:"required#ID校验失败" dc:"部门ID"`
}

type GetDeptListTreeByIdReq struct {
	g.Meta `method:"post" summary:"查询部门|树形列表" tags:"部门"`
	Id     int64 `json:"id"  dc:"部门ID 查询部门的下级"`
}
