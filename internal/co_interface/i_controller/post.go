package i_controller

import (
	"context"
	"gf-admin/api/co_v1"
	model "gf-admin/internal/model"
	"github.com/SupenBysz/gf-admin-community/api_v1"
)

type IPost interface {
	iModule
	// GetPostById 根据id获取岗位信息
	GetPostById(ctx context.Context, req *co_v1.GetPostByIdReq) (*model.PostRes, error)

	// HasPostByName 岗位名称是否存在
	HasPostByName(ctx context.Context, req *co_v1.HasPostByNameReq) (api_v1.BoolRes, error)

	// QueryPostList 查询岗位列表
	QueryPostList(ctx context.Context, req *co_v1.QueryPostListReq) (*model.PostListRes, error)

	// CreatePost 创建岗位信息
	CreatePost(ctx context.Context, req *co_v1.CreatePostReq) (*model.PostRes, error)

	// UpdatePost 更新岗位信息
	UpdatePost(ctx context.Context, req *co_v1.UpdatePostReq) (*model.PostRes, error)

	// DeletePost 删除岗位信息
	DeletePost(ctx context.Context, req *co_v1.DeletePostReq) (api_v1.BoolRes, error)
}
