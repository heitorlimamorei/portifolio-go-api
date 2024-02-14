package handlers

import "fmt"

// Defines the Struct to be hydrated with the request body

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Romote   bool   `json:"romote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r == nil {
		return fmt.Errorf("malformed request body")
	}

	if r.Role == "" {
		return fmt.Errorf("role is required")
	}

	if r.Company == "" {
		return fmt.Errorf("company is required")
	}

	if r.Location == "" {
		return fmt.Errorf("location is required")
	}

	if r.Link == "" {
		return fmt.Errorf("link is required")
	}

	if r.Salary <= 0 {
		return fmt.Errorf("salary is required")
	}

	return nil
}
