package models

import "gorm.io/gorm"

type Kelas struct {
	gorm.Model
	KodeKelas string `gorm:"index" json:"kode_kelas"`
	Nama      string `gorm:"type:varchar(30)" json:"nama"`
}
