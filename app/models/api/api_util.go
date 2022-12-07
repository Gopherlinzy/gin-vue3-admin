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
    database.Gohub_DB.Where(field + " = ?", value).First(&api)
    return
}

func All() (apis []Api) {
    database.Gohub_DB.Find(&apis)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.Gohub_DB.Model(Api{}).Where(field + " = ?", value).Count(&count)
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