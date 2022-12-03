package migrations

import (
    "database/sql"
    "github.com/Gopherlinzy/gohub/app/models"
    "github.com/Gopherlinzy/gohub/pkg/migrate"

    "gorm.io/gorm"
)

func init() {

    type Role struct {
        models.BaseModel

        RoleName string `gorm:"type:varchar(255);not null"`
        Des      string `gorm:"type:varchar(255);index;default:null"`

        models.CommonTimestampsField
    }

    up := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.AutoMigrate(&Role{})
    }

    down := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.DropTable(&Role{})
    }

    migrate.Add("2022_12_02_214125_add_roles_table", up, down)
}