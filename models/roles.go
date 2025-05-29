package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Parent struct {
	gorm.Model
	FullName       string
	Email          string `gorm:"uniqueIndex"`
	PhoneNumber    string
	PasswordHash   string
	ProfilePicture string
	Description    string
	Students       []Student `gorm:"foreignKey:ParentID"`
}

type Student struct {
	gorm.Model
	FullName          string
	GradeLevel        string
	Age               int
	ProfilePicture    string
	XP                int
	Missions          []MissionProgress
	ParentID          uint
	AssignedTeacherID *uint // Assigned by principal/admin
}

type Teacher struct {
	gorm.Model
	FullName       string
	Email          string `gorm:"uniqueIndex"`
	PhoneNumber    string
	PasswordHash   string
	ProfilePicture string
	Description    string
	CodeLanguages  pq.StringArray `gorm:"type:text[]"` // e.g., {"Python", "JavaScript"}
	Subjects       pq.StringArray `gorm:"type:text[]"`
	IsApproved     bool
}

type Admin struct {
	gorm.Model
	FullName     string
	Email        string `gorm:"uniqueIndex"`
	PasswordHash string
	Role         string // "Principal", "SuperAdmin" if you want role tiering
}

type MissionProgress struct {
	gorm.Model
	StudentID uint
	MissionID uint
	Status    string // e.g., "in-progress", "completed"
	XP        int
}
