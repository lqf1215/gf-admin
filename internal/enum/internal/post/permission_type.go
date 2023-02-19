package post

import (
	"gf-admin/internal/co_interface"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/utility/kmap"
	"github.com/SupenBysz/gf-admin-community/utility/permission"
	"github.com/gogf/gf/v2/util/gconv"
)

type PermissionEnum = *permission.SysPermissionTree

type permissionType[T co_interface.IModules] struct {
	modules    T
	enumMap    *kmap.HashMap[string, PermissionEnum]
	ViewDetail PermissionEnum
	List       PermissionEnum
	ListTree   PermissionEnum
	Create     PermissionEnum
	Update     PermissionEnum
	Delete     PermissionEnum
	SetState   PermissionEnum
}

var (
	permissionTypeMap = kmap.New[string, *permissionType[co_interface.IModules]]()
	PermissionType    = func(modules co_interface.IModules) *permissionType[co_interface.IModules] {
		result := permissionTypeMap.GetOrSet(modules.GetConfig().KeyIndex, &permissionType[co_interface.IModules]{
			modules:    modules,
			enumMap:    kmap.New[string, PermissionEnum](),
			ViewDetail: permission.NewInIdentifier("ViewDetail", "详情", "查看岗位详情"),
			List:       permission.NewInIdentifier("List", "列表", "查看岗位列表"),
			Create:     permission.NewInIdentifier("Create", "新增", "新增岗位信息"),
			Update:     permission.NewInIdentifier("Update", "更新", "更新岗位信息"),
			Delete:     permission.NewInIdentifier("Delete", "删除", "删除岗位信息"),
			SetState:   permission.NewInIdentifier("SetState", "设置状态", "设置岗位状态"),
		})

		for k, v := range gconv.Map(result) {
			result.enumMap.Set(k, v.(PermissionEnum))
		}
		return result
	}
)

// ByCode 通过枚举值取枚举类型
func (e *permissionType[T]) ByCode(identifier string) *sys_entity.SysPermission {
	v, has := e.enumMap.Search(identifier)
	if v != nil && has {
		return v.SysPermission
	}
	return nil
}
