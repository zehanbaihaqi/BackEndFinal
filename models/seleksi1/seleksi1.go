package seleksi1

import (
	"time"

	"gorm.io/gorm"
)

type Seleksi1 struct {
	Id        			uint           `gorm:"primaryKey" json:"id"`
	Seleksi1    string         `json:"seleksi1"`
	CreatedAt time.Time      		`json:"createdAt"`
	UpdatedAt time.Time     		 `json:"updatedAt"`
	DeletedAt gorm.DeletedAt 		`gorm:"index" json:"deletedAt"`
}
