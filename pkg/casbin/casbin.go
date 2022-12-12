package casbins

import (
	"fmt"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/database"
	"github.com/casbin/casbin/v2"
	casbinmodel "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

type Casbin struct {
	Enforcer *casbin.CachedEnforcer
}

var (
	once           sync.Once
	internalCasbin *Casbin
)

func NewCasbin() *Casbin {
	once.Do(func() {
		internalCasbin = &Casbin{}
		text := `
			# 请求
			# sub ——> 想要访问资源的用户角色(Subject)——请求实体
			# obj ——> 访问的资源(Object)
			# act ——> 访问的方法(Action: get、post...)
			[request_definition]
			r = sub,obj,act
			
			
			# 策略(.csv文件p的格式，定义的每一行为policy rule;p为policy rule的名字。)
			[policy_definition]
			p = sub,obj,act
			
			# 定义了RBAC中的角色继承关系
			[role_definition]
			g = _, _
			
			
			# 策略效果
			[policy_effect]
			e = some(where (p.eft == allow))
			# 上面表示有任意一条 policy rule 满足, 则最终结果为 allow；p.eft它可以是allow或deny，它是可选的，默认是allow
			
			# 匹配器
			[matchers]
			m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act || checkSuperAdmin(r.sub, "superAdmin")
			`
		m, _ := casbinmodel.NewModelFromString(text)
		//db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4"), &gorm.Config{})
		//从DB获取权限
		orm, err := gormadapter.NewAdapterByDB(database.Gohub_DB)
		if err != nil {
			fmt.Printf("gormadapter init error:%s", err.Error())
			return
		}
		internalCasbin.Enforcer, _ = casbin.NewCachedEnforcer(m, orm)

		//注册超级管理员权限判断
		internalCasbin.Enforcer.AddFunction("checkSuperAdmin", func(arguments ...interface{}) (interface{}, error) {
			username := arguments[0].(string)
			role := arguments[1].(string)
			//fmt.Println(username, role, "--------")
			// 检查用户名的角色是否为superAdmin
			return internalCasbin.Enforcer.HasRoleForUser(username, role)
		})

		internalCasbin.Enforcer.SetExpireTime(60 * 60)
		_ = internalCasbin.Enforcer.LoadPolicy()
		internalCasbin.Enforcer.EnableLog(true)
	})
	return internalCasbin
}

// Enforce 验证是否具有权限 只验证权限不验证角色
func (c *Casbin) Enforce(sub, obj, act string) bool {
	f, _ := c.Enforcer.Enforce(sub, obj, act)
	return f
}

// AddPolicy 添加角色权限
func (c *Casbin) AddPolicy(sub, obj, act string) bool {
	f, _ := c.Enforcer.AddPolicy(sub, obj, act)
	return f
}

// AddPolicies 批量添加角色权限
func (c *Casbin) AddPolicies(roles [][]string) bool {
	f, _ := c.Enforcer.AddPolicies(roles)
	return f
}

// RemovePolicy 删除角色权限
func (c *Casbin) RemovePolicy(sub, obj, act string) bool {
	f, _ := c.Enforcer.RemovePolicy(sub, obj, act)
	return f
}

// UpdatePolicy 修改角色权限
func (c *Casbin) UpdatePolicy(oldsub, oldobj, oldact, newsub, newobj, newact string) bool {
	if ok := c.AddPolicy(newsub, newobj, newact); !ok {
		return false
	}
	if ok := c.RemovePolicy(oldsub, oldobj, oldact); !ok {
		return false
	}
	return true
}

// GetPolicy 查询所有权限
func (c *Casbin) GetPolicy() [][]string {
	return c.Enforcer.GetPolicy()
}

// GetFilteredPolicy 查询指定角色所有权限
func (c *Casbin) GetFilteredPolicy(sub string) [][]string {
	return c.Enforcer.GetFilteredPolicy(0, sub)
}

// DeleteRole 删除角色所有相关权限
func (c *Casbin) DeleteRole(role string) error {
	_, err := c.Enforcer.RemoveFilteredPolicy(0, role)
	return err
}

// AddUserRole 添加用户绑定角色
func (c *Casbin) AddUserRole(sub, group string) error {
	_, err := c.Enforcer.AddGroupingPolicy(sub, group)
	return err
}

// GetRolesForUser 查询用户的角色
func (c *Casbin) GetRolesForUser(sub string) []string {
	roles, err := c.Enforcer.GetRolesForUser(sub)
	if err != nil {
		return []string{}
	}
	return roles
}

// HasRoleForUser 查询用户具有某个角色
func (c *Casbin) HasRoleForUser(sub, group string) error {
	_, err := c.Enforcer.HasRoleForUser(sub, group)
	return err
}

// DeleteUser 删除用户（同删除绑定角色）
func (c *Casbin) DeleteUser(sub, role string) error {
	err := database.Gohub_DB.Where("ptype = 'g' AND v0 = ? AND v1 = ?", sub, role).Delete(&gormadapter.CasbinRule{}).Error
	return err
}

// UpdateUserRole 修改用户绑定角色
func (c *Casbin) UpdateUserRole(oldsub, newsub, oldRole, newRole string) error {
	err := c.AddUserRole(newsub, newRole)
	if err = c.DeleteUser(oldsub, oldRole); err != nil {
		return err
	}
	return err
}
