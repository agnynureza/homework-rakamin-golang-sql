package models

import "gorm.io/gorm"

type TableMovies struct {
	gorm.Model
	Title       *string `gorm:"type:varchar(255);not null; index"`
	Slug        *string `gorm:"type:varchar(255);unique;not null; index"`
	Description *string `gorm:"type:text;not null"`
	Duration    *uint   `gorm:"type:int(5);not null"`
	Image       *string `gorm:"type:varchar(255);not null"`
}

type Movies struct {
	ID          int    `gorm:"column:id; PRIMARY_KEY" json:"id"`
	Title       string `gorm:"column:title" json:"title" validate:"required"`
	Slug        string `gorm:"column:slug" json:"Slug" validate:"required"`
	Description string `gorm:"column:description" json:"Description" validate:"required"`
	Duration    uint   `gorm:"column:duration" json:"Duration" validate:"required"`
	Image       string `gorm:"column:image" json:"Image" validate:"required"`
}
