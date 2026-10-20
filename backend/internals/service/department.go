package service

import (
	"context"

	db "Frank2006x/Vmate/db/sqlc"
	"Frank2006x/Vmate/internals/model"
)

type DepartmentService struct {
	readQueries  *db.Queries
	writeQueries *db.Queries
}

func NewDepartmentService(readQueries, writeQueries *db.Queries) *DepartmentService {
	return &DepartmentService{readQueries: readQueries, writeQueries: writeQueries}
}
func (s *DepartmentService) Create(ctx context.Context, input model.DepartmentRequest) (db.Department, error) {
	if err := ValidateDepartment(input); err != nil {
		return db.Department{}, err
	}
	return s.writeQueries.CreateDepartment(ctx, db.CreateDepartmentParams{DepartmentCode: input.DepartmentCode, DepartmentName: input.DepartmentName})
}
func (s *DepartmentService) List(ctx context.Context) ([]db.Department, error) {
	return s.readQueries.ListDepartments(ctx)
}
func (s *DepartmentService) Get(ctx context.Context, id int64) (db.Department, error) {
	return s.readQueries.GetDepartment(ctx, id)
}
func (s *DepartmentService) Update(ctx context.Context, id int64, input model.DepartmentRequest) (db.Department, error) {
	if err := ValidateDepartment(input); err != nil {
		return db.Department{}, err
	}
	return s.writeQueries.UpdateDepartment(ctx, db.UpdateDepartmentParams{ID: id, DepartmentCode: input.DepartmentCode, DepartmentName: input.DepartmentName})
}
func (s *DepartmentService) Delete(ctx context.Context, id int64) error {
	return s.writeQueries.DeleteDepartment(ctx, id)
}
