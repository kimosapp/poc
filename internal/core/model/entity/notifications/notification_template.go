package notifications

type NotificationTemplate struct {
	ID                  string `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name                string `gorm:"column:name;type:varchar(255)" json:"name"`
	Flow                string `gorm:"column:flow;type:varchar(255)" json:"flow"`
	NotificationChannel string `gorm:"column:notification_channel;type:varchar(255)" json:"notificationChannel"`
	SenderAccount       string `gorm:"column:sender_account;type:varchar(255)" json:"senderAccount"`
	Subject             string `gorm:"column:subject;type:varchar(255)" json:"subject"`
	Body                string `gorm:"column:body;type:text" json:"body"`
	Language            string `gorm:"column:language;type:varchar(255)" json:"language"`
	IsActive            bool   `gorm:"column:is_active;type:boolean" json:"isActive"`
	CreatedAt           string `gorm:"column:created_at;not null" json:"createdAt"`
	UpdatedAt           string `gorm:"column:updated_at;not null" json:"updatedAt"`
}

func (NotificationTemplate) TableName() string {
	return "notification_templates"
}
