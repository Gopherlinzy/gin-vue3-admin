package v1

import (
	"github.com/Gopherlinzy/gohub/app/models/category"
	"github.com/Gopherlinzy/gohub/app/requests"
	"github.com/Gopherlinzy/gohub/pkg/response"
	"github.com/gin-gonic/gin"
)

type CategoriesController struct {
	BaseApiController
}

func (ctrl *CategoriesController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := category.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (ctrl *CategoriesController) Show(c *gin.Context) {
	categoryModel := category.Get(c.Param("id"))
	if categoryModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, categoryModel)
}

func (ctrl *CategoriesController) Store(c *gin.Context) {

	request := requests.CategorySaveRequest{}
	if ok := requests.Validate(c, &request, requests.CategorySave); !ok {
		return
	}

	categoryModel := category.Category{
		Name:        request.Name,
		Description: request.Description,
	}
	categoryModel.Create()
	if categoryModel.ID > 0 {
		response.Created(c, categoryModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *CategoriesController) Update(c *gin.Context) {
	// 表单验证
	request := requests.CategorySaveRequest{}

	if bindOk := requests.Validate(c, &request, requests.CategorySave); !bindOk {
		return
	}

	categoryModel := category.Get(request.ID)

	categoryModel.Name = request.Name
	categoryModel.Description = request.Description
	// 保存数据
	rowsAffected := categoryModel.Save()
	if rowsAffected > 0 {
		response.Data(c, categoryModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *CategoriesController) Delete(c *gin.Context) {

	// 表单验证
	request := requests.CategoryDeleteRequest{}

	if bindOk := requests.Validate(c, &request, requests.CategoryDelete); !bindOk {
		return
	}

	categoryModel := category.Get(request.ID)

	rowsAffected := categoryModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
