package model

type FacultyRequest struct {
	FacultyCode    string `json:"faculty_code"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	DepartmentID   int64  `json:"department_id"`
	Designation    string `json:"designation"`
	Specialization string `json:"specialization"`
	HireDate       string `json:"hire_date"`
	Status         string `json:"status"`
}
