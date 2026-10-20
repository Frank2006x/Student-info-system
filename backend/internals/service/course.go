package service

import (
	db "Frank2006x/Vmate/db/sqlc"
	"Frank2006x/Vmate/internals/model"
	"context"
)

type CourseService struct {
	readQueries  *db.Queries
	writeQueries *db.Queries
}

func NewCourseService(readQueries, writeQueries *db.Queries) *CourseService {
	return &CourseService{readQueries: readQueries, writeQueries: writeQueries}
}
func (s *CourseService) Create(ctx context.Context, input model.CourseRequest) (db.Course, error) {
	if err := ValidateCourse(input); err != nil {
		return db.Course{}, err
	}
	return s.writeQueries.CreateCourse(ctx, db.CreateCourseParams{CourseCode: input.CourseCode, CourseName: input.CourseName, DepartmentID: input.DepartmentID, FacultyID: nullableID(input.FacultyID), Credits: input.Credits, Semester: input.Semester, AcademicYear: input.AcademicYear})
}
func (s *CourseService) List(ctx context.Context, params db.ListCoursesParams) ([]db.Course, error) {
	return s.readQueries.ListCourses(ctx, params)
}
func (s *CourseService) Get(ctx context.Context, id int64) (db.Course, error) {
	return s.readQueries.GetCourse(ctx, id)
}
func (s *CourseService) Update(ctx context.Context, id int64, input model.CourseRequest) (db.Course, error) {
	if err := ValidateCourse(input); err != nil {
		return db.Course{}, err
	}
	return s.writeQueries.UpdateCourse(ctx, db.UpdateCourseParams{ID: id, CourseCode: input.CourseCode, CourseName: input.CourseName, DepartmentID: input.DepartmentID, FacultyID: nullableID(input.FacultyID), Credits: input.Credits, Semester: input.Semester, AcademicYear: input.AcademicYear})
}
func (s *CourseService) Delete(ctx context.Context, id int64) error {
	return s.writeQueries.DeleteCourse(ctx, id)
}
