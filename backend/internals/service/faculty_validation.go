package service

import (
	"Frank2006x/Vmate/internals/model"
	"fmt"
	"net/mail"
	"strings"
)

func ValidateFaculty(input model.FacultyRequest) error {
	if strings.TrimSpace(input.FacultyCode) == "" || strings.TrimSpace(input.FirstName) == "" || strings.TrimSpace(input.LastName) == "" || strings.TrimSpace(input.Designation) == "" || input.DepartmentID <= 0 {
		return fmt.Errorf("faculty code, names, designation, and department are required")
	}
	if _, err := mail.ParseAddress(input.Email); err != nil {
		return fmt.Errorf("valid email is required")
	}
	if input.Status != "" && !contains([]string{"active", "on_leave", "retired"}, input.Status) {
		return fmt.Errorf("invalid faculty status")
	}
	return nil
}
