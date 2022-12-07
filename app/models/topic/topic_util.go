package topic

import (
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/app"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/database"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/paginator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func Get(idstr string) (topic Topic) {
	database.Gohub_DB.Preload(clause.Associations).Where("id", idstr).First(&topic)
	return
}

func GetBy(field, value string) (topic Topic) {
	database.Gohub_DB.Where(field+" = ?", value).First(&topic)
	return
}

func All() (topics []Topic) {
	database.Gohub_DB.Find(&topics)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.Gohub_DB.Model(Topic{}).Where(field+" = ?", value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (topics []Topic, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.Gohub_DB.Model(Topic{}),
		&topics,
		app.V1URL(database.TableName(&Topic{})),
		perPage,
	)
	return
}
