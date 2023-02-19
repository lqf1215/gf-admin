package co_enum

import (
	"gf-admin/internal/enum/internal/dept"
	"gf-admin/internal/enum/internal/post"
)

type (
	DeptPermissionType = dept.PermissionEnum
	DeptState          = dept.StateEnum

	PostPermissionType = post.PermissionEnum
	PostState          = post.StateEnum
)

var (
	Dept = dept.Dept
	Post = post.Post
)
