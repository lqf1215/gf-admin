package post

import "gf-admin/internal/co_interface"

type post struct {
	PermissionType func(modules co_interface.IModules) *permissionType[co_interface.IModules]
	State          state
}

var Post = post{
	PermissionType: PermissionType,
	State:          State,
}
