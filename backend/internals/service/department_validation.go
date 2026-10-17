package service

import (
	"Frank2006x/Vmate/internals/model"
	"fmt"
	"strings"
)

func ValidateDepartment(input model.DepartmentRequest) error {
	if strings.TrimSpace(input.DepartmentCode) == "" || strings.TrimSpace(input.DepartmentName) == "" {
		return fmt.Errorf("department code and name are required")
	}
	return nil
}
