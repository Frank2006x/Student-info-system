package service

import (
	db "Frank2006x/Vmate/db/sqlc"
	"Frank2006x/Vmate/internals/model"
	"context"
)

type StudentService struct {
	readQueries  *db.Queries
	writeQueries *db.Queries
}

func NewStudentService(readQueries, writeQueries *db.Queries) *StudentService {
	return &StudentService{readQueries: readQueries, writeQueries: writeQueries}
}
func (s *StudentService) Create(ctx context.Context, input model.StudentRequest) (db.Student, error) {
	if err := ValidateStudent(input); err != nil {
		return db.Student{}, err
	}
	date, err := optionalDate(input.EnrollmentDate)
	if err != nil {
		return db.Student{}, err
	}
	return s.writeQueries.CreateStudent(ctx, db.CreateStudentParams{StudentCode: input.StudentCode, FirstName: input.FirstName, LastName: input.LastName, Email: input.Email, DepartmentID: input.DepartmentID, EnrollmentDate: date, Status: optionalString(input.Status)})
}
func (s *StudentService) List(ctx context.Context, params db.ListStudentsParams) ([]db.Student, error) {
	return s.readQueries.ListStudents(ctx, params)
}
func (s *StudentService) Get(ctx context.Context, id int64) (db.Student, error) {
	return s.readQueries.GetStudent(ctx, id)
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
	return s.writeQueries.UpdateStudent(ctx, db.UpdateStudentParams{ID: id, StudentCode: input.StudentCode, FirstName: input.FirstName, LastName: input.LastName, Email: input.Email, DepartmentID: input.DepartmentID, EnrollmentDate: date, Status: status})
}
func (s *StudentService) Delete(ctx context.Context, id int64) error {
	return s.writeQueries.DeleteStudent(ctx, id)
}
