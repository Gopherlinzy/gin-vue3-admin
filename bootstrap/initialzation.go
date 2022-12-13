package bootstrap

import (
	"github.com/Gopherlinzy/gin-vue3-admin/app/models"
	"github.com/Gopherlinzy/gin-vue3-admin/app/models/api"
	"github.com/Gopherlinzy/gin-vue3-admin/app/models/menu"
	"github.com/Gopherlinzy/gin-vue3-admin/app/models/role"
	"github.com/Gopherlinzy/gin-vue3-admin/app/models/user"
	casbins "github.com/Gopherlinzy/gin-vue3-admin/pkg/casbin"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/database"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/logger"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"sync"
)

var once sync.Once

func IntizationData() {
	once.Do(func() {
		initUserData()
		initApiData()
		initMenuData()
		initRoleData()
		initCasbinData()
		//roleModel := role.GetBy("role_name", "superadmin")
		//apis := roleModel.GetAssociationsApis()
		//fmt.Println(apis)
		//casbins.NewCasbin().DeleteRole("guest")
		//casbins.NewCasbin().AddPolicies([][]string{
		//	{"guest", "1", "2"},
		//	{"guest", "2", "3"},
		//})
		//database.Gohub_DB.Model(&role.Role{BaseModel: models.BaseModel{ID: 7}}).Find(&role.Role{}).Association("Menus").Append(&[]menu.Menu{{BaseModel: models.BaseModel{ID: 1}}, {BaseModel: models.BaseModel{ID: 2}}})
	})
}

func initUserData() {
	// 判断数据表是否存在
	isExist := database.Gohub_DB.Migrator().HasTable(&user.User{})
	if !isExist {
		return
	}

	// 如果数据存在了则不插入
	var count int64
	database.Gohub_DB.Find(&user.User{}).Count(&count)
	if count > 0 {
		return
	}

	users := []user.User{
		{Name: "sly", Email: "123@testing.com", Phone: "00012312312", Password: "123456", City: "Huzhou", Introduction: "帅", Status: true, RoleName: "superAdmin"},
		{Name: "linzy", Email: "123456@testing.com", Phone: "00012345678", Password: "123456", City: "Hangzhou", Introduction: "很帅", Status: true, RoleName: "admin"},
		{Name: "gggg", Email: "gggg@testing.com", Phone: "12345678910", Password: "123456", City: "yingdu", Introduction: "长得不咋地", Status: true, RoleName: "user"},
	}

	database.Gohub_DB.Create(&users)
}

