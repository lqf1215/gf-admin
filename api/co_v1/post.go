package co_v1

import (
	model "gf-admin/internal/model"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/frame/g"
)

type GetPostByIdReq struct {
	g.Meta `method:"post" summary:"根据ID获取岗位|信息" tags:"岗位"`
	Id     int64 `json:"id" v:"required#ID校验失败" dc:"岗位ID"`
}

type HasPostByNameReq struct {
	g.Meta `method:"post" summary:"判断名称是否存在" tags:"岗位"`
	Name   string `json:"name" v:"required#名称不能为空" dc:"名称"`
}

type QueryPostListReq struct {
	g.Meta `method:"post" summary:"查询岗位|列表" tags:"岗位"`
	sys_model.SearchParams
}

type CreatePostReq struct {
	g.Meta `method:"post" summary:"创建岗位|信息" tags:"岗位"`
	model.Post
}

type UpdatePostReq struct {
	g.Meta `method:"post" summary:"更新岗位|信息" tags:"岗位"`
	model.Post
}

type DeletePostReq struct {
	g.Meta `method:"post" summary:"删除岗位|信息" tags:"岗位"`
	Id     int64 `json:"id" v:"required#ID校验失败" dc:"岗位ID"`
}
