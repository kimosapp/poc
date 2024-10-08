package organization

import "time"

type User struct {
	ID        string    `json:"id"`
	LastName  string    `gorm:"column:last_name;type:varchar(255);" json:"lastName"`
	FirstName string    `gorm:"column:first_name;type:varchar(255);" json:"firstName"`
	Email     string    `gorm:"column:email;type:varchar(255);" json:"email"`
	ImageUrl  string    `gorm:"column:image_url;type:varchar(255);" json:"imageUrl"`
	LastLogin time.Time `gorm:"column:last_login;"  json:"lastLogin"`
	CreatedAt time.Time `gorm:"column:created_at;" json:"createdAt"`
	DeletedAt time.Time `gorm:"column:deleted_at;" json:"deletedAt"`
	UpdatedAt time.Time `gorm:"column:updated_at;" json:"updatedAt"`
}

func (User) TableName() string {
	return "profiles"
}
