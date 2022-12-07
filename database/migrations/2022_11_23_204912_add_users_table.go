package migrations

import (
	"database/sql"
	"github.com/Gopherlinzy/gin-vue3-admin/app/models"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/migrate"

	"gorm.io/gorm"
)

type User struct {
	models.BaseModel

	Name         string `gorm:"type:varchar(255);not null;index"`
	Email        string `gorm:"type:varchar(255);index;default:null"`
	Phone        string `gorm:"type:varchar(20);index;default:null"`
	Password     string `gorm:"type:varchar(255)"`
	City         string `gorm:"type:varchar(10);"`
	Introduction string `gorm:"type:varchar(255);"`
	Avatar       string `gorm:"type:varchar(255);default:null"`
	Status       bool   `gorm:"type:TINYINT(1);default:1"`

	RoleName string `gorm:"type:varchar(255);not null"`

	models.CommonTimestampsField
}

func init() {

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&User{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&User{})
	}

	migrate.Add("2022_11_23_204912_add_users_table", up, down)
}
