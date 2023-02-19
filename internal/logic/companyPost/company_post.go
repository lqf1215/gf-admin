package companyPost

import (
	"context"
	"fmt"
	"gf-admin/internal/co_interface"
	"gf-admin/internal/dao"
	model "gf-admin/internal/model"
	"gf-admin/internal/model/do"
	"gf-admin/internal/model/entity"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/yitter/idgenerator-go/idgen"
)

type sPost struct {
	modules co_interface.IModules
	dao     *dao.XDao
}

func NewPost(modules co_interface.IModules) co_interface.IPost {
	return &sPost{
		modules: modules,
		dao:     modules.Dao(),
	}
}

// GetPostById 根据ID获取获取岗位信息
func (s *sPost) GetPostById(ctx context.Context, id int64) (*model.PostRes, error) {
	data, err := daoctl.GetByIdWithError[model.PostRes](
		s.dao.Post.Ctx(ctx),
		id,
	)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#PostName} {#error_Data_NotFound}"), s.dao.Post.Table())
	}

	return data, nil
}

// GetPostByName 根据Name获取获取岗位信息
func (s *sPost) GetPostByName(ctx context.Context, name string) (*model.PostRes, error) {
	data, err := daoctl.ScanWithError[model.PostRes](
		s.dao.Post.Ctx(ctx).
			Where(do.CompanyPost{Name: name}),
	)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#PostName} {#error_Data_NotFound}"), s.dao.Post.Table())
	}

	return data, nil
}

// HasPostByName 判断名称是否存在
func (s *sPost) HasPostByName(ctx context.Context, name string, excludeIds ...int64) bool {
	model := s.dao.Post.Ctx(ctx).Where(do.CompanyPost{
		Name: name,
	})

	if len(excludeIds) > 0 {
		var ids []int64
		for _, id := range excludeIds {
			if id > 0 {
				ids = append(ids, id)
			}
		}
		if len(ids) > 0 {
			model = model.WhereNotIn(s.dao.Post.Columns().Id, ids)
		}
	}

	count, _ := model.Count()
	return count > 0
}

// QueryPostList 查询岗位列表
func (s *sPost) QueryPostList(ctx context.Context, filter *sys_model.SearchParams) (*model.PostListRes, error) {

	data, err := daoctl.Query[*model.PostRes](
		s.dao.Post.Ctx(ctx),
		filter,
		false,
	)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#PostName} {#error_Data_NotFound}"), s.dao.Post.Table())
	}

	return (*model.PostListRes)(data), nil
}

// CreatePost 创建岗位信息
func (s *sPost) CreatePost(ctx context.Context, info *model.Post) (*model.PostRes, error) {
	info.Id = 0
	return s.savePost(ctx, info)
}

// UpdatePost 更新岗位信息
func (s *sPost) UpdatePost(ctx context.Context, info *model.Post) (*model.PostRes, error) {
	if info.Id <= 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "{#PostName} {#error_Data_NotFound}"), s.dao.Post.Table())
	}
	return s.savePost(ctx, info)
}

// SavePost 保存岗位信息
func (s *sPost) savePost(ctx context.Context, info *model.Post) (*model.PostRes, error) {
	// 名称重名检测
	if s.HasPostByName(ctx, info.Name, info.Id) {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, fmt.Sprintf("{#PostName} %v {#error_NameAlreadyExists}", info.Name)), s.dao.Post.Table())
	}

	// 获取登录用户
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 构建岗位ID
	UnionMainId := idgen.NextId()

	data := do.CompanyPost{
		Meta:   g.Meta{},
		Name:   info.Name,
		State:  info.State,
		Sort:   info.Sort,
		Remark: info.Remark,
		Code:   info.Code,
	}

	// 启用事务
	err := s.dao.Post.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if info.Id == 0 {
			data.Id = UnionMainId
			data.CreatedBy = sessionUser.Username
			data.CreatedAt = gtime.Now()

			affected, err := daoctl.InsertWithError(
				s.dao.Post.Ctx(ctx),
				data,
			)
			if affected == 0 || err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#PostName} {#error_Data_Save_Failed}"), s.dao.Post.Table())
			}
		} else {
			data.Id = info.Id
			data.UpdatedBy = sessionUser.Username
			data.UpdatedAt = gtime.Now()
			_, err = daoctl.UpdateWithError(
				s.dao.Post.Ctx(ctx).
					Where(do.CompanyPost{Id: info.Id}),
				data,
			)
			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#PostName} {#error_Data_Save_Failed}"), s.dao.Post.Table())
			}

		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return s.GetPostById(ctx, data.Id.(int64))
}

// DeletePost 删除岗位
func (s *sPost) DeletePost(ctx context.Context, id int64) (bool, error) {

	info := entity.CompanyPost{}
	err := s.dao.Post.Ctx(ctx).
		Where(do.CompanyPost{Id: id}).Scan(&info)

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#PostName}{#error_Data_Get_Failed}"), s.dao.Post.Table())

	}

	if info.Id <= 0 {
		return true, nil
	}

	_, err = dao.NewCompanyPost().Ctx(ctx).
		Delete(do.CompanyPost{Id: id})

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#PostName}{#error_Delete_Failed}"), s.dao.Post.Table())

	}

	return true, err
}
