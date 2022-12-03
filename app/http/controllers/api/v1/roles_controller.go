package v1

import (
	"github.com/Gopherlinzy/gohub/app/models/role"
	"github.com/Gopherlinzy/gohub/app/requests"
	"github.com/Gopherlinzy/gohub/pkg/response"

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

	data, pager := role.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

// Show 显示单个 role 数据
func (ctrl *RolesController) Show(c *gin.Context) {
	roleModel := role.Get(c.Param("id"))
	if roleModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, roleModel)
}

// Store 新建 role 数据
func (ctrl *RolesController) Store(c *gin.Context) {

	request := requests.RoleRequest{}
	if ok := requests.Validate(c, &request, requests.RoleSave); !ok {
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

	request := requests.RoleRequest{}

	if bindOk := requests.Validate(c, &request, requests.RoleSave); !bindOk {
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

// Delete 删除 role 数据
func (ctrl *RolesController) Delete(c *gin.Context) {

	request := requests.RoleDeleteRequest{}

	if bindOk := requests.Validate(c, &request, requests.RoleDelete); !bindOk {
		return
	}

	roleModel := role.Get(request.ID)

	rowsAffected := roleModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
