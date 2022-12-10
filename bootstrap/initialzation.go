package bootstrap

import (
	"github.com/Gopherlinzy/gin-vue3-admin/app/models"
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
		initMenuData()
		initRoleData()
		initCasbinData()
		//casbins.NewCasbin().AddPolicy("superAdmin", "/api/v1/roles/status", "PUT")
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
		{BaseModel: models.BaseModel{ID: 1}, Name: "首页", Permissions: "system:index", RouterName: "Index", RouterPath: "/", FatherID: 0, Status: true, VuePath: "@/views/index/Index.vue"},

		{BaseModel: models.BaseModel{ID: 2}, Name: "超级管理员", Permissions: "system:superAdmin", RouterName: "superAdmin", RouterPath: "/superAdmin", FatherID: 0, Status: true},
		{BaseModel: models.BaseModel{ID: 3}, Name: "角色管理", Permissions: "system:superAdmin:role", RouterName: "roles", RouterPath: "roles", FatherID: 2, Status: true, VuePath: "@/views/superAdmin/Role.vue"},
		{BaseModel: models.BaseModel{ID: 4}, Name: "API管理", Permissions: "system:superAdmin:api", RouterName: "apis", RouterPath: "apis", FatherID: 2, Status: true, VuePath: "@/views/superAdmin/Api.vue"},
		{BaseModel: models.BaseModel{ID: 5}, Name: "菜单管理", Permissions: "system:superAdmin:menu", RouterName: "menus", RouterPath: "menus", FatherID: 2, Status: true, VuePath: "@/views/superAdmin/Menu.vue"},
		{BaseModel: models.BaseModel{ID: 6}, Name: "用户管理", Permissions: "system:superAdmin:user", RouterName: "users", RouterPath: "users", FatherID: 2, Status: true, VuePath: "@/views/superAdmin/User.vue"},
		{BaseModel: models.BaseModel{ID: 7}, Name: "系统设置", Permissions: "system:superAdmin:sysSetting", RouterName: "setting", RouterPath: "setting", FatherID: 2, Status: true, VuePath: "@/views/superAdmin/SysSetting.vue"},

		{BaseModel: models.BaseModel{ID: 8}, Name: "订单管理", Permissions: "system:order", RouterName: "order", RouterPath: "/order", FatherID: 0, Status: true},
		{BaseModel: models.BaseModel{ID: 9}, Name: "订单查询", Permissions: "system:order:orderInfo", RouterName: "orderInfo", RouterPath: "orderInfo", FatherID: 8, Status: true, VuePath: "@/views/orders/OrderInfo.vue"},
		{BaseModel: models.BaseModel{ID: 10}, Name: "订单管理", Permissions: "system:order:orderManage", RouterName: "orderManage", RouterPath: "orderManage", FatherID: 8, Status: true, VuePath: "@/views/orders/OrderManage.vue"},

		{BaseModel: models.BaseModel{ID: 11}, Name: "商品管理", Permissions: "goods:goods", RouterName: "goods", RouterPath: "/goods", FatherID: 0, Status: true},
		{BaseModel: models.BaseModel{ID: 12}, Name: "商品种类", Permissions: "system:goods:goodsCategory", RouterName: "goodsCategory", RouterPath: "goodsCategory", FatherID: 11, Status: true, VuePath: "@/views/goods/GoodsCategory.vue"},
		{BaseModel: models.BaseModel{ID: 13}, Name: "商品信息", Permissions: "system:goods:goodsInfo", RouterName: "goodsInfo", RouterPath: "goodsInfo", FatherID: 11, Status: true, VuePath: "@/views/goods/GoodsInfo.vue"},
		{BaseModel: models.BaseModel{ID: 14}, Name: "添加商品", Permissions: "system:goods:goodsInfo:add", RouterName: "orderManage", RouterPath: "orderManage", FatherID: 13, Status: true},
		{BaseModel: models.BaseModel{ID: 15}, Name: "修改商品", Permissions: "system:goods:goodsInfo:update", RouterName: "orderManage", RouterPath: "orderManage", FatherID: 13, Status: true},
	}

	database.Gohub_DB.Create(&menus)
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
		{BaseModel: models.BaseModel{ID: 1}, RoleName: "superAdmin", Des: "超级管理员:拥有所有权限", Status: true, Menus: menu.ExceptID()},
		{BaseModel: models.BaseModel{ID: 2}, RoleName: "admin", Des: "管理员", Status: true, Menus: menu.ExceptID(2, 3)},
		{BaseModel: models.BaseModel{ID: 3}, RoleName: "user", Des: "普通用户", Status: true, Menus: menu.ExceptID(2, 3, 4, 5, 6, 7)},
		{BaseModel: models.BaseModel{ID: 4}, RoleName: "guest", Des: "游客模式", Status: false, Menus: menu.ExceptID(2, 3, 4, 5, 6, 7, 8, 10, 11, 13, 14, 15)},
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
		// 超级管理员
		{Ptype: "g", V0: "sly", V1: "superAdmin"},
		{Ptype: "g", V0: "linzy", V1: "admin"},
		{Ptype: "g", V0: "gggg", V1: "user"},

		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/users", V2: "GET"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/users", V2: "POST"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/users/id", V2: "POST"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/users/reset", V2: "POST"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/users", V2: "PUT"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/users/status", V2: "PUT"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/users/role", V2: "PUT"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/users/profile", V2: "PUT"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/users/email", V2: "PUT"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/users/phone", V2: "PUT"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/users/password", V2: "PUT"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/users/avatar", V2: "PUT"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/users", V2: "DELETE"},

		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/categories", V2: "GET"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/categories", V2: "POST"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/categories", V2: "PUT"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/categories", V2: "DELETE"},

		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/roles/id", V2: "POST"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/roles", V2: "GET"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/roles", V2: "POST"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/roles", V2: "PUT"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/roles", V2: "DELETE"},

		// 管理员
		{Ptype: "p", V0: "admin", V1: "/api/v1/users", V2: "GET"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/users", V2: "POST"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/users", V2: "PUT"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/users/profile", V2: "PUT"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/users/email", V2: "PUT"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/users/phone", V2: "PUT"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/users/password", V2: "PUT"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/users/avatar", V2: "PUT"},

		{Ptype: "p", V0: "admin", V1: "/api/v1/categories", V2: "GET"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/categories", V2: "POST"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/categories", V2: "PUT"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/categories", V2: "DELETE"},

		{Ptype: "p", V0: "admin", V1: "/api/v1/roles/id", V2: "GET"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/roles", V2: "GET"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/roles", V2: "POST"},

		// 普通用户
		{Ptype: "p", V0: "user", V1: "/api/v1/users", V2: "GET"},
		{Ptype: "p", V0: "user", V1: "/api/v1/users/profile", V2: "PUT"},
		{Ptype: "p", V0: "user", V1: "/api/v1/users/email", V2: "PUT"},
		{Ptype: "p", V0: "user", V1: "/api/v1/users/phone", V2: "PUT"},
		{Ptype: "p", V0: "user", V1: "/api/v1/users/password", V2: "PUT"},
		{Ptype: "p", V0: "user", V1: "/api/v1/users/avatar", V2: "PUT"},

		{Ptype: "p", V0: "user", V1: "/api/v1/categories", V2: "GET"},
		{Ptype: "p", V0: "user", V1: "/api/v1/categories", V2: "POST"},
		{Ptype: "p", V0: "user", V1: "/api/v1/categories", V2: "PUT"},
		{Ptype: "p", V0: "user", V1: "/api/v1/categories", V2: "DELETE"},

		// 游客
		{Ptype: "p", V0: "guest", V1: "/api/v1/users", V2: "GET"},
		{Ptype: "p", V0: "guest", V1: "/api/v1/categories", V2: "GET"},
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
