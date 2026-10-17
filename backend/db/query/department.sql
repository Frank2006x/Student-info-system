-- name: CreateDepartment :one
INSERT INTO departments (department_code, department_name) VALUES ($1, $2) RETURNING *;

-- name: ListDepartments :many
SELECT * FROM departments ORDER BY department_name;

-- name: GetDepartment :one
SELECT * FROM departments WHERE id = $1;

-- name: UpdateDepartment :one
UPDATE departments SET department_code = $2, department_name = $3, updated_at = NOW() WHERE id = $1 RETURNING *;

-- name: DeleteDepartment :exec
DELETE FROM departments WHERE id = $1;