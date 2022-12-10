package migrations

import (
	"database/sql"
	"github.com/Gopherlinzy/gin-vue3-admin/app/models"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/migrate"

	"gorm.io/gorm"
)

type Menu struct {
	models.BaseModel

	Name        string `gorm:"type:varchar(50);not null;index"`
	Permissions string `gorm:"type:varchar(50);not null;index;unique"`
	RouterName  string `gorm:"type:varchar(255);not null"`
	RouterPath  string `gorm:"type:varchar(255);not null"`
	FatherID    uint64 `gorm:"type:int;default:0"`
	VuePath     string `gorm:"type:varchar(255)"`
	Status      bool   `gorm:"type:TINYINT(1);default:1"`

	//Children []Menu `json:"children" gorm:"-"`

	models.CommonTimestampsField
}

func init() {

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Menu{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Menu{})
	}

	migrate.Add("2022_12_07_210542_add_menus_table", up, down)
}
