package api

import (
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/app"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/database"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (api Api) {
	database.Gohub_DB.Where("id = ?", idstr).First(&api)
	return
}

func GetBy(field, value string) (api Api) {
	database.Gohub_DB.Where(field+" = ?", value).First(&api)
	return
}

// ExceptID 排除 ID 序列的数据
func ExceptID(value []uint64) (apis []Api) {
	if len(value) == 0 {
		return All()
	}
	query := database.Gohub_DB.Model(&Api{})
	for _, v := range value {
		query.Where("id != ?", v)
	}
	query.Find(&apis)
	return
}

// AcceptID 关联 id 序列的数据
func AcceptID(value []uint64, l uint64) (apis []Api) {
	if len(value) == 0 {
		return
	}
	var re []uint64
	i, j := uint64(1), 0
	for ; i <= l; i++ {
		if j < len(value) && i == value[j] {
			j++
			continue
		}
		re = append(re, i)
	}
	//fmt.Println(re)
	apis = ExceptID(re)
	return
}

func All() (apis []Api) {
	database.Gohub_DB.Find(&apis)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.Gohub_DB.Model(Api{}).Where(field+" = ?", value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (apis []Api, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.Gohub_DB.Model(Api{}),
		&apis,
		app.V1URL(database.TableName(&Api{})),
		perPage,
	)
	return
}
