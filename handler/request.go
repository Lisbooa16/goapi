package handler

import "fmt"

// CreateOpening
func errParmsIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   bool   `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" {
		return errParmsIsRequired("role", "string")
	}
	if r.Company == "" {
		return errParmsIsRequired("company", "string")
	}
	if r.Location == "" {
		return errParmsIsRequired("location", "string")
	}
	if r.Link == "" {
		return errParmsIsRequired("link", "string")
	}
	if r.Remote != true && r.Remote != false {
		return errParmsIsRequired("remote", "bool")
	}
	if r.Salary <= 0 {
		return errParmsIsRequired("salary", "int64")
	}
	return nil
}
