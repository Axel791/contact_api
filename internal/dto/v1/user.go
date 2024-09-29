package v1

type UserDetail struct {
	ID          int       `json:"id" validate:"required"`
	FirstName   string    `json:"first_name" validate:"required"`
	LastName    string    `json:"last_name" validate:"required"`
	Age         int       `json:"age" validate:"required,gte=0"`
	Description string    `json:"description,omitempty"`
	Contacts    []Contact `json:"contacts,omitempty"`
}

type UserSummary struct {
	ID          int    `json:"id" validate:"required"`
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Age         int    `json:"age,omitempty" validate:"required,gte=0"`
	Description string `json:"description,omitempty"`
}

type UserList struct {
	Data []UserSummary `json:"data"`
}
