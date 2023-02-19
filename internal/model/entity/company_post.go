// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CompanyPost is the golang structure for table company_post.
type CompanyPost struct {
	Id        int64       `json:"id"        description:"ID"`
	Code      string      `json:"code"      description:"岗位编码"`
	Name      string      `json:"name"      description:"岗位名称"`
	State     int         `json:"state"     description:"状态（0正常 1停用）"`
	Sort      int         `json:"sort"      description:"显示顺序"`
	Remark    string      `json:"remark"    description:"备注"`
	CreatedBy string      `json:"createdBy" description:"创建人"`
	UpdatedBy string      `json:"updatedBy" description:"修改人"`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	DeletedAt *gtime.Time `json:"deletedAt" description:""`
}
