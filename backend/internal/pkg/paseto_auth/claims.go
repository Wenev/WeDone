package pasetoauth

import "github.com/google/uuid"

type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	Role string `json:"role"`
	IsAdmin bool `json:"is_admin"`
	ManagerID *uuid.UUID `json:"manager_id,omitempty"`
}