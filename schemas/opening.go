package schemas

import (
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
