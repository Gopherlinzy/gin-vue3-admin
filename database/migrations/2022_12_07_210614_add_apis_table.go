package migrations

import (
	"database/sql"
	"github.com/Gopherlinzy/gin-vue3-admin/app/models"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/migrate"

	"gorm.io/gorm"
)

type Api struct {
	models.BaseModel

	Path        string `gorm:"type:varchar(255);not null"`     // api路径
	Description string `gorm:"type:varchar(255);default:null"` // api中文描述
	ApiGroup    string `gorm:"type:varchar(255);not null"`     // api组
	Method      string `gorm:"type:varchar(50);not null"`      // 方法:创建POST(默认)|查看GET|更新PUT|删除DELETE

	models.CommonTimestampsField
}

func init() {

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Api{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Api{})
	}

	migrate.Add("2022_12_07_210614_add_apis_table", up, down)
}
