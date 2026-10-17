package service

import (
	"Frank2006x/Vmate/internals/model"
	"fmt"
	"strings"
)

func ValidateCourse(input model.CourseRequest) error {
	if strings.TrimSpace(input.CourseCode) == "" || strings.TrimSpace(input.CourseName) == "" || strings.TrimSpace(input.AcademicYear) == "" || input.DepartmentID <= 0 {
		return fmt.Errorf("course code, name, academic year, and department are required")
	}
	if input.Credits < 1 || input.Credits > 10 {
		return fmt.Errorf("credits must be between 1 and 10")
	}
	if input.Semester < 1 || input.Semester > 12 {
		return fmt.Errorf("semester must be between 1 and 12")
	}
	return nil
}

func contains(values []string, value string) bool {
	for _, item := range values {
		if item == value {
			return true
		}
	}
	return false
}
