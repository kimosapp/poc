package organization

type OrganizationUser struct {
	ID             string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	OrganizationID string `gorm:"column:organization_id"`
	UserID         string `gorm:"column:user_id"`
}

func (OrganizationUser) TableName() string {
	return "organization_users"
}
