package companyDept

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

type sDept struct {
	modules co_interface.IModules
	dao     *dao.XDao
}

func NewDept(modules co_interface.IModules) co_interface.IDept {
	return &sDept{
		modules: modules,
		dao:     modules.Dao(),
	}
}

// GetDeptById 根据ID获取获取部门信息
func (s *sDept) GetDeptById(ctx context.Context, id int64) (*model.DeptRes, error) {
	data, err := daoctl.GetByIdWithError[model.DeptRes](
		s.dao.Dept.Ctx(ctx),
		id,
	)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#DeptName} {#error_Data_NotFound}"), s.dao.Dept.Table())
	}

	return data, nil
}

// GetDeptByName 根据Name获取获取部门信息
func (s *sDept) GetDeptByName(ctx context.Context, name string) (*model.DeptRes, error) {
	data, err := daoctl.ScanWithError[model.DeptRes](
		s.dao.Dept.Ctx(ctx).
			Where(do.CompanyDept{Name: name}),
	)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#DeptName} {#error_Data_NotFound}"), s.dao.Dept.Table())
	}

	return data, nil
}

// HasDeptByName 判断名称是否存在
func (s *sDept) HasDeptByName(ctx context.Context, name string, excludeIds ...int64) bool {
	model := s.dao.Dept.Ctx(ctx).Where(do.CompanyDept{
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
			model = model.WhereNotIn(s.dao.Dept.Columns().Id, ids)
		}
	}

	count, _ := model.Count()
	return count > 0
}

// QueryDeptList 查询部门列表
func (s *sDept) QueryDeptList(ctx context.Context, filter *sys_model.SearchParams) (*model.DeptListRes, error) {

	data, err := daoctl.Query[*model.DeptRes](
		s.dao.Dept.Ctx(ctx),
		filter,
		false,
	)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#DeptName} {#error_Data_NotFound}"), s.dao.Dept.Table())
	}

	return (*model.DeptListRes)(data), nil
}

// CreateDept 创建部门信息
func (s *sDept) CreateDept(ctx context.Context, info *model.Dept) (*model.DeptRes, error) {
	info.Id = 0
	return s.saveDept(ctx, info)
}

// UpdateDept 更新部门信息
func (s *sDept) UpdateDept(ctx context.Context, info *model.Dept) (*model.DeptRes, error) {
	if info.Id <= 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "{#DeptName} {#error_Data_NotFound}"), s.dao.Dept.Table())
	}
	return s.saveDept(ctx, info)
}

// SaveDept 保存部门信息
func (s *sDept) saveDept(ctx context.Context, info *model.Dept) (*model.DeptRes, error) {
	// 名称重名检测
	if s.HasDeptByName(ctx, info.Name, info.Id) {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, fmt.Sprintf("{#DeptName} %v {#error_NameAlreadyExists}", info.Name)), s.dao.Dept.Table())
	}

	// 获取登录用户
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 构建部门ID
	UnionMainId := idgen.NextId()

	data := do.CompanyDept{
		Meta:     g.Meta{},
		ParentId: info.ParentID,
		Name:     info.Name,
		Email:    info.Email,
		State:    info.State,
		Sort:     info.Sort,
		Mobile:   info.Mobile,
		Leader:   info.Leader,
	}

	// 启用事务
	err := s.dao.Dept.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if info.Id == 0 {
			data.Id = UnionMainId
			data.CreatedBy = sessionUser.Username
			data.CreatedAt = gtime.Now()

			affected, err := daoctl.InsertWithError(
				s.dao.Dept.Ctx(ctx),
				data,
			)
			if affected == 0 || err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#DeptName} {#error_Data_Save_Failed}"), s.dao.Dept.Table())
			}
		} else {
			data.Id = info.Id
			data.UpdatedBy = sessionUser.Username
			data.UpdatedAt = gtime.Now()
			_, err = daoctl.UpdateWithError(
				s.dao.Dept.Ctx(ctx).
					Where(do.CompanyDept{Id: info.Id}),
				data,
			)
			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#DeptName} {#error_Data_Save_Failed}"), s.dao.Dept.Table())
			}

		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return s.GetDeptById(ctx, data.Id.(int64))
}

// DeleteDept 删除部门
func (s *sDept) DeleteDept(ctx context.Context, id int64) (bool, error) {
	count, err := s.dao.Dept.Ctx(ctx).Where(do.CompanyDept{ParentId: id}).Count()

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#DeptName}{#error_Delete_Failed}"), s.dao.Dept.Table())

	}

	if count > 0 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#DeptName}{#error_Not_Delete_ExistParentId}"), s.dao.Dept.Table())

	}

	info := entity.CompanyDept{}
	err = s.dao.Dept.Ctx(ctx).
		Where(do.CompanyDept{Id: id}).Scan(&info)

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#DeptName}{#error_Data_Get_Failed}"), s.dao.Dept.Table())

	}

	if info.Id <= 0 {
		return true, nil
	}

	_, err = dao.NewCompanyDept().Ctx(ctx).
		Delete(do.CompanyDept{Id: id})

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#DeptName}{#error_Delete_Failed}"), s.dao.Dept.Table())

	}

	return true, err
}

func (s *sDept) FindSonByParentId(deptList []*entity.CompanyDept, deptId int64) []*entity.CompanyDept {
	children := make([]*entity.CompanyDept, 0, len(deptList))
	for _, v := range deptList {
		if v.ParentId == deptId {
			children = append(children, v)
			fChildren := s.FindSonByParentId(deptList, v.Id)
			children = append(children, fChildren...)
		}
	}
	return children
}

// GetListTree 部门树形菜单
func (s *sDept) GetListTree(pid int64, list []*model.Dept) (deptTree []*model.DeptTreeRes) {
	deptTree = make([]*model.DeptTreeRes, 0, len(list))
	for _, v := range list {
		if v.ParentID == pid {
			t := &model.DeptTreeRes{
				Dept: v,
			}
			child := s.GetListTree(v.Id, list)
			if len(child) > 0 {
				t.Children = child
			}
			deptTree = append(deptTree, t)
		}
	}
	return
}

func (s *sDept) GetDeptListTree(ctx context.Context, parentId int64) (*model.DeptTreeSelectRes, error) {
	res := new(model.DeptTreeSelectRes)
	var deptList []*model.Dept
	err := s.dao.Dept.Ctx(ctx).Scan(&deptList)
	if err != nil {
		return res, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#DeptName}{#error_Data_Get_Failed}"), s.dao.Dept.Table())
	}

	res.Deps = s.GetListTree(parentId, deptList)
	return res, nil
}
