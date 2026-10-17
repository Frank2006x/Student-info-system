package service

import (
	"Frank2006x/Vmate/internals/model"
	"testing"
)

func TestValidateFaculty(t *testing.T) {
	input := model.FacultyRequest{FacultyCode: "F-1", FirstName: "Grace", LastName: "Hopper", Email: "grace@example.com", DepartmentID: 1, Designation: "Professor"}
	if err := ValidateFaculty(input); err != nil {
		t.Fatalf("valid faculty rejected: %v", err)
	}
	input.Email = "bad"
	if err := ValidateFaculty(input); err == nil {
		t.Fatal("invalid faculty email should be rejected")
	}
}
