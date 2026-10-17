package model

type CourseRequest struct {
	CourseCode   string `json:"course_code"`
	CourseName   string `json:"course_name"`
	DepartmentID int64  `json:"department_id"`
	FacultyID    *int64 `json:"faculty_id"`
	Credits      int16  `json:"credits"`
	Semester     int16  `json:"semester"`
	AcademicYear string `json:"academic_year"`
}
