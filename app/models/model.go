package models

import (
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"time"
)

// BaseModel 模型基础类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

// CommonTimestampsField 时间戳
type CommonTimestampsField struct {
	CreatedAt time.Time      `gorm:"type:datetime(0);column:created_at;index;autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt time.Time      `gorm:"type:datetime(0);column:updated_at;index;autoUpdateTime" json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index;" json:"deleted_at,omitempty"`
}

// GetStringID 获取 ID 的字符串格式
func (a BaseModel) GetStringID() string {
	return cast.ToString(a.ID)
}