func initMenuData() {
	// 判断数据表是否存在
	isExist := database.Gohub_DB.Migrator().HasTable(&menu.Menu{})
	if !isExist {
		return
	}

	// 如果数据存在了则不插入
	var count int64
	database.Gohub_DB.Find(&menu.Menu{}).Count(&count)
	if count > 0 {
		return
	}

	menus := []menu.Menu{
		{BaseModel: models.BaseModel{ID: 1}, Name: "首页", Permission: "system:index", RouterName: "Index", RouterPath: "/", FatherID: 0, Status: true, VuePath: "@/views/index/Index.vue"},

		{BaseModel: models.BaseModel{ID: 2}, Name: "超级管理员", Permission: "system:superAdmin", RouterName: "superAdmin", RouterPath: "/superAdmin", FatherID: 0, Status: true},
		{BaseModel: models.BaseModel{ID: 3}, Name: "角色管理", Permission: "system:superAdmin:role", RouterName: "roles", RouterPath: "roles", FatherID: 2, Status: true, VuePath: "@/views/superAdmin/Role.vue"},
		{BaseModel: models.BaseModel{ID: 4}, Name: "API管理", Permission: "system:superAdmin:api", RouterName: "apis", RouterPath: "apis", FatherID: 2, Status: true, VuePath: "@/views/superAdmin/Api.vue"},
		{BaseModel: models.BaseModel{ID: 5}, Name: "菜单管理", Permission: "system:superAdmin:menu", RouterName: "menus", RouterPath: "menus", FatherID: 2, Status: true, VuePath: "@/views/superAdmin/Menu.vue"},
		{BaseModel: models.BaseModel{ID: 6}, Name: "用户管理", Permission: "system:superAdmin:user", RouterName: "users", RouterPath: "users", FatherID: 2, Status: true, VuePath: "@/views/superAdmin/User.vue"},
		{BaseModel: models.BaseModel{ID: 7}, Name: "系统设置", Permission: "system:superAdmin:sysSetting", RouterName: "setting", RouterPath: "setting", FatherID: 2, Status: true, VuePath: "@/views/superAdmin/SysSetting.vue"},

		{BaseModel: models.BaseModel{ID: 8}, Name: "订单管理", Permission: "system:order", RouterName: "order", RouterPath: "/order", FatherID: 0, Status: true},
		{BaseModel: models.BaseModel{ID: 9}, Name: "订单查询", Permission: "system:order:orderInfo", RouterName: "orderInfo", RouterPath: "orderInfo", FatherID: 8, Status: true, VuePath: "@/views/orders/OrderInfo.vue"},
		{BaseModel: models.BaseModel{ID: 10}, Name: "订单管理", Permission: "system:order:orderManage", RouterName: "orderManage", RouterPath: "orderManage", FatherID: 8, Status: true, VuePath: "@/views/orders/OrderManage.vue"},

		{BaseModel: models.BaseModel{ID: 11}, Name: "商品管理", Permission: "goods:goods", RouterName: "goods", RouterPath: "/goods", FatherID: 0, Status: true},
		{BaseModel: models.BaseModel{ID: 12}, Name: "商品种类", Permission: "system:goods:goodsCategory", RouterName: "goodsCategory", RouterPath: "goodsCategory", FatherID: 11, Status: true, VuePath: "@/views/goods/GoodsCategory.vue"},
		{BaseModel: models.BaseModel{ID: 13}, Name: "商品信息", Permission: "system:goods:goodsInfo", RouterName: "goodsInfo", RouterPath: "goodsInfo", FatherID: 11, Status: true, VuePath: "@/views/goods/GoodsInfo.vue"},
	}

	database.Gohub_DB.Create(&menus)
}

