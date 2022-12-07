package migrations

import (
	"database/sql"
	"github.com/Gopherlinzy/gin-vue3-admin/app/models"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/migrate"

	"gorm.io/gorm"
)

type Api struct {
	models.BaseModel

	Path    string `gorm:"type:varchar(255);not null"`
	Group   string `gorm:"type:varchar(50);not null"`
	Des     string `gorm:"type:varchar(255);default:null"`
	Request string `gorm:"type:varchar(255);not null"`

	Roles []Role `gorm:"many2many:roles_menus;"`

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
