package account

import "time"

type Users struct {
	ID              uint     `json:"id"`
	AccountUniqueID string   `json:"account_unique_id"`
	Firstname       string   `json:"firstname"`
	Surename        string   `json:"surename"`
	Username        string   `json:"username"`
	Password        string   `json:"password"`
	Email           string   `json:"email"`
	Phone           string   `json:"phone"`
	Role            []string `json:"role"`
	Timestamp
}

type Timestamp struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
