package models

import (
	"time"
)

// PasswordResetToken represents the password_reset_tokens table in the database
// D8 Table password_reset_tokens - ใช้สำหรับเก็บโทเคนชั่วคราวสำหรับรีเซ็ตรหัสผ่าน
type PasswordResetToken struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"not null;index" json:"user_id" binding:"required"`
	Token     string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"token" binding:"required"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at" binding:"required"`
	CreatedAt time.Time `json:"created_at"`

	// Relationships
	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
}

// TableName specifies the table name for PasswordResetToken model
func (PasswordResetToken) TableName() string {
	return "password_reset_tokens"
}

// IsExpired checks if the token has expired
func (p *PasswordResetToken) IsExpired() bool {
	return time.Now().After(p.ExpiresAt)
}

// IsValid checks if the token is still valid (not expired)
func (p *PasswordResetToken) IsValid() bool {
	return !p.IsExpired()
}

