package common

import "time"

type Model struct {
	Id        int        `json:"id" gorm:"column:id; AUTO_INCREMENT;"`
	CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;not null;default:current_timestamp;comment:创建时间"`
	CreatedBy int        `json:"createdBy" gorm:"column:created_by;type:int; unsigned;not null;default:0"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updated_at;type:datetime;not null;default:current_timestamp;autoUpdateTime;comment:更新时间"`
	UpdatedBy int        `json:"updatedBy" gorm:"column:updated_by;type:int; unsigned;not null;default:0"`
}
