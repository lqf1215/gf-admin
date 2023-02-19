package i_controller

import "gf-admin/internal/co_interface"

type iModule interface {
	GetModules() co_interface.IModules
}
