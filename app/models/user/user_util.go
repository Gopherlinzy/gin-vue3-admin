package user

import (
	"github.com/Gopherlinzy/gin-vue3-admin/app/models/menu"
	"github.com/Gopherlinzy/gin-vue3-admin/app/models/role"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/app"
	casbins "github.com/Gopherlinzy/gin-vue3-admin/pkg/casbin"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/database"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/paginator"
	"github.com/gin-gonic/gin"
)

// IsEmailExist 判断 Email 是否被注册
func IsEmailExist(email string) bool {
	var count int64
	database.Gohub_DB.Model(&User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// IsPhoneExist 判断 Phone 是否被注册
func IsPhoneExist(phone string) bool {
	var count int64
	database.Gohub_DB.Model(&User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

func All() (users []User) {
	database.Gohub_DB.Find(&users)
	return
}

// GetByMulti 通过 手机号/Email/用户名 来获取用户
func GetByMulti(loginID string) (userModel User) {
	database.Gohub_DB.Where("name = ?", loginID).
		Or("phone = ?", loginID).
		Or("email = ?", loginID).
		First(&userModel)
	return
}

// GetByPhone 通过手机号来获取用户
func GetByPhone(phone string) (userModel User) {
	database.Gohub_DB.Where("phone = ?", phone).First(&userModel)
	return
}

// Get 通过 ID 获取用户
func Get(idstr string) (userModel User) {
	database.Gohub_DB.Where("id = ?", idstr).First(&userModel)
	return
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int) (users []User, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.Gohub_DB.Model(&User{}),
		&users,
		app.V1URL(database.TableName(&User{})),
		perPage,
	)
	return
}

// GetMenus 获取用户角色的菜单权限
func GetMenus(userName string) (menus []menu.Menu) {
	r := casbins.NewCasbin().GetRolesForUser(userName)[0]
	//fmt.Println(r)
	roleModel := role.GetBy("role_name", r)
	menus = roleModel.GetAssociationsMenus()
	return
}
