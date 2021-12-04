package pemberitahuan

import (
	"time"

	"gorm.io/gorm"
)

type Pemberitahuans struct {
	Id        			uint           `gorm:"primaryKey" json:"id"`
	No_Pemberitahuan    string         `json:"pemberitahuan"`
	CreatedAt time.Time      		`json:"createdAt"`
	UpdatedAt time.Time     		 `json:"updatedAt"`
	DeletedAt gorm.DeletedAt 		`gorm:"index" json:"deletedAt"`
}
