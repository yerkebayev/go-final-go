package models

import (
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	Title    string
	Sessions []Session
	Teachers []Teacher `gorm:"many2many:teacher_courses;"`
}

type Session struct {
	gorm.Model
	TeacherID   uint
	CourseID    uint
	Date        string
	Attendances []Attendance
}

type Attendance struct {
	gorm.Model
	SessionID uint
	StudentID uint
	Time      string
}

type Image struct {
	gorm.Model
	UserID    int
	ImageData []byte `gorm:"type:bytea"`
}
