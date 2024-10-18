package organization

import (
	"gorm.io/gorm"
	"time"
)

type Organization struct {
	ID                    string             `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name                  string             `gorm:"column:name;type:varchar(255)" json:"name"`
	Slug                  string             `gorm:"column:slug;type:varchar(255)" json:"slug"`
	BillingEmail          string             `gorm:"column:billing_email;type:varchar(255)" json:"billingEmail"`
	URL                   string             `gorm:"column:url;type:varchar(255)" json:"url"`
	About                 string             `gorm:"column:about;type:text" json:"about"`
	LogoURL               string             `gorm:"column:logo_url;type:varchar(255)" json:"logoUrl"`
	CreatedBy             string             `gorm:"column:created_by;type:varchar(255);default:null" json:"createdBy"`
	BackgroundImageURL    string             `gorm:"column:background_image_url;type:varchar(255)" json:"backgroundImageUrl"`
	Plan                  string             `gorm:"column:plan;type:varchar(255)" json:"plan"`
	CurrentPeriodStartsAt *time.Time         `gorm:"column:current_period_starts_at" json:"currentPeriodStartsAt"`
	CurrentPeriodEndsAt   *time.Time         `gorm:"column:current_period_ends_at" json:"currentPeriodEndsAt"`
	SubscriptionID        string             `gorm:"column:subscription_id;type:varchar(255)" json:"subscriptionId"`
	Status                string             `gorm:"column:status;type:varchar(255)" json:"status"`
	Timezone              string             `gorm:"column:timezone;type:varchar(255)" json:"timezone"`
	CreatedAt             time.Time          `gorm:"column:created_at;not null" json:"createdAt"`
	UpdatedAt             time.Time          `gorm:"column:updated_at;not null" json:"updatedAt"`
	DeletedAt             gorm.DeletedAt     `gorm:"column:deleted_at;index" json:"deletedAt"`
	UserOrganizations     []UserOrganization `gorm:"foreignKey:OrganizationID" json:"userOrganizations"`
}

func (Organization) TableName() string {
	return "organizations"
}
