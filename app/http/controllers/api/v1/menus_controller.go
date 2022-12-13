package v1

import (
	"github.com/Gopherlinzy/gin-vue3-admin/app/models/menu"
	"github.com/Gopherlinzy/gin-vue3-admin/app/requests"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MenusController struct {
	BaseApiController
}

// Index 显示 menu 列表
func (ctrl *MenusController) Index(c *gin.Context) {
	menus := menu.All()
	response.Data(c, menus)
}

// IndexPagination 显示 menus 分页列表
func (ctrl *MenusController) IndexPagination(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := menu.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

// Show 显示单个 menu 数据
func (ctrl *MenusController) Show(c *gin.Context) {
	request := requests.MenuIDRequest{}

	if bindOk := requests.Validate(c, &request, requests.MenuID); !bindOk {
		return
	}
	menuModel := menu.Get(request.ID)

	response.Data(c, menuModel)
}

// Store 新建 menu 数据
func (ctrl *MenusController) Store(c *gin.Context) {

	request := requests.MenuSaveRequest{}
	if ok := requests.Validate(c, &request, requests.MenuSave); !ok {
		return
	}

	faID, _ := strconv.Atoi(request.FatherID)
	status, _ := strconv.ParseBool(request.Status)

	menuModel := menu.Menu{
		Name:       request.Name,
		Permission: request.Permission,
		RouterName: request.RouterName,
		RouterPath: request.RouterPath,
		FatherID:   uint64(faID),
		VuePath:    request.VuePath,
		Status:     status,
	}
	menuModel.Create()
	if menuModel.ID > 0 {
		response.Created(c, menuModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

// Update 修改 menu 数据
func (ctrl *MenusController) Update(c *gin.Context) {

	request := requests.MenuSaveRequest{}

	if bindOk := requests.Validate(c, &request, requests.MenuSave); !bindOk {
		return
	}

	menuModel := menu.Get(request.ID)

	faID, _ := strconv.Atoi(request.FatherID)
	status, _ := strconv.ParseBool(request.Status)

	menuModel.Name = request.Name
	menuModel.Permission = request.Permission
	menuModel.RouterName = request.RouterName
	menuModel.RouterPath = request.RouterPath
	menuModel.FatherID = uint64(faID)
	menuModel.VuePath = request.VuePath
	menuModel.Status = status

	rowsAffected := menuModel.Save()
	if rowsAffected > 0 {
		response.Data(c, menuModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

// Delete 删除 menu 数据
func (ctrl *MenusController) Delete(c *gin.Context) {

	// 表单验证
	request := requests.MenuIDRequest{}

	if bindOk := requests.Validate(c, &request, requests.MenuID); !bindOk {
		return
	}

	menuModel := menu.Get(request.ID)

	rowsAffected := menuModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
