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

type PostController struct {
	modules co_interface.IModules
}

var Post = func(modules co_interface.IModules) i_controller.IPost {
	return &PostController{
		modules: modules,
	}
}

func (c *PostController) GetModules() co_interface.IModules {
	return c.modules
}

func (c *PostController) GetPostById(ctx context.Context, req *co_v1.GetPostByIdReq) (*model.PostRes, error) {
	return funs.CheckPermission(ctx,
		func() (*model.PostRes, error) {
			return c.modules.Post().GetPostById(ctx, req.Id)
		},
		enum.Post.PermissionType(c.modules).ViewDetail,
	)
}

// HasPostByName 岗位名称是否存在
func (c *PostController) HasPostByName(ctx context.Context, req *co_v1.HasPostByNameReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.modules.Post().HasPostByName(ctx, req.Name) == true, nil
		},
	)
}

// QueryPostList 查询岗位列表
func (c *PostController) QueryPostList(ctx context.Context, req *co_v1.QueryPostListReq) (*model.PostListRes, error) {
	return funs.CheckPermission(ctx,
		func() (*model.PostListRes, error) {
			return c.modules.Post().QueryPostList(ctx, &req.SearchParams)
		},
		enum.Post.PermissionType(c.modules).List,
	)
}

// CreatePost 创建岗位信息
func (c *PostController) CreatePost(ctx context.Context, req *co_v1.CreatePostReq) (*model.PostRes, error) {
	//req.UnionMainId = sys_service.SysSession().Get(ctx).JwtClaimsUser.UnionMainId

	return funs.CheckPermission(ctx,
		func() (*model.PostRes, error) {
			ret, err := c.modules.Post().CreatePost(ctx, &req.Post)
			return ret, err
		},
		enum.Post.PermissionType(c.modules).Create,
	)
}

// UpdatePost 更新岗位信息
func (c *PostController) UpdatePost(ctx context.Context, req *co_v1.UpdatePostReq) (*model.PostRes, error) {
	return funs.CheckPermission(ctx,
		func() (*model.PostRes, error) {
			ret, err := c.modules.Post().UpdatePost(ctx, &req.Post)
			return ret, err
		},
		enum.Post.PermissionType(c.modules).Update,
	)
}

// DeletePost 删除岗位信息
func (c *PostController) DeletePost(ctx context.Context, req *co_v1.DeletePostReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := c.modules.Post().DeletePost(ctx, req.Id)
			return ret == true, err
		},
		enum.Post.PermissionType(c.modules).Delete,
	)
}