func initApiData() {
	// 判断数据表是否存在
	isExist := database.Gohub_DB.Migrator().HasTable(&api.Api{})
	if !isExist {
		return
	}

	// 如果数据存在了则不插入
	var count int64
	database.Gohub_DB.Find(&api.Api{}).Count(&count)
	if count > 0 {
		return
	}

	apis := []api.Api{
		// role api权限
		{BaseModel: models.BaseModel{ID: 1}, Path: "/api/v1/roles", Method: "GET", ApiGroup: "roles", Description: "查询角色列表"},
		{BaseModel: models.BaseModel{ID: 2}, Path: "/api/v1/roles", Method: "POST", ApiGroup: "roles", Description: "新增角色数据"},
		{BaseModel: models.BaseModel{ID: 3}, Path: "/api/v1/roles/id", Method: "POST", ApiGroup: "roles", Description: "查询指定角色 id 信息"},
		{BaseModel: models.BaseModel{ID: 4}, Path: "/api/v1/roles/apis", Method: "POST", ApiGroup: "roles", Description: "查询指定角色 id 的所有api权限"},
		{BaseModel: models.BaseModel{ID: 5}, Path: "/api/v1/roles/menus", Method: "POST", ApiGroup: "roles", Description: "查询指定角色 id 所有菜单权限"},
		{BaseModel: models.BaseModel{ID: 6}, Path: "/api/v1/roles", Method: "PUT", ApiGroup: "roles", Description: "更新角色信息"},
		{BaseModel: models.BaseModel{ID: 7}, Path: "/api/v1/roles/menuPermissions", Method: "PUT", ApiGroup: "roles", Description: "更新角色的菜单权限"},
		{BaseModel: models.BaseModel{ID: 8}, Path: "/api/v1/roles/apiPolicy", Method: "PUT", ApiGroup: "roles", Description: "更新角色的api权限"},
		{BaseModel: models.BaseModel{ID: 9}, Path: "/api/v1/roles/status", Method: "PUT", ApiGroup: "roles", Description: "更新角色启用状态"},
		{BaseModel: models.BaseModel{ID: 10}, Path: "/api/v1/roles", Method: "DELETE", ApiGroup: "roles", Description: "删除角色"},

		// user api权限
		{BaseModel: models.BaseModel{ID: 11}, Path: "/api/v1/users", Method: "GET", ApiGroup: "users", Description: "查询用户列表"},
		{BaseModel: models.BaseModel{ID: 12}, Path: "/api/v1/users", Method: "POST", ApiGroup: "users", Description: "新增用户数据"},
		{BaseModel: models.BaseModel{ID: 13}, Path: "/api/v1/users/id", Method: "POST", ApiGroup: "users", Description: "查询指定用户 id 的信息"},
		{BaseModel: models.BaseModel{ID: 14}, Path: "/api/v1/users/reset", Method: "POST", ApiGroup: "users", Description: "重置指定用户的密码为 '123456'"},
		{BaseModel: models.BaseModel{ID: 15}, Path: "/api/v1/users", Method: "PUT", ApiGroup: "users", Description: "更新用户信息"},
		{BaseModel: models.BaseModel{ID: 16}, Path: "/api/v1/users/status", Method: "PUT", ApiGroup: "users", Description: "更新用户启用状态"},
		{BaseModel: models.BaseModel{ID: 17}, Path: "/api/v1/users/role", Method: "PUT", ApiGroup: "users", Description: "更新用户的角色权限"},
		{BaseModel: models.BaseModel{ID: 18}, Path: "/api/v1/users/email", Method: "PUT", ApiGroup: "users", Description: "更新用户的邮箱"},
		{BaseModel: models.BaseModel{ID: 19}, Path: "/api/v1/users/phone", Method: "PUT", ApiGroup: "users", Description: "更新用户的手机号"},
		{BaseModel: models.BaseModel{ID: 20}, Path: "/api/v1/users/password", Method: "PUT", ApiGroup: "users", Description: "更新用户的密码"},
		{BaseModel: models.BaseModel{ID: 21}, Path: "/api/v1/users/avatar", Method: "PUT", ApiGroup: "users", Description: "更新用户的头像"},
		{BaseModel: models.BaseModel{ID: 22}, Path: "/api/v1/users", Method: "DELETE", ApiGroup: "users", Description: "删除用户信息"},

		// menu api权限
		{BaseModel: models.BaseModel{ID: 23}, Path: "/api/v1/menus", Method: "GET", ApiGroup: "menus", Description: "查询菜单列表"},
		{BaseModel: models.BaseModel{ID: 24}, Path: "/api/v1/menus/pag", Method: "GET", ApiGroup: "menus", Description: "查询菜单分页列表"},
		{BaseModel: models.BaseModel{ID: 25}, Path: "/api/v1/menus", Method: "POST", ApiGroup: "menus", Description: "新增菜单信息"},
		{BaseModel: models.BaseModel{ID: 26}, Path: "/api/v1/menus/id", Method: "POST", ApiGroup: "menus", Description: "查询指定菜单 id 数据"},
		{BaseModel: models.BaseModel{ID: 27}, Path: "/api/v1/menus", Method: "PUT", ApiGroup: "menus", Description: "更新菜单信息"},
		{BaseModel: models.BaseModel{ID: 28}, Path: "/api/v1/menus", Method: "DELETE", ApiGroup: "menus", Description: "删除菜单信息"},

		// apis api权限
		{BaseModel: models.BaseModel{ID: 29}, Path: "/api/v1/apis", Method: "GET", ApiGroup: "apis", Description: "查询api列表"},
		{BaseModel: models.BaseModel{ID: 30}, Path: "/api/v1/apis/pag", Method: "GET", ApiGroup: "apis", Description: "查询api分页列表"},
		{BaseModel: models.BaseModel{ID: 31}, Path: "/api/v1/apis", Method: "POST", ApiGroup: "apis", Description: "新增api信息"},
		{BaseModel: models.BaseModel{ID: 32}, Path: "/api/v1/apis/id", Method: "POST", ApiGroup: "apis", Description: "查询指定api id 数据"},
		{BaseModel: models.BaseModel{ID: 33}, Path: "/api/v1/apis", Method: "PUT", ApiGroup: "apis", Description: "更新api信息"},
		{BaseModel: models.BaseModel{ID: 34}, Path: "/api/v1/apis", Method: "DELETE", ApiGroup: "apis", Description: "删除api信息"},
	}

	database.Gohub_DB.Create(&apis)
}

