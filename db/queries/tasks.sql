-- name: ListTasks :many
SELECT * FROM tasks ORDER BY created_at OFFSET $1 LIMIT $2;

-- name: GetTask :one
SELECT * FROM tasks WHERE id = $1 LIMIT 1;

-- name: InsertTask :one
INSERT INTO tasks(id, name, description)
VALUES (uuid_generate_v4(), $1, $2)
RETURNING *;

-- name: UpdateTask :one
UPDATE tasks
SET name = $2,
    description = $3,
    done = $4,
    updated_at = NOW()
WHERE id = $1
RETURNING *;