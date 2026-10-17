package service

import (
	"Frank2006x/Vmate/internals/model"
	"testing"
)

func TestValidateCourse(t *testing.T) {
	input := model.CourseRequest{CourseCode: "CS101", CourseName: "Algorithms", DepartmentID: 1, Credits: 3, Semester: 1, AcademicYear: "2026-2027"}
	if err := ValidateCourse(input); err != nil {
		t.Fatalf("valid course rejected: %v", err)
	}
	input.Credits = 11
	if err := ValidateCourse(input); err == nil {
		t.Fatal("out-of-range credits should be rejected")
	}
}
