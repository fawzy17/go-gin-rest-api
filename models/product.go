package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name     string  `gorm:"type:varchar(30)" json:"name"`
	Price    float64 `gorm:"type:float" json:"price"`
	Quantity int64   `gorm:"type:int" json:"quantity"`
}
