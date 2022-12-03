package role

import (
    "github.com/Gopherlinzy/gohub/pkg/app"
    "github.com/Gopherlinzy/gohub/pkg/database"
    "github.com/Gopherlinzy/gohub/pkg/paginator"

    "github.com/gin-gonic/gin"
)

func Get(idstr string) (role Role) {
    database.Gohub_DB.Where("id", idstr).First(&role)
    return
}

func GetBy(field, value string) (role Role) {
    database.Gohub_DB.Where(field + " = ?", value).First(&role)
    return
}

func All() (roles []Role) {
    database.Gohub_DB.Find(&roles)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.Gohub_DB.Model(Role{}).Where(field + " = ?", value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (roles []Role, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.Gohub_DB.Model(Role{}),
        &roles,
        app.V1URL(database.TableName(&Role{})),
        perPage,
    )
    return
}