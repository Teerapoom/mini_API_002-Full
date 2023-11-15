package model

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Isbn       string
	Title      string
	Score      float64
	DirectorFK uint
	Director   *Director `gorm:"foreignKey:DirectorFK"`
}

type Director struct {
	DirectorID uint64 `gorm:"primaryKey"`
	Fistname   string
	Lastname   string
}