func initRoleData() {
	// 判断数据表是否存在
	isExist := database.Gohub_DB.Migrator().HasTable(&role.Role{})
	if !isExist {
		return
	}

	// 如果数据存在了则不插入
	var count int64
	database.Gohub_DB.Find(&role.Role{}).Count(&count)
	if count > 0 {
		return
	}

	roles := []role.Role{
		{BaseModel: models.BaseModel{ID: 1}, RoleName: "superAdmin", Des: "超级管理员:拥有所有权限", Status: true, Menus: menu.ExceptID(), Apis: api.ExceptID([]uint64{})},
		{BaseModel: models.BaseModel{ID: 2}, RoleName: "admin", Des: "管理员", Status: true, Menus: menu.ExceptID(2, 3), Apis: api.ExceptID([]uint64{4, 5, 6, 7, 8, 9, 10})},
		{BaseModel: models.BaseModel{ID: 3}, RoleName: "user", Des: "普通用户", Status: true, Menus: menu.ExceptID(2, 3, 4, 5, 6, 7), Apis: api.AcceptID([]uint64{11, 18, 19, 20, 21, 23, 28}, 32)},
		{BaseModel: models.BaseModel{ID: 4}, RoleName: "guest", Des: "游客模式", Status: false, Menus: menu.ExceptID(2, 3, 4, 5, 6, 7, 8, 10, 11, 13, 14, 15), Apis: api.AcceptID([]uint64{11}, 32)},
	}

	database.Gohub_DB.Create(&roles)
}

