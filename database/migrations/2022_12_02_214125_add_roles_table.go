package migrations

import (
	"database/sql"
	"github.com/Gopherlinzy/gin-vue3-admin/app/models"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/migrate"

	"gorm.io/gorm"
)

type Role struct {
	models.BaseModel

	RoleName string `gorm:"type:varchar(255);not null;unique"`
	Des      string `gorm:"type:varchar(255);default:null"`
	Status   bool   `gorm:"type:TINYINT(1);default:1"`

	Menus []Menu `gorm:"many2many:roles_menus;"`
	Apis  []Api  `gorm:"many2many:roles_apis;"`

	models.CommonTimestampsField
}

func init() {

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Role{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Role{})
	}

	migrate.Add("2022_12_02_214125_add_roles_table", up, down)
}
