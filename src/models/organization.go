package models

import "github.com/google/uuid"

type Organization struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	BelongToUser uuid.UUID `json:"belong_to_user"`
}

type OrganizationMember struct {
	OrganizationID uuid.UUID `json:"organization_id"`
	UserID         uuid.UUID `json:"user_id"`
	RoleID         string    `json:"role_id"`
	CreatedAt      string    `json:"created_at"`
	UpdatedAt      string    `json:"updated_at"`
	DeletedAt      string    `json:"deleted_at"`
}

type OrganizationRole struct {
	ID                   int       `json:"id"`
	Name                 string    `json:"name"`
	Description          string    `json:"description"`
	Category             string    `json:"category"`
	BelongToOrganization uuid.UUID `json:"belong_to_organization"`
	CreatedAt            string    `json:"created_at"`
	UpdatedAt            string    `json:"updated_at"`
	DeletedAt            string    `json:"deleted_at"`
}
