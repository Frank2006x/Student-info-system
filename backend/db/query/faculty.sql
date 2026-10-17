-- name: CreateFaculty :one
INSERT INTO faculty (faculty_code, first_name, last_name, email, department_id, designation, specialization, hire_date, status)
VALUES ($1, $2, $3, $4, $5, $6, sqlc.narg('specialization'), COALESCE(sqlc.narg('hire_date'), CURRENT_DATE), COALESCE(sqlc.narg('status'), 'active')) RETURNING *;

-- name: ListFaculty :many
SELECT * FROM faculty
WHERE (sqlc.narg('search')::text IS NULL OR faculty_code ILIKE '%' || sqlc.narg('search') || '%' OR first_name ILIKE '%' || sqlc.narg('search') || '%' OR last_name ILIKE '%' || sqlc.narg('search') || '%' OR email ILIKE '%' || sqlc.narg('search') || '%')
  AND (sqlc.narg('department_id')::bigint IS NULL OR department_id = sqlc.narg('department_id'))
  AND (sqlc.narg('status')::text IS NULL OR status = sqlc.narg('status'))
ORDER BY last_name, first_name LIMIT sqlc.arg('page_size') OFFSET sqlc.arg('page_offset');

-- name: GetFaculty :one
SELECT * FROM faculty WHERE id = $1;

-- name: UpdateFaculty :one
UPDATE faculty SET faculty_code = $2, first_name = $3, last_name = $4, email = $5, department_id = $6, designation = $7, specialization = $8, hire_date = $9, status = $10, updated_at = NOW() WHERE id = $1 RETURNING *;

-- name: DeleteFaculty :exec
DELETE FROM faculty WHERE id = $1;