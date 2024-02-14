package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Romote   bool   `json:"romote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

type OpeningResp struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	Role      string    `json:"role"`
	Company   string    `json:"company"`
	Location  string    `json:"location"`
	Romote    bool      `json:"romote"`
	Link      string    `json:"link"`
	Salary    int64     `json:"salary"`
}
