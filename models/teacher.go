package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Name    string
	Courses []Course `gorm:"many2many:teacher_courses;"`
}
