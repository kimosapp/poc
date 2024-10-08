package organization

import (
	"gorm.io/gorm"
	"time"
)

type TeamMember struct {
	ID        string         `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	TeamID    string         `gorm:"column:team_id;type:uuid;not null"`
	UserID    string         `gorm:"column:user_id;type:uuid;not null"`
	RoleID    string         `gorm:"column:role_id;type:uuid;not null"`
	IsActive  bool           `gorm:"column:is_active;type:boolean;default:true"`
	Status    string         `gorm:"column:status;type:varchar(255);default:'pending'"`
	InvitedAt time.Time      `gorm:"column:invited_at;not null"`
	CreatedAt time.Time      `gorm:"column:created_at;not null"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (TeamMember) TableName() string {
	return "team_members"
}
