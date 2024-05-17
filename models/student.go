package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	StudentNumberId string
	Name            string
	Attendances     []Attendance
}
