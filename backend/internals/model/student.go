package model

type StudentRequest struct {
	StudentCode    string `json:"student_code"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	DepartmentID   int64  `json:"department_id"`
	EnrollmentDate string `json:"enrollment_date"`
	Status         string `json:"status"`
}
