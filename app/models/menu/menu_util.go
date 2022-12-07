package menu

import (
    "github.com/Gopherlinzy/gin-vue3-admin/pkg/app"
    "github.com/Gopherlinzy/gin-vue3-admin/pkg/database"
    "github.com/Gopherlinzy/gin-vue3-admin/pkg/paginator"

    "github.com/gin-gonic/gin"
)

func Get(idstr string) (menu Menu) {
    database.Gohub_DB.Where("id = ?", idstr).First(&menu)
    return
}

func GetBy(field, value string) (menu Menu) {
    database.Gohub_DB.Where(field + " = ?", value).First(&menu)
    return
}

func All() (menus []Menu) {
    database.Gohub_DB.Find(&menus)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.Gohub_DB.Model(Menu{}).Where(field + " = ?", value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (menus []Menu, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.Gohub_DB.Model(Menu{}),
        &menus,
        app.V1URL(database.TableName(&Menu{})),
        perPage,
    )
    return
}