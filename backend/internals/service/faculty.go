package service

import (
	db "Frank2006x/Vmate/db/sqlc"
	"Frank2006x/Vmate/internals/model"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type FacultyService struct {
	readQueries  *db.Queries
	writeQueries *db.Queries
}

func NewFacultyService(readQueries, writeQueries *db.Queries) *FacultyService {
	return &FacultyService{readQueries: readQueries, writeQueries: writeQueries}
}
func (s *FacultyService) Create(ctx context.Context, input model.FacultyRequest) (db.Faculty, error) {
	if err := ValidateFaculty(input); err != nil {
		return db.Faculty{}, err
	}
	hireDate, err := optionalDate(input.HireDate)
	if err != nil {
		return db.Faculty{}, err
	}
	return s.writeQueries.CreateFaculty(ctx, db.CreateFacultyParams{FacultyCode: input.FacultyCode, FirstName: input.FirstName, LastName: input.LastName, Email: input.Email, DepartmentID: input.DepartmentID, Designation: input.Designation, Specialization: pgtype.Text{String: input.Specialization, Valid: input.Specialization != ""}, HireDate: hireDate, Status: optionalString(input.Status)})
}
func (s *FacultyService) List(ctx context.Context, params db.ListFacultyParams) ([]db.Faculty, error) {
	return s.readQueries.ListFaculty(ctx, params)
}
func (s *FacultyService) Get(ctx context.Context, id int64) (db.Faculty, error) {
	return s.readQueries.GetFaculty(ctx, id)
}
func (s *FacultyService) Update(ctx context.Context, id int64, input model.FacultyRequest) (db.Faculty, error) {
	if err := ValidateFaculty(input); err != nil {
		return db.Faculty{}, err
	}
	hireDate, err := requiredDate(input.HireDate)
	if err != nil {
		return db.Faculty{}, err
	}
	status := input.Status
	if status == "" {
		status = "active"
	}
	return s.writeQueries.UpdateFaculty(ctx, db.UpdateFacultyParams{ID: id, FacultyCode: input.FacultyCode, FirstName: input.FirstName, LastName: input.LastName, Email: input.Email, DepartmentID: input.DepartmentID, Designation: input.Designation, Specialization: pgtype.Text{String: input.Specialization, Valid: input.Specialization != ""}, HireDate: hireDate, Status: status})
}
func (s *FacultyService) Delete(ctx context.Context, id int64) error {
	return s.writeQueries.DeleteFaculty(ctx, id)
}
