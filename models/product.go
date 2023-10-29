package models

type Product struct {
	Id       int64   `gorm:"primarykey" json:"id"`
	Name     string  `gorm:"type:varchar(30)" json:"name"`
	Price    float64 `gorm:"type:float" json:"price"`
	Quantity int64   `gorm:"type:int" json:"quantity"`
}
