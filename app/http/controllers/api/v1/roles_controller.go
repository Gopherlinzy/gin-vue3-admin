package v1

import (
	"github.com/Gopherlinzy/gin-vue3-admin/app/models/role"
	"github.com/Gopherlinzy/gin-vue3-admin/app/requests"
	casbins "github.com/Gopherlinzy/gin-vue3-admin/pkg/casbin"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RolesController struct {
	BaseApiController
}

// Index 显示 role 列表
func (ctrl *RolesController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := role.Paginate(c, 5)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

// IndexPolicies 显示该 role 所有的权限
func (ctrl *RolesController) IndexPolicies(c *gin.Context) {
	request := requests.RoleIDRequest{}

	if bindOk := requests.Validate(c, &request, requests.RoleID); !bindOk {
		return
	}

	roleModel := role.Get(request.ID)
	//fmt.Println("-------", request)
	policy := casbins.NewCasbin().GetFilteredPolicy(roleModel.RoleName)
	response.Data(c, policy)
}

// Show 显示单个 role 数据
func (ctrl *RolesController) Show(c *gin.Context) {
	request := requests.RoleIDRequest{}

	if bindOk := requests.Validate(c, &request, requests.RoleID); !bindOk {
		return
	}

	roleModel := role.Get(request.ID)
	response.Data(c, roleModel)
}

// ShowMenus 显示角色关联的 menus
func (ctrl *RolesController) ShowMenus(c *gin.Context) {
	request := requests.RoleIDRequest{}

	if bindOk := requests.Validate(c, &request, requests.RoleID); !bindOk {
		return
	}
	roleModel := role.GetByAssociated("Menus", request.ID)
	//fmt.Println(roleModel)
	response.Data(c, roleModel.Menus)
}

// ShowApis 显示角色关联的 apis
func (ctrl *RolesController) ShowApis(c *gin.Context) {
	request := requests.RoleIDRequest{}

	if bindOk := requests.Validate(c, &request, requests.RoleID); !bindOk {
		return
	}
	roleModel := role.GetByAssociated("Apis", request.ID)
	//fmt.Println(roleModel)
	response.Data(c, roleModel.Apis)
}

// Store 新建 role 数据
func (ctrl *RolesController) Store(c *gin.Context) {

	request := requests.RoleStoreRequest{}
	if ok := requests.Validate(c, &request, requests.RoleStore); !ok {
		return
	}

	roleModel := role.Role{
		RoleName: request.RoleName,
		Des:      request.Des,
	}
	roleModel.Create()
	if roleModel.ID > 0 {
		response.Created(c, roleModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

// Update 修改 role 数据
func (ctrl *RolesController) Update(c *gin.Context) {

	request := requests.RoleUpdateRequest{}

	if bindOk := requests.Validate(c, &request, requests.RoleUpdate); !bindOk {
		return
	}

	roleModel := role.Get(request.ID)

	roleModel.Des = request.Des
	rowsAffected := roleModel.Save()
	if rowsAffected > 0 {
		response.Data(c, roleModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

// SetMenuPermissions 添加 role_id 的 菜单权限
func (ctrl *RolesController) SetMenuPermissions(c *gin.Context) {
	request := requests.RolePermissionsRequest{}
	if ok := requests.Validate(c, &request, requests.RolePermissions); !ok {
		return
	}
	roleModel := role.Get(request.ID)

	err := roleModel.AppendAssociation("Menus", request.Permissions)
	if err != nil {
		response.Abort500(c, "创建失败，请稍后尝试~")
	} else {
		response.Data(c, request)
	}

}

// SetApiPolicies 添加 role_id 的 api权限
func (ctrl *RolesController) SetApiPolicies(c *gin.Context) {
	request := requests.RoleApiPolicyRequest{}
	if ok := requests.Validate(c, &request, requests.RoleApiPolicy); !ok {
		return
	}
	roleModel := role.Get(request.ID)

	err := roleModel.AppendAssociation("Apis", request.ApiPolicies)
	if err != nil {
		response.Abort500(c, "创建失败，请稍后尝试~")
	} else {
		response.Data(c, request)
	}
}

// Delete 删除 role 数据
func (ctrl *RolesController) Delete(c *gin.Context) {

	request := requests.RoleIDRequest{}

	if bindOk := requests.Validate(c, &request, requests.RoleID); !bindOk {
		return
	}

	roleModel := role.Get(request.ID)

	roleModel.AssociationClear("Menus")
	roleModel.AssociationClear("Apis")
	casbins.NewCasbin().DeleteRole(roleModel.RoleName)
	rowsAffected := roleModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}

// UpdateRoleStatus 修改角色启用状态
func (ctrl *RolesController) UpdateRoleStatus(c *gin.Context) {
	request := requests.UpdateRoleStatusRequest{}
	if ok := requests.Validate(c, &request, requests.UpdateRoleStatus); !ok {
		return
	}

	roleModel := role.Get(request.ID)

	status, _ := strconv.ParseBool(request.Status)

	roleModel.Status = status

	rowsAffected := roleModel.Save()

	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "更新失败，请稍后尝试~")
}
