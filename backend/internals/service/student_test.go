package service

import (
	"Frank2006x/Vmate/internals/model"
	"testing"
)

func TestValidateStudent(t *testing.T) {
	input := model.StudentRequest{StudentCode: "S-1", FirstName: "Ada", LastName: "Lovelace", Email: "ada@example.com", DepartmentID: 1}
	if err := ValidateStudent(input); err != nil {
		t.Fatalf("valid student rejected: %v", err)
	}
	input.Status = "unknown"
	if err := ValidateStudent(input); err == nil {
		t.Fatal("invalid student status should be rejected")
	}
}