func initCasbinData() {
	cs := casbins.NewCasbin()
	// 如果数据存在了则不插入
	var count int64
	database.Gohub_DB.Find(&gormadapter.CasbinRule{}).Count(&count)
	if count > 0 {
		return
	}

	policies := []gormadapter.CasbinRule{
		// 用户绑定角色
		{Ptype: "g", V0: "sly", V1: "superAdmin"},
		{Ptype: "g", V0: "linzy", V1: "admin"},
		{Ptype: "g", V0: "gggg", V1: "user"},

		// 超级管理员 其实可以不写 我只是记录我有多少api
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/roles", V2: "GET"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/roles/id", V2: "POST"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/roles/policies", V2: "POST"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/roles/menus", V2: "POST"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/roles/apis", V2: "POST"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/roles", V2: "POST"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/roles", V2: "PUT"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/roles/menuPermissions", V2: "PUT"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/roles/apiPolicy", V2: "PUT"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/roles/status", V2: "PUT"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/roles", V2: "DELETE"},

		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/users", V2: "GET"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/users", V2: "POST"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/users/id", V2: "POST"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/users/reset", V2: "POST"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/users", V2: "PUT"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/users/status", V2: "PUT"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/users/role", V2: "PUT"},
		//{Ptype: "p", V0: "superAdmin", V1: "/api/v1/users/profile", V2: "PUT"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/users/email", V2: "PUT"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/users/phone", V2: "PUT"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/users/password", V2: "PUT"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/users/avatar", V2: "PUT"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/users", V2: "DELETE"},

		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/menus", V2: "GET"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/menus/pag", V2: "GET"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/menus", V2: "POST"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/menus/id", V2: "POST"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/menus", V2: "PUT"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/menus", V2: "DELETE"},

		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/apis", V2: "GET"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/apis/pag", V2: "GET"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/apis", V2: "POST"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/apis/id", V2: "POST"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/apis", V2: "PUT"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/apis", V2: "DELETE"},

		// 管理员
		{Ptype: "p", V0: "admin", V1: "/api/v1/roles/id", V2: "GET"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/roles", V2: "GET"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/roles", V2: "POST"},

		{Ptype: "p", V0: "admin", V1: "/api/v1/users", V2: "GET"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/users", V2: "POST"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/users/id", V2: "POST"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/users/reset", V2: "POST"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/users", V2: "PUT"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/users/status", V2: "PUT"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/users/role", V2: "PUT"},
		//{Ptype: "p", V0: "superAdmin", V1: "/api/v1/users/profile", V2: "PUT"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/users/email", V2: "PUT"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/users/phone", V2: "PUT"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/users/password", V2: "PUT"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/users/avatar", V2: "PUT"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/users", V2: "DELETE"},

		{Ptype: "p", V0: "admin", V1: "/api/v1/menus", V2: "GET"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/menus/pag", V2: "GET"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/menus", V2: "POST"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/menus/id", V2: "POST"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/menus", V2: "PUT"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/menus", V2: "DELETE"},

		{Ptype: "p", V0: "admin", V1: "/api/v1/apis", V2: "GET"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/apis/pag", V2: "GET"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/apis", V2: "POST"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/apis/id", V2: "POST"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/apis", V2: "PUT"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/apis", V2: "DELETE"},

		// 普通用户
		{Ptype: "p", V0: "user", V1: "/api/v1/users", V2: "GET"},
		//{Ptype: "p", V0: "user", V1: "/api/v1/users/profile", V2: "PUT"},
		{Ptype: "p", V0: "user", V1: "/api/v1/users/email", V2: "PUT"},
		{Ptype: "p", V0: "user", V1: "/api/v1/users/phone", V2: "PUT"},
		{Ptype: "p", V0: "user", V1: "/api/v1/users/password", V2: "PUT"},
		{Ptype: "p", V0: "user", V1: "/api/v1/users/avatar", V2: "PUT"},

		{Ptype: "p", V0: "user", V1: "/api/v1/menus", V2: "GET"},
		{Ptype: "p", V0: "user", V1: "/api/v1/menus/pag", V2: "GET"},

		{Ptype: "p", V0: "user", V1: "/api/v1/apis", V2: "GET"},

		//{Ptype: "p", V0: "user", V1: "/api/v1/categories", V2: "GET"},
		//{Ptype: "p", V0: "user", V1: "/api/v1/categories", V2: "POST"},
		//{Ptype: "p", V0: "user", V1: "/api/v1/categories", V2: "PUT"},
		//{Ptype: "p", V0: "user", V1: "/api/v1/categories", V2: "DELETE"},

		// 游客
		{Ptype: "p", V0: "guest", V1: "/api/v1/users", V2: "GET"},
		{Ptype: "p", V0: "guest", V1: "/api/v1/menus/pag", V2: "GET"},
		{Ptype: "p", V0: "guest", V1: "/api/v1/apis/pag", V2: "GET"},
		//{Ptype: "p", V0: "guest", V1: "/api/v1/categories", V2: "GET"},
	}

	database.Gohub_DB.Model(&gormadapter.CasbinRule{}).Create(policies)

	// 清除内存中的读取的缓存
	err := cs.Enforcer.InvalidateCache()
	if err != nil {
		logger.LogIf(err)
		return
	}
	// 重新缓存
	_ = cs.Enforcer.LoadPolicy()
}
