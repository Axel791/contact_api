package v1

type ContactType string

const (
	ContactTypePhone    ContactType = "phone"
	ContactTypeEmail    ContactType = "email"
	ContactTypeTelegram ContactType = "telegram"
	ContactTypeWhatsApp ContactType = "whatsapp"
	ContactTypeVK       ContactType = "vk"
)

type Contact struct {
	ID          int         `json:"id" validate:"required"`
	Type        ContactType `json:"number" validate:"required"`
	Description string      `json:"Description,omitempty"`
}

type UpdateOrCreateContact struct {
	Type        ContactType `json:"number" validate:"required"`
	Description string      `json:"Description,omitempty"`
}
