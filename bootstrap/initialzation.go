package bootstrap

import (
	"github.com/Gopherlinzy/gohub/app/models/role"
	casbins "github.com/Gopherlinzy/gohub/pkg/casbin"
	"github.com/Gopherlinzy/gohub/pkg/database"
	"github.com/Gopherlinzy/gohub/pkg/logger"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"sync"
)

var once sync.Once

func IntizationData() {
	once.Do(func() {
		initRoleData()
		initCasbinData()
	})
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
		{Ptype: "g", V0: "sly", V1: "superAdmin"},
		{Ptype: "g", V0: "linzy", V1: "admin"},
		{Ptype: "g", V0: "gg", V1: "user"},

		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/users", V2: "GET"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/users", V2: "POST"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/users/role", V2: "POST"},
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

		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/roles", V2: "GET"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/roles", V2: "POST"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/roles", V2: "PUT"},
		{Ptype: "p", V0: "superAdmin", V1: "/api/v1/roles", V2: "DELETE"},

		{Ptype: "p", V0: "admin", V1: "/api/v1/users", V2: "GET"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/users", V2: "POST"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/users/profile", V2: "PUT"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/users/email", V2: "PUT"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/users/phone", V2: "PUT"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/users/password", V2: "PUT"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/users/avatar", V2: "PUT"},

		{Ptype: "p", V0: "admin", V1: "/api/v1/categories", V2: "GET"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/categories", V2: "POST"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/categories", V2: "PUT"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/categories", V2: "DELETE"},

		{Ptype: "p", V0: "admin", V1: "/api/v1/roles", V2: "GET"},
		{Ptype: "p", V0: "admin", V1: "/api/v1/roles", V2: "POST"},

		{Ptype: "p", V0: "user", V1: "/api/v1/users", V2: "GET"},
		{Ptype: "p", V0: "user", V1: "/api/v1/users", V2: "POST"},
		{Ptype: "p", V0: "user", V1: "/api/v1/users/role", V2: "POST"},
		{Ptype: "p", V0: "user", V1: "/api/v1/users/profile", V2: "PUT"},
		{Ptype: "p", V0: "user", V1: "/api/v1/users/email", V2: "PUT"},
		{Ptype: "p", V0: "user", V1: "/api/v1/users/phone", V2: "PUT"},
		{Ptype: "p", V0: "user", V1: "/api/v1/users/password", V2: "PUT"},
		{Ptype: "p", V0: "user", V1: "/api/v1/users/avatar", V2: "PUT"},
		{Ptype: "p", V0: "user", V1: "/api/v1/users", V2: "DELETE"},

		{Ptype: "p", V0: "user", V1: "/api/v1/categories", V2: "GET"},
		{Ptype: "p", V0: "user", V1: "/api/v1/categories", V2: "POST"},
		{Ptype: "p", V0: "user", V1: "/api/v1/categories", V2: "PUT"},
		{Ptype: "p", V0: "user", V1: "/api/v1/categories", V2: "DELETE"},
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
		{RoleName: "superAdmin", Des: "超级管理员:拥有所有权限", Status: true},
		{RoleName: "admin", Des: "管理员", Status: true},
		{RoleName: "user", Des: "普通用户", Status: true},
		{RoleName: "guest", Des: "游客模式", Status: false},
	}

	database.Gohub_DB.Create(&roles)
}
