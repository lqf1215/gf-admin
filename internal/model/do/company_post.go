// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CompanyPost is the golang structure of table co_company_post for DAO operations like Where/Data.
type CompanyPost struct {
	g.Meta    `orm:"table:co_company_post, do:true"`
	Id        interface{} // ID
	Code      interface{} // 岗位编码
	Name      interface{} // 岗位名称
	State     interface{} // 状态（0正常 1停用）
	Sort      interface{} // 显示顺序
	Remark    interface{} // 备注
	CreatedBy interface{} // 创建人
	UpdatedBy interface{} // 修改人
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
