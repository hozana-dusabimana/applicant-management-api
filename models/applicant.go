package models

import (
	"time"
	"gorm.io/gorm"
)

type Applicant struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
	
	Name     string `json:"name" gorm:"not null;size:100"`
	Email    string `json:"email" gorm:"unique;not null;size:150"`
	Position string `json:"position" gorm:"not null;size:100"`
	Status   string `json:"status" gorm:"default:'pending';size:20"`
	Phone    string `json:"phone,omitempty" gorm:"size:20"`
	Resume   string `json:"resume,omitempty" gorm:"type:text"`
	Notes    string `json:"notes,omitempty" gorm:"type:text"`
}

// TableName returns the table name for the Applicant model
func (Applicant) TableName() string {
	return "applicants"
}
