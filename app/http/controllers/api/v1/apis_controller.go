package v1

import (
    "github.com/Gopherlinzy/gin-vue3-admin/app/models/api"
    "github.com/Gopherlinzy/gin-vue3-admin/app/requests"
    "github.com/Gopherlinzy/gin-vue3-admin/pkg/response"

    "github.com/gin-gonic/gin"
)

type ApisController struct {
    BaseApiController
}

// Index 显示 api 列表
func (ctrl *ApisController) Index(c *gin.Context) {
    apis := api.All()
    response.Data(c, apis)
}

// Show 显示单个 api 数据
func (ctrl *ApisController) Show(c *gin.Context) {
    request := requests.ApiIDRequest{}

    if bindOk := requests.Validate(c, &request, requests.ApiID); !bindOk {
        return
    }
    apiModel := api.Get(request.ID)

    response.Data(c, apiModel)
}

// Store 新建 api 数据
func (ctrl *ApisController) Store(c *gin.Context) {

    request := requests.ApiSaveRequest{}
    if ok := requests.Validate(c, &request, requests.ApiSave); !ok {
        return
    }

    apiModel := api.Api{
        Path:        request.Path,
        ApiGroup:    request.ApiGroup,
        Description: request.Description,
        Method:      request.Method,
    }

    apiModel.Create()
    if apiModel.ID > 0 {
        response.Created(c, apiModel)
    } else {
        response.Abort500(c, "创建失败，请稍后尝试~")
    }
}

// Update 修改 api 数据
func (ctrl *ApisController) Update(c *gin.Context) {

    request := requests.ApiSaveRequest{}

    if bindOk := requests.Validate(c, &request, requests.ApiSave); !bindOk {
        return
    }

    apiModel := api.Get(request.ID)

    apiModel.Path = request.Path
    apiModel.ApiGroup = request.ApiGroup
    apiModel.Description = request.Description
    apiModel.Method = request.Method
    
    rowsAffected := apiModel.Save()
    if rowsAffected > 0 {
        response.Data(c, apiModel)
    } else {
        response.Abort500(c, "更新失败，请稍后尝试~")
    }
}

// Delete 删除 api 数据
func (ctrl *ApisController) Delete(c *gin.Context) {

    // 表单验证
    request := requests.ApiIDRequest{}

    if bindOk := requests.Validate(c, &request, requests.ApiID); !bindOk {
        return
    }

    apiModel := api.Get(request.ID)

    rowsAffected := apiModel.Delete()
    if rowsAffected > 0 {
        response.Success(c)
        return
    }

    response.Abort500(c, "删除失败，请稍后尝试~")
}
