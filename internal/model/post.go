package co_model

import (
	"gf-admin/internal/model/entity"
	"github.com/SupenBysz/gf-admin-community/sys_model"
)

type Post struct {
	Id     int64  `json:"id"           description:"ID，保持与USERID一致"`
	Code   string `p:"code" v:"required#岗位编码不能为空" ds:"岗位编码"`
	Name   string `p:"name" v:"required#岗位名称不能为空" ds:"岗位名称"`
	Sort   int    `p:"sort" v:"required#岗位排序不能为空"`
	Remark string `p:"remark"`
	State  int    `p:"state"  v:"required#状态必须"  description:"状态（0正常 1停用）"`
}

type PostRes struct {
	*entity.CompanyPost
}

type PostSearchRes struct {
	PostList []*entity.CompanyPost `json:"PostList"`
}

type PostListRes sys_model.CollectRes[*PostRes]
