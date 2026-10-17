package service

import (
	"Frank2006x/Vmate/internals/model"
	"testing"
)

func TestValidateDepartment(t *testing.T) {
	if err := ValidateDepartment(model.DepartmentRequest{DepartmentCode: "CS", DepartmentName: "Computer Science"}); err != nil {
		t.Fatalf("valid department rejected: %v", err)
	}
	if err := ValidateDepartment(model.DepartmentRequest{}); err == nil {
		t.Fatal("empty department should be rejected")
	}
}
