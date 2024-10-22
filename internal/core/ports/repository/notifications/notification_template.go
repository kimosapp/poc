package notifications

import "github.com/kimosapp/poc/internal/core/model/entity/notifications"

type NotificationTemplateRepository interface {
	GetTemplateById(id string) (notifications.NotificationTemplate, error)
	GetTemplateByChannelAndFlow(channel string, flow string) (
		notifications.NotificationTemplate,
		error,
	)
	GetAllTemplatesByFlow(flow string) ([]notifications.NotificationTemplate, error)
	GetAllTemplatesByChannel(channel string) ([]notifications.NotificationTemplate, error)
	CreateNewTemplate(template *notifications.NotificationTemplate) (
		notifications.
			NotificationTemplate, error,
	)
	UpdateTemplate(template *notifications.NotificationTemplate) (
		notifications.NotificationTemplate, error,
	)
	DeleteTemplate(id string) error
}
