package link

import (
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/app"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/cache"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/database"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/helpers"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/paginator"
	"time"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (link Link) {
	database.Gohub_DB.Where("id", idstr).First(&link)
	return
}

func GetBy(field, value string) (link Link) {
	database.Gohub_DB.Where(field+" = ?", value).First(&link)
	return
}

func All() (links []Link) {
	database.Gohub_DB.Find(&links)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.Gohub_DB.Model(Link{}).Where(field+" = ?", value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (links []Link, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.Gohub_DB.Model(Link{}),
		&links,
		app.V1URL(database.TableName(&Link{})),
		perPage,
	)
	return
}

func AllCached() (links []Link) {
	// 设置缓存
	cacheKey := "links:all"
	// 设置过期时间
	expireTime := 120 * time.Minute
	// 取数据
	cache.GetObject(cacheKey, &links)

	// 如果数据为空
	if helpers.Empty(links) {
		// 查询数据库
		links = All()
		if helpers.Empty(links) {
			return links
		}
		// 设置缓存
		cache.Set(cacheKey, links, expireTime)
	}
	return
}
