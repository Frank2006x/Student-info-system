package service

import (
	"Frank2006x/Vmate/internals/model"
	"fmt"
	"net/mail"
	"strings"
)

func ValidateStudent(input model.StudentRequest) error {
	if strings.TrimSpace(input.StudentCode) == "" || strings.TrimSpace(input.FirstName) == "" || strings.TrimSpace(input.LastName) == "" || input.DepartmentID <= 0 {
		return fmt.Errorf("student code, names, and department are required")
	}
	if _, err := mail.ParseAddress(input.Email); err != nil {
		return fmt.Errorf("valid email is required")
	}
	if input.Status != "" && !contains([]string{"active", "graduated", "withdrawn"}, input.Status) {
		return fmt.Errorf("invalid student status")
	}
	return nil
}
