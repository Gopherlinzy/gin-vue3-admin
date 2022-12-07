package seeders

import (
	"fmt"
	"github.com/Gopherlinzy/gin-vue3-admin/database/factories"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/console"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/logger"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/seed"

	"gorm.io/gorm"
)

func init() {

	seed.Add("SeedCategoriesTable", func(db *gorm.DB) {

		categories := factories.MakeCategories(10)

		result := db.Table("categories").Create(&categories)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
