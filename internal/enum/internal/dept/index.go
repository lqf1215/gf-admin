package dept

import "gf-admin/internal/co_interface"

type dept struct {
	PermissionType func(modules co_interface.IModules) *permissionType[co_interface.IModules]
	State          state
}

var Dept = dept{
	PermissionType: PermissionType,
	State:          State,
}
