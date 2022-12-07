package migrations

import (
	"database/sql"
	"github.com/Gopherlinzy/gin-vue3-admin/app/models"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/migrate"

	"gorm.io/gorm"
)

type Link struct {
	models.BaseModel

	Name string `gorm:"type:varchar(255);not null"`
	URL  string `gorm:"type:varchar(255);default:null"`

	models.CommonTimestampsField
}

func init() {

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Link{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Link{})
	}

	migrate.Add("2022_11_25_152436_add_links_table", up, down)
}
