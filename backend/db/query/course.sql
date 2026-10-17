-- name: CreateCourse :one
INSERT INTO courses (course_code, course_name, department_id, faculty_id, credits, semester, academic_year)
VALUES ($1, $2, $3, sqlc.narg('faculty_id'), $4, $5, $6) RETURNING *;

-- name: ListCourses :many
SELECT * FROM courses
WHERE (sqlc.narg('search')::text IS NULL OR course_code ILIKE '%' || sqlc.narg('search') || '%' OR course_name ILIKE '%' || sqlc.narg('search') || '%')
  AND (sqlc.narg('department_id')::bigint IS NULL OR department_id = sqlc.narg('department_id'))
  AND (sqlc.narg('faculty_id')::bigint IS NULL OR faculty_id = sqlc.narg('faculty_id'))
  AND (sqlc.narg('semester')::smallint IS NULL OR semester = sqlc.narg('semester'))
  AND (sqlc.narg('academic_year')::text IS NULL OR academic_year = sqlc.narg('academic_year'))
ORDER BY course_name LIMIT sqlc.arg('page_size') OFFSET sqlc.arg('page_offset');

-- name: GetCourse :one
SELECT * FROM courses WHERE id = $1;

-- name: UpdateCourse :one
UPDATE courses SET course_code = $2, course_name = $3, department_id = $4, faculty_id = $5, credits = $6, semester = $7, academic_year = $8 WHERE id = $1 RETURNING *;

-- name: DeleteCourse :exec
DELETE FROM courses WHERE id = $1;