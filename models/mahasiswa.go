package models

import "gorm.io/gorm"

type Mahasiswa struct {
	gorm.Model
	Nama      string `gorm:"type:varchar(30)" json:"nama"`
	KodeKelas string `gorm:"type:varchar(3)" json:"kode_kelas"`
}
