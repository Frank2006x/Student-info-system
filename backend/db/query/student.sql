-- name: CreateStudent :one
INSERT INTO students (student_code, first_name, last_name, email, department_id, enrollment_date, status)
VALUES ($1, $2, $3, $4, $5, COALESCE(sqlc.narg('enrollment_date'), CURRENT_DATE), COALESCE(sqlc.narg('status'), 'active')) RETURNING *;

-- name: ListStudents :many
SELECT * FROM students
WHERE (sqlc.narg('search')::text IS NULL OR student_code ILIKE '%' || sqlc.narg('search') || '%' OR first_name ILIKE '%' || sqlc.narg('search') || '%' OR last_name ILIKE '%' || sqlc.narg('search') || '%' OR email ILIKE '%' || sqlc.narg('search') || '%')
  AND (sqlc.narg('department_id')::bigint IS NULL OR department_id = sqlc.narg('department_id'))
  AND (sqlc.narg('status')::text IS NULL OR status = sqlc.narg('status'))
ORDER BY last_name, first_name LIMIT sqlc.arg('page_size') OFFSET sqlc.arg('page_offset');

-- name: GetStudent :one
SELECT * FROM students WHERE id = $1;

-- name: UpdateStudent :one
UPDATE students SET student_code = $2, first_name = $3, last_name = $4, email = $5, department_id = $6, enrollment_date = $7, status = $8, updated_at = NOW() WHERE id = $1 RETURNING *;

-- name: DeleteStudent :exec
DELETE FROM students WHERE id = $1;