package controller

import (
	"gf-admin/internal/co_interface/i_controller"
	"gf-admin/internal/controller/internal"
)

type (
	DeptController internal.DeptController
	PostController internal.PostController
)

type CoController struct {
	Dept i_controller.IDept
	Post i_controller.IPost
}

var (
	Dept = internal.Dept
	Post = internal.Post
)
