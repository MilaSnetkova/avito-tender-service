package models

import "time"

// Пользователь
type User struct {
	ID        int       `db:"id"`
	Username  string    `db:"username"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

//тип Организации
type OrganizationType string

const (
	IE  OrganizationType = "IE"
	LLC OrganizationType = "LLC"
	JSC OrganizationType = "JSC"
)

// Организация
type Organization struct {
	ID          int              `db:"id"`
	Name        string           `db:"name"`
	Description string           `db:"description"`
	Type        OrganizationType `db:"type"`
	CreatedAt   time.Time        `db:"created_at"`
	UpdatedAt   time.Time        `db:"updated_at"`
}

// Ответственный за организацию
type OrganizationResponsible struct {
	ID             int `db:"id"`
	OrganizationID int `db:"organization_id"`
	UserID         int `db:"user_id"`
}

// Тендер
type Tender struct {
	ID             int       `db:"id"`
	Title          string    `db:"title"`
	Description    string    `db:"description"`
	Status         string    `db:"status"`
	Version        int       `db:"version"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
	OrganizationID int       `db:"organization_id"`
}

// Предложение
type Offer struct {
	ID          int       `db:"id"`
	Description string    `db:"description"`
	Status      string    `db:"status"`
	Version     int       `db:"version"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	TenderID    int       `db:"tender_id"`
	UserID      int       `db:"user_id"`
}
