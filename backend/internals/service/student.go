package service

import (
	db "Frank2006x/Vmate/db/sqlc"
	"Frank2006x/Vmate/internals/model"
	"context"
)

type StudentService struct{ queries *db.Queries }

func NewStudentService(queries *db.Queries) *StudentService { return &StudentService{queries: queries} }
func (s *StudentService) Create(ctx context.Context, input model.StudentRequest) (db.Student, error) {
	if err := ValidateStudent(input); err != nil {
		return db.Student{}, err
	}
	date, err := optionalDate(input.EnrollmentDate)
	if err != nil {
		return db.Student{}, err
	}
	return s.queries.CreateStudent(ctx, db.CreateStudentParams{StudentCode: input.StudentCode, FirstName: input.FirstName, LastName: input.LastName, Email: input.Email, DepartmentID: input.DepartmentID, EnrollmentDate: date, Status: optionalString(input.Status)})
}
func (s *StudentService) List(ctx context.Context, params db.ListStudentsParams) ([]db.Student, error) {
	return s.queries.ListStudents(ctx, params)
}
func (s *StudentService) Get(ctx context.Context, id int64) (db.Student, error) {
	return s.queries.GetStudent(ctx, id)
}
func (s *StudentService) Update(ctx context.Context, id int64, input model.StudentRequest) (db.Student, error) {
	if err := ValidateStudent(input); err != nil {
		return db.Student{}, err
	}
	date, err := requiredDate(input.EnrollmentDate)
	if err != nil {
		return db.Student{}, err
	}
	status := input.Status
	if status == "" {
		status = "active"
	}
	return s.queries.UpdateStudent(ctx, db.UpdateStudentParams{ID: id, StudentCode: input.StudentCode, FirstName: input.FirstName, LastName: input.LastName, Email: input.Email, DepartmentID: input.DepartmentID, EnrollmentDate: date, Status: status})
}
func (s *StudentService) Delete(ctx context.Context, id int64) error {
	return s.queries.DeleteStudent(ctx, id)
}
