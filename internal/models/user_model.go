package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents the users table in the database
// D1 Table users - ใช้สำหรับเก็บข้อมูลบัญชีผู้ใช้งานและผู้ดูแลระบบ
type User struct {
	ID          uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Username    string         `gorm:"type:varchar(255);uniqueIndex;not null" json:"username" binding:"required"`
	DisplayName string         `gorm:"type:varchar(255);not null" json:"display_name" binding:"required"`
	Email       string         `gorm:"type:varchar(255);uniqueIndex;not null" json:"email" binding:"required,email"`
	Password    string         `gorm:"type:varchar(255);not null" json:"-"` // "-" means this field won't be included in JSON response
	Role        string         `gorm:"type:varchar(50);not null;default:'member'" json:"role"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"` // Soft delete support

	// Relationships
	PasswordResetTokens []PasswordResetToken `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
}

// TableName specifies the table name for User model
func (User) TableName() string {
	return "users"
}

// UserRole constants
const (
	RoleMember = "member"
	RoleAdmin  = "admin"
)

// BeforeCreate hook - will be called before inserting a new user
func (u *User) BeforeCreate(tx *gorm.DB) error {
	// Set default role if not specified
	if u.Role == "" {
		u.Role = RoleMember
	}
	return nil
}

// PublicUser represents the public-facing user data (without sensitive fields)
type PublicUser struct {
	ID          uint      `json:"id"`
	Username    string    `json:"username"`
	DisplayName string    `json:"display_name"`
	Email       string    `json:"email"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
}

// ToPublicUser converts User to PublicUser
func (u *User) ToPublicUser() PublicUser {
	return PublicUser{
		ID:          u.ID,
		Username:    u.Username,
		DisplayName: u.DisplayName,
		Email:       u.Email,
		Role:        u.Role,
		CreatedAt:   u.CreatedAt,
	}
}

// IsAdmin checks if the user is an admin
func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
}

// IsMember checks if the user is a member
func (u *User) IsMember() bool {
	return u.Role == RoleMember
}
