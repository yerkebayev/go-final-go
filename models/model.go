package models

import (
	"gorm.io/gorm"
	"time"
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
	Students    []Student `gorm:"many2many:student_sessions;"`
	Attendances []Attendance
}

type Attendance struct {
	gorm.Model
	SessionID uint
	StudentID uint
	Present   bool
	Time      time.Time `gorm:"not null"`
}

type Image struct {
	gorm.Model
	UserID    int
	ImageData []byte `gorm:"type:bytea"`
}
