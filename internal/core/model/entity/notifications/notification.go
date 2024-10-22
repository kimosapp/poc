package notifications

type Notification struct {
	ID        string `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Template  string `gorm:"column:template;type:uuid" json:"template"`
	Recipient string `gorm:"column:recipient;type:varchar(255)" json:"recipient"`
	Subject   string `gorm:"column:subject;type:varchar(255)" json:"subject"`
	Body      string `gorm:"column:body;type:text" json:"body"`
	Language  string `gorm:"column:language;type:varchar(255)" json:"language"`
	CreatedAt string `gorm:"column:created_at;not null" json:"createdAt"`
}
