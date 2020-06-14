package sysController

import (
    "gin-test/common"
    "gin-test/utils/code"
    "github.com/gin-gonic/gin"
    "net/http"
)

var Routers gin.RoutesInfo

// @Summary 菜单列表
// @Description get router list
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Tags SYS
// @Success 200 {object} common.Response
// @Router /sys/menu_list [get]
func GetMenuList(c *gin.Context) {
    testMenuList := `{"code":200,"data":[{"menuId":2,"menuName":"Upms","title":"系统管理","icon":"example","path":"/upms","paths":"/0/2","menuType":"M","action":"无","permission":"","parentId":0,"noCache":true,"breadcrumb":"","component":"Layout","sort":1,"visible":"0","createBy":"1","updateBy":"1","isFrame":"0","dataScope":"","params":"","RoleId":0,"children":[{"menuId":3,"menuName":"Sysuser","title":"用户管理","icon":"user","path":"sysuser","paths":"","menuType":"C","action":"无","permission":"system:sysuser:list","parentId":2,"noCache":false,"breadcrumb":"","component":"/sysuser/index","sort":1,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[{"menuId":46,"menuName":"","title":"删除用户","icon":"","path":"/api/v1/sysuser/","paths":"","menuType":"F","action":"DELETE","permission":"system:sysuser:remove","parentId":3,"noCache":false,"breadcrumb":"","component":"","sort":0,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":45,"menuName":"","title":"修改用户","icon":"","path":"/api/v1/sysuser/","paths":"","menuType":"F","action":"PUT","permission":"system:sysuser:edit","parentId":3,"noCache":false,"breadcrumb":"","component":"","sort":0,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":44,"menuName":"","title":"查询用户","icon":"","path":"/api/v1/sysuser","paths":"","menuType":"F","action":"GET","permission":"system:sysuser:query","parentId":3,"noCache":false,"breadcrumb":"","component":"","sort":0,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":43,"menuName":"","title":"新增用户","icon":"","path":"/api/v1/sysuser","paths":"","menuType":"F","action":"POST","permission":"system:sysuser:add","parentId":3,"noCache":false,"breadcrumb":"","component":"","sort":0,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null}],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":52,"menuName":"Role","title":"角色管理","icon":"peoples","path":"role","paths":"","menuType":"C","action":"无","permission":"system:sysrole:list","parentId":2,"noCache":true,"breadcrumb":"","component":"/role/index","sort":2,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[{"menuId":227,"menuName":"","title":"删除角色","icon":"","path":"","paths":"","menuType":"F","action":"","permission":"system:sysrole:remove","parentId":52,"noCache":false,"breadcrumb":"","component":"","sort":1,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":226,"menuName":"","title":"修改角色","icon":"","path":"","paths":"","menuType":"F","action":"","permission":"system:sysrole:edit","parentId":52,"noCache":false,"breadcrumb":"","component":"","sort":1,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":225,"menuName":"","title":"查询角色","icon":"","path":"","paths":"","menuType":"F","action":"","permission":"system:sysrole:query","parentId":52,"noCache":false,"breadcrumb":"","component":"","sort":1,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":224,"menuName":"","title":"新增角色","icon":"","path":"","paths":"","menuType":"F","action":"","permission":"system:sysrole:add","parentId":52,"noCache":false,"breadcrumb":"","component":"","sort":1,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null}],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":51,"menuName":"Menu","title":"菜单管理","icon":"tree-table","path":"menu","paths":"","menuType":"C","action":"无","permission":"system:sysmenu:list","parentId":2,"noCache":true,"breadcrumb":"","component":"/menu/index","sort":3,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[{"menuId":223,"menuName":"","title":"删除菜单","icon":"","path":"","paths":"","menuType":"F","action":"","permission":"system:sysmenu:remove","parentId":51,"noCache":false,"breadcrumb":"","component":"","sort":1,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":222,"menuName":"","title":"查询菜单","icon":"search","path":"","paths":"","menuType":"F","action":"","permission":"system:sysmenu:query","parentId":51,"noCache":false,"breadcrumb":"","component":"","sort":1,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":221,"menuName":"","title":"修改菜单","icon":"edit","path":"","paths":"","menuType":"F","action":"","permission":"system:sysmenu:edit","parentId":51,"noCache":false,"breadcrumb":"","component":"","sort":1,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":220,"menuName":"","title":"新增菜单","icon":"","path":"","paths":"","menuType":"F","action":"","permission":"system:sysmenu:add","parentId":51,"noCache":false,"breadcrumb":"","component":"","sort":1,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null}],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":58,"menuName":"Dict","title":"字典管理","icon":"education","path":"dict","paths":"","menuType":"C","action":"无","permission":"system:sysdicttype:list","parentId":2,"noCache":false,"breadcrumb":"","component":"/dict/index","sort":4,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[{"menuId":239,"menuName":"","title":"删除类型","icon":"","path":"","paths":"","menuType":"F","action":"","permission":"system:sysdicttype:remove","parentId":58,"noCache":false,"breadcrumb":"","component":"","sort":0,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":238,"menuName":"","title":"修改类型","icon":"","path":"","paths":"","menuType":"F","action":"","permission":"system:sysdicttype:edit","parentId":58,"noCache":false,"breadcrumb":"","component":"","sort":0,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":237,"menuName":"","title":"新增类型","icon":"","path":"","paths":"","menuType":"F","action":"","permission":"system:sysdicttype:add","parentId":58,"noCache":false,"breadcrumb":"","component":"","sort":0,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":236,"menuName":"","title":"字典查询","icon":"","path":"","paths":"","menuType":"F","action":"","permission":"system:sysdicttype:query","parentId":58,"noCache":false,"breadcrumb":"","component":"","sort":0,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null}],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":59,"menuName":"DictData","title":"字典数据","icon":"education","path":"dict/data/:dictId","paths":"","menuType":"C","action":"无","permission":"system:sysdictdata:list","parentId":2,"noCache":false,"breadcrumb":"","component":"/dict/data","sort":5,"visible":"1","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[{"menuId":242,"menuName":"","title":"修改数据","icon":"","path":"","paths":"","menuType":"F","action":"","permission":"system:sysdictdata:edit","parentId":59,"noCache":false,"breadcrumb":"","component":"","sort":0,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":241,"menuName":"","title":"新增数据","icon":"","path":"","paths":"","menuType":"F","action":"","permission":"system:sysdictdata:add","parentId":59,"noCache":false,"breadcrumb":"","component":"","sort":0,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":240,"menuName":"","title":"查询数据","icon":"","path":"","paths":"","menuType":"F","action":"","permission":"system:sysdictdata:query","parentId":59,"noCache":false,"breadcrumb":"","component":"","sort":0,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":243,"menuName":"","title":"删除数据","icon":"","path":"","paths":"","menuType":"F","action":"","permission":"system:sysdictdata:remove","parentId":59,"noCache":false,"breadcrumb":"","component":"","sort":0,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null}],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":57,"menuName":"post","title":"岗位管理","icon":"pass","path":"post","paths":"","menuType":"C","action":"无","permission":"system:syspost:list","parentId":2,"noCache":false,"breadcrumb":"","component":"/post/index","sort":6,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[{"menuId":235,"menuName":"","title":"删除岗位","icon":"","path":"","paths":"","menuType":"F","action":"","permission":"system:syspost:remove","parentId":57,"noCache":false,"breadcrumb":"","component":"","sort":0,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":234,"menuName":"","title":"修改岗位","icon":"","path":"","paths":"","menuType":"F","action":"","permission":"system:syspost:edit","parentId":57,"noCache":false,"breadcrumb":"","component":"","sort":0,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":233,"menuName":"","title":"新增岗位","icon":"","path":"","paths":"","menuType":"F","action":"","permission":"system:syspost:add","parentId":57,"noCache":false,"breadcrumb":"","component":"","sort":0,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":232,"menuName":"","title":"查询岗位","icon":"","path":"","paths":"","menuType":"F","action":"","permission":"system:syspost:query","parentId":57,"noCache":false,"breadcrumb":"","component":"","sort":0,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null}],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":56,"menuName":"Dept","title":"部门管理","icon":"tree","path":"dept","paths":"","menuType":"C","action":"无","permission":"system:sysdept:list","parentId":2,"noCache":false,"breadcrumb":"","component":"/dept/index","sort":7,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[{"menuId":231,"menuName":"","title":"删除部门","icon":"","path":"","paths":"","menuType":"F","action":"","permission":"system:sysdept:remove","parentId":56,"noCache":false,"breadcrumb":"","component":"","sort":0,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":230,"menuName":"","title":"修改部门","icon":"","path":"","paths":"","menuType":"F","action":"","permission":"system:sysdept:edit","parentId":56,"noCache":false,"breadcrumb":"","component":"","sort":0,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":228,"menuName":"","title":"查询部门","icon":"","path":"","paths":"","menuType":"F","action":"","permission":"system:sysdept:query","parentId":56,"noCache":false,"breadcrumb":"","component":"","sort":1,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":229,"menuName":"","title":"新增部门","icon":"","path":"","paths":"","menuType":"F","action":"","permission":"system:sysdept:add","parentId":56,"noCache":false,"breadcrumb":"","component":"","sort":1,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null}],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":62,"menuName":"Config","title":"参数设置","icon":"list","path":"/config","paths":"","menuType":"C","action":"无","permission":"system:sysconfig:list","parentId":2,"noCache":false,"breadcrumb":"","component":"/config/index","sort":9,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[{"menuId":247,"menuName":"","title":"删除参数","icon":"","path":"","paths":"","menuType":"F","action":"","permission":"system:sysconfig:remove","parentId":62,"noCache":false,"breadcrumb":"","component":"","sort":0,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":246,"menuName":"","title":"修改参数","icon":"","path":"","paths":"","menuType":"F","action":"","permission":"system:sysconfig:edit","parentId":62,"noCache":false,"breadcrumb":"","component":"","sort":0,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":245,"menuName":"","title":"新增参数","icon":"","path":"","paths":"","menuType":"F","action":"","permission":"system:sysconfig:add","parentId":62,"noCache":false,"breadcrumb":"","component":"","sort":0,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":244,"menuName":"","title":"查询参数","icon":"","path":"","paths":"","menuType":"F","action":"","permission":"system:sysconfig:query","parentId":62,"noCache":false,"breadcrumb":"","component":"","sort":0,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null}],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":211,"menuName":"Log","title":"日志管理","icon":"log","path":"/log","paths":"","menuType":"M","action":"","permission":"","parentId":2,"noCache":false,"breadcrumb":"","component":"/log/index","sort":10,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[{"menuId":216,"menuName":"OperLog","title":"操作日志","icon":"skill","path":"/operlog","paths":"","menuType":"C","action":"","permission":"system:sysoperlog:list","parentId":211,"noCache":false,"breadcrumb":"","component":"/operlog/index","sort":1,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[{"menuId":251,"menuName":"","title":"删除操作日志","icon":"","path":"","paths":"","menuType":"F","action":"","permission":"system:sysoperlog:remove","parentId":216,"noCache":false,"breadcrumb":"","component":"","sort":0,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":250,"menuName":"","title":"查询操作日志","icon":"","path":"","paths":"","menuType":"F","action":"","permission":"system:sysoperlog:query","parentId":216,"noCache":false,"breadcrumb":"","component":"","sort":0,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null}],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":212,"menuName":"LoginLog","title":"登录日志","icon":"logininfor","path":"/loginlog","paths":"","menuType":"C","action":"","permission":"system:sysloginlog:list","parentId":211,"noCache":false,"breadcrumb":"","component":"/loginlog/index","sort":1,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[{"menuId":249,"menuName":"","title":"删除登录日志","icon":"","path":"","paths":"","menuType":"F","action":"","permission":"system:sysloginlog:remove","parentId":212,"noCache":false,"breadcrumb":"","component":"","sort":0,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":248,"menuName":"","title":"查询登录日志","icon":"","path":"","paths":"","menuType":"F","action":"","permission":"system:sysloginlog:query","parentId":212,"noCache":false,"breadcrumb":"","component":"","sort":0,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null}],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null}],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null}],"is_select":false,"createdAt":"2019-11-26T21:22:09+08:00","updatedAt":"2020-04-08T15:07:57+08:00","deletedAt":null},{"menuId":60,"menuName":"Tools","title":"系统工具","icon":"component","path":"/tools","paths":"/0/60","menuType":"M","action":"无","permission":"","parentId":0,"noCache":false,"breadcrumb":"","component":"Layout","sort":3,"visible":"0","createBy":"1","updateBy":"1","isFrame":"0","dataScope":"","params":"","RoleId":0,"children":[{"menuId":262,"menuName":"EditTable","title":"数据表修改","icon":"","path":"editTable","paths":"","menuType":"C","action":"","permission":"","parentId":60,"noCache":false,"breadcrumb":"","component":"/tools/gen/editTable","sort":0,"visible":"1","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":261,"menuName":"Gen","title":"代码生成","icon":"code","path":"gen","paths":"","menuType":"C","action":"","permission":"","parentId":60,"noCache":false,"breadcrumb":"","component":"/tools/gen/index","sort":0,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":61,"menuName":"Swagger","title":"系统接口","icon":"guide","path":"swagger","paths":"","menuType":"C","action":"无","permission":"","parentId":60,"noCache":false,"breadcrumb":"","component":"/tools/swagger/index","sort":1,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null},{"menuId":260,"menuName":"Build","title":"表单构建","icon":"build","path":"build","paths":"","menuType":"C","action":"","permission":"tools:build:list","parentId":60,"noCache":false,"breadcrumb":"","component":"/tools/build/index","sort":1,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null}],"is_select":false,"createdAt":"2020-02-28T23:36:21+08:00","updatedAt":"2020-04-08T14:31:35+08:00","deletedAt":null},{"menuId":264,"menuName":"Monitor","title":"系统监控","icon":"monitor","path":"/monitor","paths":"/0/264","menuType":"M","action":"","permission":"","parentId":0,"noCache":false,"breadcrumb":"","component":"Layout","sort":4,"visible":"0","createBy":"1","updateBy":"","isFrame":"1","dataScope":"","params":"","RoleId":0,"children":[{"menuId":265,"menuName":"Server","title":"服务监控","icon":"druid","path":"/server","paths":"","menuType":"C","action":"","permission":"monitor:server:list","parentId":264,"noCache":false,"breadcrumb":"","component":"/monitor/server/index","sort":0,"visible":"0","createBy":"","updateBy":"","isFrame":"","dataScope":"","params":"","RoleId":0,"children":[],"is_select":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null}],"is_select":false,"createdAt":"2020-04-14T00:26:15+08:00","updatedAt":"2020-04-14T00:26:15+08:00","deletedAt":null}],"msg":""}`
    c.Writer.Header().Set("Content-Type", "application/json")
    c.Writer.WriteString(testMenuList)
}

// @Summary 后端存在路由列表
// @Description get router list
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Tags SYS
// @Success 200 {object} common.Response
// @Router /sys/router [get]
func GetRouterList(c *gin.Context) {
    appG := common.Gin{C: c}

    type Router struct {
        Path   string `json:"path"`
        Method string `json:"method"`
    }

    data := make([]Router, 0)

    routers := Routers

    for index := range routers{
        var router Router
        router.Method = Routers[index].Method
        router.Path = Routers[index].Path
        data = append(data, router)
    }

    appG.Response(http.StatusOK, code.SUCCESS, "获取存在路由列表成功", data)
}