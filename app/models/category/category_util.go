package category

import (
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/app"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/database"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (category Category) {
	database.Gohub_DB.Where("id", idstr).First(&category)
	return
}

func GetBy(field, value string) (category Category) {
	database.Gohub_DB.Where(field+" = ?", value).First(&category)
	return
}

func All() (categories []Category) {
	database.Gohub_DB.Find(&categories)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.Gohub_DB.Model(Category{}).Where(field+" = ?", value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (categories []Category, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.Gohub_DB.Model(Category{}),
		&categories,
		app.V1URL(database.TableName(&Category{})),
		perPage,
	)
	return
}
