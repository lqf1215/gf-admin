package co_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/gogf/gf/v2/database/gdb"
)

type TableName struct {
	Dept string `p:"dept" dc:"部门表名"`
	Post string `p:"post" dc:"岗位表名"`
}

type Identifier struct {
	Dept string `p:"dept" dc:"部门标识符"`
	Post string `p:"post" dc:"岗位标识符"`
}

type Config struct {
	DB                             gdb.DB            `p:"-" dc:"数据库连接"`
	AllowEmptyNo                   bool              `p:"allowEmptyNo" dc:"允许员工工号为空" default:"false"`
	IsCreateDefaultEmployeeAndRole bool              `p:"isCreateDefaultEmployeeAndRole" dc:"是否创建默认员工和角色"`
	HardDeleteWaitAt               int64             `p:"hardDeleteWaitAt" dc:"硬删除等待时限,单位/小时" default:"12"`
	KeyIndex                       string            `p:"keyIndex" dc:"配置索引"`
	RoutePrefix                    string            `p:"routePrefix" dc:"路由前缀"`
	StoragePath                    string            `p:"storagePath" dc:"资源存储路径"`
	UserType                       sys_enum.UserType `p:"userType" dc:"用户类型"`
	Identifier                     Identifier        `p:"identifier" dc:"标识符"`
	TableName                      TableName         `p:"tableName" dc:"模块表名"`
}
