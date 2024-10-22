package notifications

import (
	"github.com/kimosapp/poc/internal/core/model/entity/notifications"
	notificationRepository "github.com/kimosapp/poc/internal/core/ports/repository/notifications"
	"gorm.io/gorm"
)

type NotificationTemplateRepositoryPostgres struct {
	db *gorm.DB
}

func NewNotificationTemplateRepository(db *gorm.DB) notificationRepository.NotificationTemplateRepository {
	return &NotificationTemplateRepositoryPostgres{db: db}
}

func (repo *NotificationTemplateRepositoryPostgres) GetTemplateById(id string) (
	notifications.
		NotificationTemplate, error,
) {
	//TODO implement
	panic("implement me")
}
func (repo *NotificationTemplateRepositoryPostgres) GetTemplateByChannelAndFlow(
	channel string,
	flow string,
) (
	notifications.NotificationTemplate,
	error,
) {
	//TODO implement
	panic("implement me")
}
func (repo *NotificationTemplateRepositoryPostgres) GetAllTemplatesByFlow(flow string) (
	[]notifications.NotificationTemplate,
	error,
) {
	//TODO implement
	panic("implement me")
}
func (repo *NotificationTemplateRepositoryPostgres) GetAllTemplatesByChannel(channel string) (
	[]notifications.NotificationTemplate,
	error,
) {
	//TODO implement
	panic("implement me")
}
func (repo *NotificationTemplateRepositoryPostgres) CreateNewTemplate(template *notifications.NotificationTemplate) (
	notifications.
		NotificationTemplate, error,
) {
	//TODO implement
	panic("implement me")
}
func (repo *NotificationTemplateRepositoryPostgres) UpdateTemplate(template *notifications.NotificationTemplate) (
	notifications.NotificationTemplate, error,
) {
	//TODO implement
	panic("implement me")
}
func (repo *NotificationTemplateRepositoryPostgres) DeleteTemplate(id string) error {
	//TODO implement
	panic("implement me")
}
