package co_interface

import (
	"context"
	"gf-admin/internal/dao"
	model "gf-admin/internal/model"
	"gf-admin/internal/model/entity"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/i18n/gi18n"
)

type (
	IDept interface {
		GetDeptById(ctx context.Context, id int64) (*model.DeptRes, error)
		GetDeptByName(ctx context.Context, name string) (*model.DeptRes, error)
		HasDeptByName(ctx context.Context, name string, excludeIds ...int64) bool
		QueryDeptList(ctx context.Context, filter *sys_model.SearchParams) (*model.DeptListRes, error)
		CreateDept(ctx context.Context, info *model.Dept) (*model.DeptRes, error)
		UpdateDept(ctx context.Context, info *model.Dept) (*model.DeptRes, error)
		DeleteDept(ctx context.Context, id int64) (bool, error)
		FindSonByParentId(deptList []*entity.CompanyDept, deptId int64) []*entity.CompanyDept
		GetListTree(pid int64, list []*model.Dept) (deptTree []*model.DeptTreeRes)
		GetDeptListTree(ctx context.Context, parentId int64) (*model.DeptTreeSelectRes, error)
	}

	IPost interface {
		GetPostById(ctx context.Context, id int64) (*model.PostRes, error)
		GetPostByName(ctx context.Context, name string) (*model.PostRes, error)
		HasPostByName(ctx context.Context, name string, excludeIds ...int64) bool
		QueryPostList(ctx context.Context, filter *sys_model.SearchParams) (*model.PostListRes, error)
		CreatePost(ctx context.Context, info *model.Post) (*model.PostRes, error)
		UpdatePost(ctx context.Context, info *model.Post) (*model.PostRes, error)
		DeletePost(ctx context.Context, id int64) (bool, error)
	}
)

type IModules interface {
	GetConfig() *model.Config
	Dept() IDept
	Post() IPost

	SetI18n(i18n *gi18n.Manager) error
	T(ctx context.Context, content string) string
	Tf(ctx context.Context, format string, values ...interface{}) string
	Dao() *dao.XDao
}
