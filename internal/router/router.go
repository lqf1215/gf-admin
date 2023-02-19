package router

import (
	"gf-admin/internal/co_interface"
	"gf-admin/internal/controller"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
)

func ModulesGroup(modules co_interface.IModules, group *ghttp.RouterGroup) *ghttp.RouterGroup {
	DeptGroup(modules, group)
	PostGroup(modules, group)
	return group
}

func DeptGroup(modules co_interface.IModules, group *ghttp.RouterGroup) *ghttp.RouterGroup {
	routePrefix := modules.GetConfig().RoutePrefix + "/" + gstr.LcFirst(modules.GetConfig().Identifier.Dept)
	controller := controller.Dept(modules)
	group.POST(routePrefix+"/getDeptById", controller.GetDeptById)
	group.POST(routePrefix+"/hasDeptByName", controller.HasDeptByName)
	group.POST(routePrefix+"/queryDeptList", controller.QueryDeptList)
	group.POST(routePrefix+"/getDeptListTree", controller.GetDeptListTree)
	group.POST(routePrefix+"/createDept", controller.CreateDept)
	group.POST(routePrefix+"/updateDept", controller.UpdateDept)
	group.POST(routePrefix+"/deleteDept", controller.DeleteDept)
	return group
}
func PostGroup(modules co_interface.IModules, group *ghttp.RouterGroup) *ghttp.RouterGroup {
	routePrefix := modules.GetConfig().RoutePrefix + "/" + gstr.LcFirst(modules.GetConfig().Identifier.Post)
	controller := controller.Post(modules)
	group.POST(routePrefix+"/getPostById", controller.GetPostById)
	group.POST(routePrefix+"/hasPostByName", controller.HasPostByName)
	group.POST(routePrefix+"/queryPostList", controller.QueryPostList)
	group.POST(routePrefix+"/createPost", controller.CreatePost)
	group.POST(routePrefix+"/updatePost", controller.UpdatePost)
	group.POST(routePrefix+"/deletePost", controller.DeletePost)
	return group
}
