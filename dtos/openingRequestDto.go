package dtos

import "fmt"

func ErrParamIsRequired(name, typ string) error {
	return fmt.Errorf("paramr: %s (type: %s) is required", name, typ)
}

type CreateOpeningDto struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningDto) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Remote == nil && r.Link == "" && r.Salary <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}

	if r.Role == "" {
		return ErrParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return ErrParamIsRequired("company", "string")
	}
	if r.Location == "" {
		return ErrParamIsRequired("location", "string")
	}
	if r.Remote == nil {
		return ErrParamIsRequired("remote", "bool")
	}
	if r.Link == "" {
		return ErrParamIsRequired("link", "string")
	}
	if r.Salary <= 0 {
		return ErrParamIsRequired("salary", "int64")
	}

	return nil
}
